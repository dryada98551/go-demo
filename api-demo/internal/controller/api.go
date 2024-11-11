package route

import (
	"github.com/gin-gonic/gin"

	apiv1 "main/internal/controller/v1"
)

func RegisterRoutesV1(rg *gin.RouterGroup) {
	rg.POST("/path/demo", apiv1.Demo01)
	// rg.POST("/UserSocial/Logout", apiv1.POST02)
	// 其他路由...
}
