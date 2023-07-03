package controllers

import (
	"fmt"
	"net/http"
	"restapi/initializers"
	"restapi/models"
	"restapi/utils"
	"restapi/utils/customTypes"
	"time"

	"github.com/gin-gonic/gin"
)

var otpDataMap map[string]customTypes.OTPData

func Login(c *gin.Context) {
	// Mendapatkan data dari body request
	var request customTypes.ReqFormEmail

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	var user models.User
	result := initializers.DB.Where("email = ?", request.Email).Find(&user)
	if result.Error != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Data Not Found", result.Error)
		return
	}

	otp := utils.GenerateOTP()

	initializers.DB.Model(&models.User{}).Where("email = ?", request.Email).Update("otp", otp)

	fmt.Println(otp)

	type Response struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	response := Response{
		Email:    user.Email,
		Username: user.Username,
	}

	utils.JSONResponse(c, http.StatusOK, response, "OTP has been sent to your email", nil)
}

func OTP(c *gin.Context) {
	// Mendapatkan Otp
	var request customTypes.ReqFormUser

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	// validasi Otp, waktu dan sesuai atau tidak
	var user models.User
	result := initializers.DB.Where("email = ?", request.Email).Find(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid OTP",
		})
		utils.JSONResponse(c, http.StatusForbidden, nil, "Invalid OTP", result.Error)
		return
	}

	// Cek apakah waktu update OTP sudah lebih dari 5 menit
	if time.Since(user.UpdatedAt) > 5*time.Minute {
		utils.JSONResponse(c, http.StatusForbidden, nil, "OTP has expired", result.Error)
		return
	}

	// Cek apakah OTP yang dimasukkan sesuai dengan OTP yang tersimpan
	if request.OTP != user.Otp {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid OTP",
		})
		utils.JSONResponse(c, http.StatusForbidden, nil, "Invalid OTP", result.Error)
		return
	}

	// siapkan accessToken dan simpan di cookies
	accessToken, err := utils.GenerateAccessToken(int(user.ID))
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to generate token", result.Error)
		return
	}

	c.SetCookie("Authorization", accessToken, 3600*24, "/", "", false, true)

	utils.JSONResponse(c, http.StatusOK, nil, "Access granted", nil)
}

func Logout(c *gin.Context) {
	userReq, _ := c.Get("user")

	var user = userReq.(models.User)

	initializers.DB.Model(&models.User{}).Where("ID = ?", user.ID).Update("otp", nil)

	// Menghapus cookie dengan mengatur waktu kadaluarsa yang sudah lewat
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	utils.JSONResponse(c, http.StatusOK, nil, "Success Logout", nil)
}
