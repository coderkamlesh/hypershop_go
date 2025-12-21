package dto

// ============ REQUEST DTOs ============

type RegistrationOtpRequest struct {
	Mobile string `json:"mobile" binding:"required,len=10,numeric"`
}

type RegistrationOtpVerifyRequest struct {
	Mobile string `json:"mobile" binding:"required,len=10,numeric" validate:"required"`
	Name   string `json:"name" binding:"required,min=2,max=50" validate:"required"`
	OTP    string `json:"otp" binding:"required,len=6,numeric" validate:"required"`
}

type OtpRequest struct {
	Mobile string `json:"mobile" binding:"required,len=10,numeric" validate:"required"`
}

type OtpVerifyRequest struct {
	Mobile     string `json:"mobile" binding:"required,len=10,numeric" validate:"required"`
	OTP        string `json:"otp" binding:"required,len=6,numeric" validate:"required"`
	DeviceInfo string `json:"device_info"`
	Source     string `json:"source"` // "android", "ios", "web"
}

// ============ RESPONSE DTOs ============

type AuthResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Role   string `json:"role"`
}
