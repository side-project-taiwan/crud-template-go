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

func InitializeUserController(userService *service.UserService, gin_Instance *gin.Engine) *UserController {
	_target := &UserController{
		UserService: userService,
	}
	_target.setAssignRoutes(gin_Instance)
	return _target
}

func (_target *UserController) setAssignRoutes(gin_Instance *gin.Engine) {

	util.PrintLogWithColor("Enter setAssignRoutes")
	//userService := _target.UserService
	gin_Instance.GET("/CreateNewUser", func(_ginCTX *gin.Context) {
		newUserID, err := _target.UserService.CreateNewUser()
		if err != nil {
			_ginCTX.String(http.StatusInternalServerError, "Error creating new user")
			return
		}

		_ginCTX.String(http.StatusOK, "New user created: %v", newUserID)
	})

	gin_Instance.GET("/UserController", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----UserController 20240401")
	})

}
