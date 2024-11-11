package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 記錄每個請求的處理時間
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 處理請求
		c.Next()

		// 請求完成後，記錄時間
		duration := time.Since(start)
		log.Printf("請求 [%s] %s 花費 %v", c.Request.Method, c.Request.URL, duration)
	}
}
