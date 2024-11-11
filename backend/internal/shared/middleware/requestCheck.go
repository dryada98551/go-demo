package middleware

import (
	"net/http"
	// "encoding/json"
	"github.com/gin-gonic/gin"

	"main/internal/shared/utils/http"
)

// RequestCheckMiddleware 檢查 JSON 格式和必須字段
func RequestCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req httpStruct.Request

		// 解析請求中的 JSON 數據
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			c.Abort()
			return
		}

		// 檢查是否存在必須的字段
		if req.APIInfo == nil || req.UserInfo == nil || req.Payload == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 若檢查通過，繼續處理後續的 handler
		c.Next()
	}
}

