package apps

import (
	"be/dependencies"
	"be/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	AdminDependency  *dependencies.AdminDependency
	MemberDependency *dependencies.MemberDependency

	JwtKey string
	Engine *gin.Engine
}

func NewRouter(r Router) *Router {
	// Admin
	adminGroup := r.Engine.Group("api/v1/admin")
	{
		categoryGroup := adminGroup.Group("/category")
		{
			categoryGroup.POST("/", r.AdminDependency.CategoryCont.Create)
			categoryGroup.PUT("/:categoryId", r.AdminDependency.CategoryCont.Update)
			categoryGroup.DELETE("/:categoryId", r.AdminDependency.CategoryCont.Delete)
		}
	}

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
		AdminDependency:  r.AdminDependency,
		MemberDependency: r.MemberDependency,

		JwtKey: r.JwtKey,
		Engine: r.Engine,
	}
}
