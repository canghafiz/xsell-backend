package app

import (
	"be/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	// Member
	MemberDependency *MemberDependency

	JwtKey string
	Engine *gin.Engine
}

func NewRouter(r Router) *Router {
	// Member Route
	memberGroup := r.Engine.Group("api/v1/member")
	{
		userGroup := memberGroup.Group("/user")
		{
			userGroup.POST("/register", r.MemberDependency.UserCont.Register)
			userGroup.POST("/login", r.MemberDependency.UserCont.Login)

			authMiddleware := userGroup.Use(middlewares.AuthMiddleware(r.MemberDependency.Db, r.MemberDependency.UserRepo, r.JwtKey))
			{
				authMiddleware.PUT("/:userId", r.MemberDependency.UserCont.UpdateData)
				authMiddleware.DELETE("/logout/:email", r.MemberDependency.UserCont.Logout)
			}
		}
	}

	return &Router{
		MemberDependency: r.MemberDependency,

		JwtKey: r.JwtKey,
		Engine: r.Engine,
	}
}
