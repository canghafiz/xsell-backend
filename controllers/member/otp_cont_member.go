package member

import "github.com/gin-gonic/gin"

type OtpContMember interface {
	SendOtpPhoneVerify(context *gin.Context)
	CheckOtpPhoneVerify(context *gin.Context)
}
