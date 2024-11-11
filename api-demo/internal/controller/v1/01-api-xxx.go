package apiv1

import (
	"net/http"
	"time"
	"context"

	"github.com/gin-gonic/gin"

	lib "main/internal/lib"
	libmodel "main/internal/lib/model"
	configs "main/internal/lib/init"

	pb "main/internal/lib/jwt"
	service01 "main/internal/service/01-service"
	libreturncode "main/internal/lib/returnCode"
	// GLogging "main/internal/lib/gcplog"
)

// type Request01 struct {
// 	Factory string `json:"factory"`
// 	GameId  string `json:"gameId"`
// }

type Request01 struct {
	// Userid  string `json:"userid"`
	Factory string `json:"factory"`
	GameId  string `json:"gameId"`
}
type Response01 struct {
	Url string `json:"Url"`
}

// @Summary game galaxyagroup login
// @Schemes
// @Description do user game login
// @Tags V1
// @Accept json
// @Produce json
// @Param   connection_token  formData  string  true  "connection_token"
// @Param X-Custom-Header1 header string true "自定义头信息1"
// @Param user body Request01 true "用戶信息"
// @Success 200 {string} Response01
// @Failure 400 {object} object
// @Router /api/v1/game/galaxyagroup/start [post]
func Demo01(c *gin.Context) {

	// start time
	start := time.Now()

	// header check
	_, err := lib.HeaderCheck(c)
	if err != nil {
		res := libmodel.Response{
			Code: libreturncode.Cobelib.ServiceFail.ApiHeader.Code,
			Message: libreturncode.Cobelib.ServiceFail.ApiHeader.Message,
			RequestedTime: start,
			ResponsedTime: time.Now(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Grpc 解JWTToken拿取UserID
	validateResponse, err := configs.Global.GrpcConn.Client.Validate(context.Background(), &pb.ValidateRequest{Token: c.GetHeader("JWTToken"), Loginat: c.GetHeader("LoginAt")})
	if err != nil {
		res := libmodel.Response{
			Code: libreturncode.Cobelib.ServiceFail.JwtInput.Code,
			Message: libreturncode.Cobelib.ServiceFail.JwtInput.Message,
			RequestedTime: start,
			ResponsedTime: time.Now(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	UserID := validateResponse.Userid

	// 檢查Request
	var req Request01
	if err := c.ShouldBindJSON(&req); err != nil {
		res := libmodel.Response{
			Code: libreturncode.Cobelib.ServiceFail.DataParse.Code,
			Message: libreturncode.Cobelib.ServiceFail.DataParse.Message,
			RequestedTime: start,
			ResponsedTime: time.Now(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// UserID := req.Userid

	// 查看登入狀況
	sysReturn := lib.LoginCheck(c.Request.Context(), configs.Global.EnvVar.JWT.JWTREDISKEY, UserID)
	if sysReturn.Code != libreturncode.Cobelib.Success.Other.Code {
		res := libmodel.Response{
			Code: sysReturn.Code,
			Message: sysReturn.Message,
			RequestedTime: start,
			ResponsedTime: time.Now(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// ...

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api fail"})
		return
	} else {
		res := libmodel.Response{
			Code: 200,
			Data: Response01{
				// Url:  url,
			},
			RequestedTime: start,
			ResponsedTime: time.Now(),
		}
		c.JSON(http.StatusOK, res)
		return
	}
}
