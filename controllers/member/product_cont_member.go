package member

import "github.com/gin-gonic/gin"

type ProductContMember interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetSingleBySlug(context *gin.Context)
	Delete(context *gin.Context)
}
