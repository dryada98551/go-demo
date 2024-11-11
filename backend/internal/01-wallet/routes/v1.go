package route

import (
	"github.com/gin-gonic/gin"

	apiv1 "main/internal/01-wallet/controllers/v1"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	rg.POST("/wallet/create", apiv1.CreateWallet)
	// 其他路由...
}
