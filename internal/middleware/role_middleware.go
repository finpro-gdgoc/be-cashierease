package middleware

import (
	"cashierease/internal/models" //
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoleAny, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			return
		}

		userRole, ok := userRoleAny.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid user role format in context"})
			return
		}

		if userRole != string(requiredRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Next()
	}
}