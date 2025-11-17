package controllers

import "github.com/gin-gonic/gin"

type UserCont interface {
	RegisterMember(context *gin.Context)
	LoginMember(context *gin.Context)
	Logout(context *gin.Context)
	UpdateData(context *gin.Context)
}
