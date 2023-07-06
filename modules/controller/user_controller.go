package controller

import (
	"fmt"
	"net/http"
	"restapi/initializers"
	"restapi/modules/model"
	"restapi/modules/repository"
	"restapi/modules/request"
	"restapi/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Mendapatkan data dari body request
	var request request.FormEmail

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	userRepo := repository.NewUserRepository(initializers.DB)

	user, err := userRepo.FindByEmail(request.Email)

	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Data Not Found", err)
		return
	}

	otp := utils.GenerateOTP()

	err = userRepo.UpdateOtp(user.ID, otp)

	fmt.Println(otp)

	utils.JSONResponse(c, http.StatusOK, user.GetUser(), "OTP has been sent to your email", nil)
}

func OTP(c *gin.Context) {
	// Mendapatkan Otp
	var request request.FormUser

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, nil, "Invalid request body", err)
		return
	}

	// Find User By Email
	userRepo := repository.NewUserRepository(initializers.DB)

	user, err := userRepo.FindByEmail(request.Email)

	// validasi Otp, waktu dan sesuai atau tidak
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, nil, "Email not found", err)
		return
	}

	// Cek apakah waktu update OTP sudah lebih dari 5 menit
	if time.Since(user.UpdatedAt) > 5*time.Minute {
		utils.JSONResponse(c, http.StatusForbidden, nil, "OTP has expired", err)
		return
	}

	// Cek apakah OTP yang dimasukkan sesuai dengan OTP yang tersimpan
	if request.OTP != user.Otp {
		utils.JSONResponse(c, http.StatusForbidden, nil, "Invalid OTP", err)
		return
	}

	// siapkan accessToken dan simpan di cookies
	accessToken, errToken := utils.GenerateAccessToken(int(user.ID))
	if errToken != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, nil, "Failed to generate token", errToken)
		return
	}

	c.SetCookie("Authorization", accessToken, 3600*24, "/", "", false, true)

	utils.JSONResponse(c, http.StatusOK, nil, "Access granted", nil)
}

func Logout(c *gin.Context) {
	userReq, _ := c.Get("user")

	var user = userReq.(model.User)
	userRepo := repository.NewUserRepository(initializers.DB)

	// Menghapus catatan Otp yang tersimpan
	userRepo.UpdateResetOtp(user.ID)

	// Menghapus cookie dengan mengatur waktu kadaluarsa yang sudah lewat
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	utils.JSONResponse(c, http.StatusOK, nil, "Success Logout", nil)
}
