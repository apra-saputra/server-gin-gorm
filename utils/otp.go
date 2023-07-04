package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi untuk generate OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return fmt.Sprintf("%06d", rand.Intn(max-min+1)+min)
}
