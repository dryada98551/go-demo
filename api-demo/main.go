package main

import (

	docs "main/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	route "main/internal/controller"
	configs "main/internal/lib/init"
	// "fmt"
)

func main() {
	//=============================================================
	configs.Init(true, false, false)
	swaggerInfoSetting()

	// group: v1
	v1 := configs.Global.GinSet.Engine.Group(configs.Global.EnvVar.GIN.BASEPATH + "/api/v1")
	{
		route.RegisterRoutesV1(v1)
	}

	configs.Global.GinSet.Engine.Run(":" + configs.Global.EnvVar.GIN.PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func swaggerInfoSetting() {
	docs.SwaggerInfo.Title = "Hello World API"
	docs.SwaggerInfo.Description = "Hello World API 範例"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = configs.Global.EnvVar.GIN.HOST + ":" + configs.Global.EnvVar.GIN.PORT
	docs.SwaggerInfo.BasePath = configs.Global.EnvVar.GIN.BASEPATH
	// 使用 Swaggo 提供的 Swagger 服務
	configs.Global.GinSet.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}