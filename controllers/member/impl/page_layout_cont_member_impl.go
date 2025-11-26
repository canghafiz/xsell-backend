package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageLayoutContMemberImpl struct {
	PageLayoutServ member.PageLayoutServMember
}

func NewPageLayoutContMemberImpl(pageLayoutServ member.PageLayoutServMember) *PageLayoutContMemberImpl {
	return &PageLayoutContMemberImpl{PageLayoutServ: pageLayoutServ}
}

func (cont *PageLayoutContMemberImpl) GetHomeLayouts(context *gin.Context) {
	limitStr := context.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	// Call Service
	result, errServ := cont.PageLayoutServ.GetHomeLayouts(limit)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *PageLayoutContMemberImpl) GetProductDetailLayouts(context *gin.Context) {
	limitStr := context.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	// Call Service
	result, errServ := cont.PageLayoutServ.GetHomeLayouts(limit)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}
