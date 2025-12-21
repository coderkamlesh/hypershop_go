package handler

import (
	"strings"

	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/coderkamlesh/hypershop_go/internal/service"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// ==================== REGISTRATION ====================

func (h *AuthHandler) RequestConsumerRegistrationOTP(c *gin.Context) {
	var req dto.RegistrationOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	err := h.authService.SendConsumerRegistrationOTP(req.Mobile)
	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	maskedMobile := utils.MaskMobile(req.Mobile)
	response := dto.SuccessWithoutData("Registration OTP sent successfully to " + maskedMobile)
	c.JSON(200, response)
}

func (h *AuthHandler) VerifyConsumerRegistrationOTP(c *gin.Context) {
	var req dto.RegistrationOtpVerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	authResponse, err := h.authService.VerifyConsumerRegistrationOTP(
		req.Mobile,
		req.Name,
		req.OTP,
		c.GetHeader("User-Agent"),
		"mobile",
	)

	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	response := dto.Success("Registration + login successful", authResponse)
	c.JSON(200, response)
}

// ==================== LOGIN ====================

func (h *AuthHandler) RequestConsumerLoginOTP(c *gin.Context) {
	var req dto.OtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err) // ✅ Proper validation
		return
	}

	err := h.authService.SendConsumerLoginOTP(req.Mobile)
	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	maskedMobile := utils.MaskMobile(req.Mobile)
	response := dto.SuccessWithoutData("OTP sent successfully to " + maskedMobile)
	c.JSON(200, response)
}

func (h *AuthHandler) VerifyConsumerLoginOTP(c *gin.Context) {
	var req dto.OtpVerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err) // ✅ Proper validation
		return
	}

	authResponse, err := h.authService.VerifyLoginOTP(
		req.Mobile,
		req.OTP,
		req.DeviceInfo,
		req.Source,
	)

	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	response := dto.Success("Login successful", authResponse)
	c.JSON(200, response)
}

// ==================== RIDER LOGIN ====================

func (h *AuthHandler) RequestRiderLoginOTP(c *gin.Context) {
	var req dto.OtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	err := h.authService.SendRiderLoginOTP(req.Mobile)
	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	maskedMobile := utils.MaskMobile(req.Mobile)
	response := dto.SuccessWithoutData("OTP sent successfully to " + maskedMobile)
	c.JSON(200, response)
}

func (h *AuthHandler) VerifyRiderLoginOTP(c *gin.Context) {
	var req dto.OtpVerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	authResponse, err := h.authService.VerifyLoginOTP(
		req.Mobile,
		req.OTP,
		req.DeviceInfo,
		req.Source,
	)

	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	response := dto.Success("Login successful", authResponse)
	c.JSON(200, response)
}

// ==================== MANAGER LOGIN ====================

func (h *AuthHandler) RequestManagerLoginOTP(c *gin.Context) {
	var req dto.OtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	err := h.authService.SendManagerLoginOTP(req.Mobile)
	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	maskedMobile := utils.MaskMobile(req.Mobile)
	response := dto.SuccessWithoutData("OTP sent successfully to " + maskedMobile)
	c.JSON(200, response)
}

func (h *AuthHandler) VerifyManagerLoginOTP(c *gin.Context) {
	var req dto.OtpVerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	authResponse, err := h.authService.VerifyLoginOTP(
		req.Mobile,
		req.OTP,
		req.DeviceInfo,
		req.Source,
	)

	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	response := dto.Success("Login successful", authResponse)
	c.JSON(200, response)
}

// ==================== ADMIN LOGIN ====================

func (h *AuthHandler) RequestAdminLoginOTP(c *gin.Context) {
	var req dto.OtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	err := h.authService.SendAdminLoginOTP(req.Mobile)
	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	maskedMobile := utils.MaskMobile(req.Mobile)
	response := dto.SuccessWithoutData("OTP sent successfully to " + maskedMobile)
	c.JSON(200, response)
}

func (h *AuthHandler) VerifyAdminLoginOTP(c *gin.Context) {
	var req dto.OtpVerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	authResponse, err := h.authService.VerifyLoginOTP(
		req.Mobile,
		req.OTP,
		req.DeviceInfo,
		req.Source,
	)

	if err != nil {
		response := dto.Failure(err.Error())
		c.JSON(400, response)
		return
	}

	response := dto.Success("Login successful", authResponse)
	c.JSON(200, response)
}

// ==================== TOKEN VALIDATION ====================

func (h *AuthHandler) ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		response := dto.InvalidToken()
		c.JSON(401, response)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	err := h.authService.ValidateToken(token)
	if err != nil {
		response := dto.InvalidTokenWithMessage(err.Error())
		c.JSON(401, response)
		return
	}

	response := dto.SuccessWithoutData("Token is valid")
	c.JSON(200, response)
}
