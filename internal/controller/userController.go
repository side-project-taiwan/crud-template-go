package controller

import (
	"net/http"
	"sample/internal/repository/model"
	"sample/internal/service"
	"sample/internal/util"
	"time"

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
		newUserData := &model.User{
			Account:   userInput.Account,
			Username:  userInput.Username,
			Password:  userInput.Password,
			Email:     userInput.Email,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newUserID, err := _target.UserService.CreateNewUser(newUserData)
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
