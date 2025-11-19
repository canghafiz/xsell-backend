package middlewares

import (
	"be/helper"
	"be/models/domains"
	"be/models/repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, userRepo repositories.UserRepo, jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		// 1. Check Auth Exists
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header required",
			})
			return
		}

		// 2. Check Format Bearer
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			return
		}

		// 3. Extract token from header
		tokenString := strings.TrimPrefix(header, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is required",
			})
			return
		}

		// 4. Decode Token
		result, errDecode := helper.DecodeJWT(tokenString, jwtKey)
		if errDecode != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			return
		}

		// 5. Extract email with type safety
		email, ok := result["email"].(string)
		if !ok || email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			return
		}

		// 6. Check Token in DB
		tokenModel := domains.Users{
			Email: email,
			Token: &tokenString,
		}
		tokenValid := userRepo.CheckTokenValid(db, tokenModel)
		if !tokenValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or revoked",
			})
			return
		}

		// 7. Set user data to context
		c.Set("user_email", email)
		if userID, exists := result["user_id"]; exists {
			c.Set("user_id", userID)
		}

		c.Next()
	}
}
