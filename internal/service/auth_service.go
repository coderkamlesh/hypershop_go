package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/constants"
	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/repository"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
)

const (
	OTP_RATE_LIMIT_SECONDS = 60 // 1 minute
	OTP_EXPIRY_MINUTES     = 5  // 5 minutes
)

type AuthService interface {
	SendConsumerRegistrationOTP(mobile string) error
	VerifyConsumerRegistrationOTP(mobile, name, otp, deviceInfo, source string) (*dto.AuthResponse, error)
	SendConsumerLoginOTP(mobile string) error
	SendRiderLoginOTP(mobile string) error
	SendManagerLoginOTP(mobile string) error
	SendAdminLoginOTP(mobile string) error
	VerifyLoginOTP(mobile, otp, deviceInfo, source string) (*dto.AuthResponse, error)
	ValidateToken(token string) error
}

type authService struct {
	userRepo        repository.UserRepository
	regOtpRepo      repository.RegistrationOtpRepository
	userOtpRepo     repository.UserOtpRepository
	userSessionRepo repository.UserSessionRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	regOtpRepo repository.RegistrationOtpRepository,
	userOtpRepo repository.UserOtpRepository,
	userSessionRepo repository.UserSessionRepository,
) AuthService {
	return &authService{
		userRepo:        userRepo,
		regOtpRepo:      regOtpRepo,
		userOtpRepo:     userOtpRepo,
		userSessionRepo: userSessionRepo,
	}
}

// ==================== REGISTRATION ====================

func (s *authService) SendConsumerRegistrationOTP(mobile string) error {
	log.Printf("Registration OTP requested for mobile: %s", mobile)

	// Check if user already exists
	existingUser, _ := s.userRepo.FindByMobile(mobile)
	if existingUser != nil {
		return errors.New("user already exists with this mobile")
	}

	now := time.Now()
	existingOtp, _ := s.regOtpRepo.FindByMobile(mobile)

	otpValue := generateOTP(mobile)

	if existingOtp != nil {
		// Check rate limiting
		secondsSinceLastOtp := int(now.Sub(existingOtp.CreatedAt).Seconds())
		if secondsSinceLastOtp < OTP_RATE_LIMIT_SECONDS {
			secondsLeft := OTP_RATE_LIMIT_SECONDS - secondsSinceLastOtp
			return fmt.Errorf("please wait %d seconds before requesting a new OTP", secondsLeft)
		}

		// Update existing OTP
		existingOtp.OTP = otpValue
		existingOtp.CreatedAt = now
		existingOtp.ExpiredAt = now.Add(time.Minute * OTP_EXPIRY_MINUTES)
		existingOtp.AttemptCount = 0
		existingOtp.Status = true

		if err := s.regOtpRepo.Update(existingOtp); err != nil {
			return err
		}
	} else {
		// Create new OTP
		newOtp := &models.RegistrationOtp{
			Mobile:       mobile,
			OTP:          otpValue,
			AttemptCount: 0,
			Status:       true,
			ExpiredAt:    now.Add(time.Minute * OTP_EXPIRY_MINUTES),
		}

		if err := s.regOtpRepo.Create(newOtp); err != nil {
			return err
		}
	}

	// TODO: Send SMS
	log.Printf("Registration OTP sent via SMS: %s", otpValue)

	return nil
}

func (s *authService) VerifyConsumerRegistrationOTP(mobile, name, otp, deviceInfo, source string) (*dto.AuthResponse, error) {
	// Find OTP
	regOtp, _ := s.regOtpRepo.FindByMobile(mobile)
	if regOtp == nil {
		return nil, errors.New("no registration OTP found. Please request a new OTP")
	}

	if !regOtp.Status {
		return nil, errors.New("OTP is no longer active. Please request a new one")
	}

	if time.Now().After(regOtp.ExpiredAt) {
		regOtp.Status = false
		s.regOtpRepo.Update(regOtp)
		return nil, errors.New("OTP has expired. Please request a new one")
	}

	if regOtp.AttemptCount >= 3 {
		regOtp.Status = false
		s.regOtpRepo.Update(regOtp)
		return nil, errors.New("maximum attempts exceeded. Please request a new OTP")
	}

	// Increment attempt
	regOtp.AttemptCount++
	s.regOtpRepo.Update(regOtp)

	// Verify OTP
	if regOtp.OTP != otp {
		attemptsLeft := 3 - regOtp.AttemptCount
		return nil, fmt.Errorf("invalid OTP. %d attempt(s) remaining", attemptsLeft)
	}

	// Mark OTP as used
	regOtp.Status = false
	s.regOtpRepo.Update(regOtp)

	// Check if user already exists
	existingUser, _ := s.userRepo.FindByMobile(mobile)
	if existingUser != nil {
		return nil, errors.New("user already exists. Please login instead")
	}

	// Create user
	user := &models.User{
		Name:     name,
		Mobile:   mobile,
		Email:    "",
		Role:     constants.RoleConsumer,
		Password: "NA",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	// Create session
	session := &models.UserSession{
		UserID:     user.ID,
		Token:      token,
		DeviceInfo: deviceInfo,
		Source:     source,
	}
	s.userSessionRepo.Create(session)

	response := &dto.AuthResponse{
		Token:  token,
		UserID: user.ID,
		Name:   user.Name,
		Mobile: user.Mobile,
		Role:   string(user.Role),
	}

	log.Printf("User registered successfully: %s", user.ID)

	return response, nil
}

// ==================== LOGIN ====================

func (s *authService) SendConsumerLoginOTP(mobile string) error {
	return s.sendLoginOTP(mobile, constants.RoleConsumer)
}

func (s *authService) SendRiderLoginOTP(mobile string) error {
	return s.sendLoginOTP(mobile, constants.RoleRider)
}

func (s *authService) SendManagerLoginOTP(mobile string) error {
	return s.sendLoginOTP(mobile, constants.RoleWarehouseManager, constants.RoleStoreManager)
}

func (s *authService) SendAdminLoginOTP(mobile string) error {
	return s.sendLoginOTP(mobile, constants.RoleAdmin, constants.RoleCatalogAdmin)
}

func (s *authService) sendLoginOTP(mobile string, allowedRoles ...constants.Role) error {
	log.Printf("Login OTP requested for mobile: %s", mobile)

	// Find user
	user, _ := s.userRepo.FindByMobile(mobile)
	if user == nil {
		return errors.New("user not found. Please register first")
	}

	// Check role
	roleAllowed := false
	for _, role := range allowedRoles {
		if user.Role == role {
			roleAllowed = true
			break
		}
	}

	if !roleAllowed {
		return errors.New("invalid user role for this login endpoint")
	}

	now := time.Now()
	existingOtp, _ := s.userOtpRepo.FindByMobile(mobile)

	otpValue := generateOTP(mobile)

	if existingOtp != nil {
		// Check rate limiting
		secondsSinceLastOtp := int(now.Sub(existingOtp.CreatedAt).Seconds())
		if secondsSinceLastOtp < OTP_RATE_LIMIT_SECONDS {
			secondsLeft := OTP_RATE_LIMIT_SECONDS - secondsSinceLastOtp
			return fmt.Errorf("please wait %d seconds before requesting a new OTP", secondsLeft)
		}

		// Update existing OTP
		existingOtp.OTP = otpValue
		existingOtp.UserID = user.ID
		existingOtp.CreatedAt = now
		existingOtp.ExpiredAt = now.Add(time.Minute * OTP_EXPIRY_MINUTES)
		existingOtp.AttemptCount = 0
		existingOtp.Status = true

		if err := s.userOtpRepo.Update(existingOtp); err != nil {
			return err
		}
	} else {
		// Create new OTP
		newOtp := &models.UserOtp{
			UserID:       user.ID,
			Mobile:       mobile,
			OTP:          otpValue,
			AttemptCount: 0,
			Status:       true,
			ExpiredAt:    now.Add(time.Minute * OTP_EXPIRY_MINUTES),
		}

		if err := s.userOtpRepo.Create(newOtp); err != nil {
			return err
		}
	}

	// TODO: Send SMS
	log.Printf("Login OTP sent via SMS: %s", otpValue)

	return nil
}

func (s *authService) VerifyLoginOTP(mobile, otp, deviceInfo, source string) (*dto.AuthResponse, error) {
	// Find OTP
	userOtp, _ := s.userOtpRepo.FindByMobile(mobile)
	if userOtp == nil {
		return nil, errors.New("no OTP found. Please request a new OTP")
	}

	if !userOtp.Status {
		return nil, errors.New("OTP is no longer active. Please request a new one")
	}

	if time.Now().After(userOtp.ExpiredAt) {
		userOtp.Status = false
		s.userOtpRepo.Update(userOtp)
		return nil, errors.New("OTP has expired. Please request a new one")
	}

	if userOtp.AttemptCount >= 3 {
		userOtp.Status = false
		s.userOtpRepo.Update(userOtp)
		return nil, errors.New("maximum attempts exceeded. Please request a new OTP")
	}

	// Increment attempt
	userOtp.AttemptCount++
	s.userOtpRepo.Update(userOtp)

	// Verify OTP
	if userOtp.OTP != otp {
		attemptsLeft := 3 - userOtp.AttemptCount
		return nil, fmt.Errorf("invalid OTP. %d attempt(s) remaining", attemptsLeft)
	}

	// Mark OTP as used
	userOtp.Status = false
	s.userOtpRepo.Update(userOtp)

	// Find user
	user, err := s.userRepo.FindByMobile(mobile)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	// Create session
	session := &models.UserSession{
		UserID:     user.ID,
		Token:      token,
		DeviceInfo: deviceInfo,
		Source:     source,
	}
	s.userSessionRepo.Create(session)

	response := &dto.AuthResponse{
		Token:  token,
		UserID: user.ID,
		Name:   user.Name,
		Mobile: user.Mobile,
		Role:   string(user.Role),
	}

	log.Printf("User logged in successfully: %s", user.ID)

	return response, nil
}

// ==================== TOKEN VALIDATION ====================

func (s *authService) ValidateToken(token string) error {
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	if utils.IsTokenExpired(token) {
		return errors.New("token expired")
	}

	// Optional: Check session
	session, _ := s.userSessionRepo.FindByToken(token)
	if session == nil || !session.IsActive {
		return errors.New("session not found or inactive")
	}

	// Update last used
	s.userSessionRepo.UpdateLastUsed(session.ID)

	log.Printf("Token validated for user: %s", claims.UserID)

	return nil
}

// ==================== HELPER FUNCTIONS ====================

func generateOTP(mobile string) string {
	// Test mobiles for development
	testMobiles := map[string]string{
		"1111111111": "111111",
		"9222222222": "222222",
		"9333333333": "333333",
		"9444444444": "444444",
		"9555555555": "555555",
		"9666666666": "666666",
		"9777777777": "777777",
	}

	if otp, exists := testMobiles[mobile]; exists {
		return otp
	}

	return utils.GenerateOTP()
}
