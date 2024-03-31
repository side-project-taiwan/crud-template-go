package controller

import (
	"net/http"
	"sample/internal/service"
	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func InitializeUserController(userService *service.UserService, stocksService *service.StocksService, gin_Instance *gin.Engine) *UserController {
	sc := &UserController{
		UserService: userService,
	}
	sc.setAssignRoutes(gin_Instance)
	return sc
}

func (sc *UserController) setAssignRoutes(gin_Instance *gin.Engine) {

	util.PrintLogWithColor("Enter setAssignRoutes")

	gin_Instance.GET("/UserController", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----UserController 20240401")
	})
}
