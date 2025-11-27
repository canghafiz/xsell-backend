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
	validSections := map[string]bool{
		"newly_listed":     true,
		"trending":         true,
		"like_new_items":   true,
		"good_condition":   true,
		"used_electronics": true,
		"home_items":       true,
		"fashion_items":    true,
		"editor_pick":      true,
	}

	sectionLimits := make(map[string]int)
	
	for key, values := range context.Request.URL.Query() {
		if !validSections[key] {
			continue
		}

		if len(values) == 0 {
			continue
		}

		limit, err := strconv.Atoi(values[0])
		if err != nil || limit <= 0 {
			continue
		}

		sectionLimits[key] = limit
	}

	// Call Serv
	result, errServ := cont.PageLayoutServ.GetHomeLayouts(sectionLimits)
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
	validSections := map[string]bool{
		"trending":       true,
		"like_new_items": true,
	}

	sectionLimits := make(map[string]int)

	for key, values := range context.Request.URL.Query() {
		if !validSections[key] {
			continue
		}

		if len(values) == 0 {
			continue
		}

		limit, err := strconv.Atoi(values[0])
		if err != nil || limit <= 0 {
			continue
		}

		sectionLimits[key] = limit
	}

	// Call Serv
	result, errServ := cont.PageLayoutServ.GetProductDetailLayouts(sectionLimits)
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
