package configs

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "main/internal/shared/utils/log"

	"main/internal/shared/middleware"
	"main/internal/models"
)

type InitParmsConfig struct {
	EnvVar EnvVar
	GinSet GinSet
	Sqlite Sqlite
}

var Global InitParmsConfig

func (param InitParmsConfig) JsonSetting() InitParmsConfig {
	GIN_MODE := os.Getenv("GIN_MODE")

	switch GIN_MODE {
	case "release":
		gin.SetMode(gin.ReleaseMode)
		// param.HttpUrl = libhttp.HttpUrlRelease
	case "uat":
		gin.SetMode(gin.TestMode)
		// param.HttpUrl = libhttp.HttpUrlUat
	case "debug":
		gin.SetMode(gin.DebugMode)
		// param.HttpUrl = libhttp.HttpUrlDebug
	default:
		// 如果没有设置环境变量，默认为 debug 模式
		gin.SetMode(gin.DebugMode)
		GIN_MODE = "debug"
		// param.HttpUrl = libhttp.HttpUrlDebug
	}

	path, _ := os.Getwd()
	file, err := os.Open(path + "/config/config-" + GIN_MODE + ".json")
	if err != nil {
		log.ERROR("config file not fund")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	defer file.Close()

	// 創建 JSON 解碼器
	decoder := json.NewDecoder(file)

	// 解碼 JSON
	err = decoder.Decode(&param.EnvVar)
	if err != nil {
		log.ERROR("Error decoding JSON")
		os.Exit(1) // 解碼出錯，提前終止函數
	}
	param.EnvVar.GIN.Mode = GIN_MODE
	return param
}

func (param InitParmsConfig) GinSetting() InitParmsConfig {
	param.GinSet.Engine = gin.Default()
	// 於此可定義Middleware
	param.GinSet.Engine.Use(middleware.LoggerMiddleware())
	param.GinSet.Engine.Use(middleware.IPWhitelistMiddleware(param.EnvVar.GIN.IPWhitelist))
	param.GinSet.Engine.Use(middleware.CORSMiddleware())
	param.GinSet.Engine.Use(middleware.RequestCheckMiddleware())
	// // 使用domainCheckMiddleware中間件
	// if os.Getenv("GIN_MODE") == "release" {
	// 	param.GinSet.Engine.Use(middleware.DomainCheckMiddleware(param.EnvVar.GIN.HOST))
	// }

	return param
}

func (param InitParmsConfig) SqliteSetting() InitParmsConfig {
	// 初始化 SQLite 資料庫連接，使用 GORM
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.ERROR("無法打開 SQLite 資料庫:" + err.Error())
	}
	

	// 自動遷移用戶表結構
	db.AutoMigrate(&models.UserWallet{})

	param.Sqlite.Engine = db

	return param
}

func (param InitParmsConfig) LivenessSet() {
	param.GinSet.Engine.GET("/_liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}

func Init() {
	Global = Global.JsonSetting()
	Global = Global.GinSetting()
	Global = Global.SqliteSetting()
	Global.LivenessSet()
}
