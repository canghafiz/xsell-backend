package impl

import (
	"be/exception"
	"be/helper"
	"be/models/requests/user"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserContMemberImpl struct {
	UserService member.UserServMember
}

func NewUserContMemberImpl(userService member.UserServMember) *UserContMemberImpl {
	return &UserContMemberImpl{UserService: userService}
}

func (cont *UserContMemberImpl) Register(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.UserService.Register(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helper.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helper.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserContMemberImpl) Login(context *gin.Context) {
	// Parse Request Body
	request := user.LoginRequest{}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	result, errServ := cont.UserService.Login(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helper.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	errResponse := helper.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserContMemberImpl) Logout(context *gin.Context) {
	email := context.Param("email")

	// Call Service
	errServ := cont.UserService.Logout(email)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helper.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helper.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserContMemberImpl) UpdateData(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// Parse Request Body
	request := user.UpdateDataRequest{
		UserId: userId,
	}
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	result, errServ := cont.UserService.UpdateData(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helper.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	errResponse := helper.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}
