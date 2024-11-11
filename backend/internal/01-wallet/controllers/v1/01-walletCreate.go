package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"main/internal/01-wallet/service"
	"main/internal/repositories"
	httpStruct "main/internal/shared/utils/http"
)

type payload struct {
	WalletAddress string `json:"walletAddress"`
	FirstToken    []byte `json:"firstToken"`
	SecondToken   []byte `json:"secondToken"`
}

// type response struct {
// 	Url string `json:"Url"`
// }

// @Summary Wallet Create
// @Schemes
// @Description do user Wallet Create
// @Tags V1
// @Accept json
// @Produce json
// @Param   connection_token  formData  string  true  "connection_token"
// @Param X-Custom-Header1 header string true "自定义头信息1"
// @Param user body httpStruct.Request true "用戶信息"
// @Success 200 {string} Response01
// @Failure 400 {object} object
// @Router /api/v1/wallet/create [post]
func CreateWallet(c *gin.Context) {
	var res httpStruct.Response
	var payload payload

	// create web3 wallet
	walletPrivateKey, walletPublicKey, walletAddress, err := service.CreateWallet()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create wallet fail"})
		return
	}
	payload.WalletAddress = walletAddress

	// encrypt key
	data, err := service.Encrypt(&walletPrivateKey, &walletPublicKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "encrypt key fail"})
		return
	}
	payload.FirstToken = data.FirstToken
	payload.SecondToken = data.SecondToken

	userID := uuid.New()

	// save key
	// 創建 UserWallet 控制器
	userWalletController := repositories.NewUserWalletRepository()
	err = userWalletController.Create(&userID, &walletAddress, &data)

	// return response
	res.Payload = payload
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api fail"})
		return
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}
