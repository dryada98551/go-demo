package main

import (
	// "fmt"
	// "main/internal/pkg"
	// "main/internal/service"
	// "math/big"
	// "time"
	configs "main/internal/shared/utils/init"
	route01 "main/internal/01-wallet/routes"
)

func main() {
	//=============================================================
	configs.Init()
	// swaggerInfoSetting()

	// group: v1
	v1 := configs.Global.GinSet.Engine.Group("/api/v1")
	{
		route01.RegisterRoutesV1(v1)
	}

	
	// v2 := configs.Global.GinSet.Engine.Group("/api/v2")
	// {
	// 	route01.RegisterRoutesV1(v2)
	// }

	configs.Global.GinSet.Engine.Run(":" + configs.Global.EnvVar.GIN.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func swaggerInfoSetting() {
// 	docs.SwaggerInfo.Title = "Game Login API"
// 	docs.SwaggerInfo.Description = "Game Login API"
// 	docs.SwaggerInfo.Version = "1.0"
// 	docs.SwaggerInfo.Host = configs.Global.EnvVar.GIN.Host + ":" + configs.Global.EnvVar.GIN.Port
// 	docs.SwaggerInfo.BasePath = "/"
// 	// 使用 Swaggo 提供的 Swagger 服務
// 	// configs.Global.GinSet.Engine.GET(configs.Global.EnvVar.GIN.BASEPATH + "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// 	configs.Global.GinSet.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// }

// func main() {
// 	start := time.Now()
// 	walletPrivateKey, walletPublicKey, walletAddress := pkg.CreateWallet()
// 	fmt.Println("wallet privateKey: ", walletPrivateKey)
// 	fmt.Println("wallet publicKey: ", walletPublicKey)
// 	fmt.Println("wallet address: ", walletAddress)

// 	step1 := time.Now()
// 	fmt.Println("wallet spend: ", step1.Sub(start))

// 	data, _ := service.Encrypt(walletPrivateKey, walletPublicKey)
// 	fmt.Printf("FirstToken: %x\n", data.FirstToken)
// 	fmt.Printf("SecondToken: %x\n", data.SecondToken)
// 	step2 := time.Now()
// 	fmt.Println("Encrypt spend: ", step2.Sub(step1))

// 	decryptwalletPrivateKey, _ := service.Decrypt(data)
// 	fmt.Println("decrypt wallet privateKey: ", decryptwalletPrivateKey)
// 	step3 := time.Now()
// 	fmt.Println("Decrypt spend: ", step3.Sub(step2))

// 	// 取得 BTC/USD 價格 (Sepolia鏈)
// 	contractAddress := "0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43"

// 	lastPrice, _ := pkg.ChainlinkPriceFeed(contractAddress)

// 	fmt.Println("lastPrice: ", lastPrice)

// 	// mint USDT 測試交易
// 	walletPrivateKeyTest := "b7772e5de15f0acad85690b15a018d1cb41c8178db3cf32e2aa67634ad3f234e"
// 	walletAddrTest := "0x90665861799f25CBEe92cB2Ad2B0A7C661376b1c"
// 	usdtContractAddr := "0xaA8E23Fb1079EA71e0a56F48a2aA51851D8433D0"
// 	amount := new(big.Int)
// 	amount.Mul(big.NewInt(100), big.NewInt(1e6)) // 100 USDT, 因為 decimal 是 6

// 	txHash, err := pkg.MintUSDT(usdtContractAddr, walletPrivateKeyTest, walletAddrTest, amount)
// 	if err != nil {
// 		fmt.Println("err : ", err)
// 	}

// 	fmt.Println("txHash: ", txHash)

// }
