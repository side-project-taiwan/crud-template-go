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
	gin_Instance.POST("/CreateNewUser", func(_ginCTX *gin.Context) {
		var userInput struct {
			Account  string `json:"account"`
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}
		if err := _ginCTX.ShouldBindJSON(&userInput); err != nil {
			_ginCTX.String(http.StatusBadRequest, "Invalid request data")
			return
		}

		newUserID, err := _target.UserService.CreateNewUser(userInput)
		if err != nil {
			_ginCTX.String(http.StatusInternalServerError, "Error creating new user")
			return
		}

		_ginCTX.String(http.StatusOK, "New user created with ID: %v", newUserID)
	})

	gin_Instance.GET("/UserController", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----UserController 20240401")
	})

}
