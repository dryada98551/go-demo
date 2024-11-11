package httpStruct

// import (
// 	"encoding/json"
// )

type Request struct {
	APIInfo  interface{} `json:"apiInfo"`
	UserInfo interface{} `json:"userInfo"`
	Payload  interface{} `json:"payload"`
}

type Response struct {
	APIInfo  interface{} `json:"apiInfo"`
	UserInfo interface{} `json:"userInfo"`
	Payload  interface{} `json:"payload"`
}
