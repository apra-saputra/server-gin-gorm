package customTypes

import "time"

type ReqFormUser struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type OTPData struct {
	OTP       string    `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ReqFormEmail struct {
	Email string `json:"email"`
}

type ReqFormTask struct {
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date" time_format:"2006-01-02"`
}
