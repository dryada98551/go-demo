package configs

import (
	"github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

type EnvVar struct {
	GIN       EnvGIN   `json:"gin"`
	HealthUrl []string `json:"healthUrl"`
}

type GinSet struct {
	Engine *gin.Engine
}

type Sqlite struct {
	Engine *gorm.DB
}

type EnvGIN struct {
	Mode        string   `json:"mode"`
	Host        string   `json:"host"`
	Port        string   `json:"port"`
	IPWhitelist []string `json:"ipWhitelist"`
}
