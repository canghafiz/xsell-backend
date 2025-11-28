package controllers

import "github.com/gin-gonic/gin"

type OtpCont interface {
	SendEmailVerification(context *gin.Context)
	CheckOtp(context *gin.Context)
}
