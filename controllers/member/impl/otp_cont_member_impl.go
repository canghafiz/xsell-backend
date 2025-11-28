package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/otp"
	"be/models/services/member"

	"github.com/gin-gonic/gin"
)

type OtpContMemberImpl struct {
	OtpServ member.OtpServMember
}

func NewOtpContMemberImpl(otpServ member.OtpServMember) *OtpContMemberImpl {
	return &OtpContMemberImpl{OtpServ: otpServ}
}

func (cont *OtpContMemberImpl) SendOtpPhoneVerify(context *gin.Context) {
	// Parse Request Body
	request := otp.SendRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.OtpServ.SendOtpToPhoneVerify(request)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *OtpContMemberImpl) CheckOtpPhoneVerify(context *gin.Context) {
	// Parse Request Body
	request := otp.CheckRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.OtpServ.CheckOtpPhoneVerify(request)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}
