package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/product"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductContMemberImpl struct {
	ProductServ member.ProductServMember
}

func NewProductContMemberImpl(productServ member.ProductServMember) *ProductContMemberImpl {
	return &ProductContMemberImpl{ProductServ: productServ}
}

func (cont *ProductContMemberImpl) Create(context *gin.Context) {
	// Parse Request Body
	request := product.CreateProductRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.ProductServ.Create(request)
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

func (cont *ProductContMemberImpl) Update(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Parse Request Body
	request := product.UpdateProductRequest{
		ProductId: productId,
	}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.ProductServ.Update(request)
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

func (cont *ProductContMemberImpl) GetSingleBySlug(context *gin.Context) {
	productSlug := context.Param("productSlug")

	// Call Service
	result, errServ := cont.ProductServ.GetSingleBySlug(productSlug)
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

func (cont *ProductContMemberImpl) Delete(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Call Service
	errServ := cont.ProductServ.Delete(productId)
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
