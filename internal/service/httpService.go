package service

import (
	"sample/internal/configs"

	"github.com/gin-gonic/gin"
)

type httpService struct {
	Gin_Instance *gin.Engine
}

var (
	HttpService *httpService
)

func init() {
	HttpService = &httpService{
		Gin_Instance: gin.Default(),
	}

	if err := HttpService.Gin_Instance.Run(":" + configs.PORT); err != nil {
		panic(err)
	}
}

func GetHttpServiceInstance() *httpService {
	return HttpService
}
