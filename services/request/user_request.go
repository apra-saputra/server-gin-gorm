package request

import "time"

type FormUser struct {
	Email string `form:"email" json:"email" binding:"required"`
	OTP   string `form:"otp" json:"otp" binding:"required"`
}

type FormEmail struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type OTPData struct {
	OTP       string    `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`
}
