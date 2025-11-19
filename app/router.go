package app

import "github.com/gin-gonic/gin"

type Router struct {
	// Member
	MemberDependency MemberDependency

	JwtKey string
	Engine *gin.Engine
}
