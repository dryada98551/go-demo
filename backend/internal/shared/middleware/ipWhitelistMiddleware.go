package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckIPInWhitelist 檢查 IP 是否在白名單中
func CheckIPInWhitelist(ip string, ipWhitelist []string) bool {
	for _, whitelistIP := range ipWhitelist {
		if ip == whitelistIP {
			return true
		}
	}
	return false
}

// IPWhitelistMiddleware 是一個基於來源 IP 的白名單中間件
func IPWhitelistMiddleware(ipWhitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 獲取來源 IP
		clientIP := c.ClientIP()

		// 檢查來源 IP 是否在白名單中
		if !CheckIPInWhitelist(clientIP, ipWhitelist) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "來源 IP 被拒絕",
			})
			c.Abort()
			return
		}

		// 如果 IP 在白名單中，繼續執行後續處理器
		c.Next()
	}
}
