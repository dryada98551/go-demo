package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DomainCheckMiddleware(allowedDomain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		if host != allowedDomain {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}
		c.Next()
	}
}
