package service

import (
	"sample/internal/repository/model"
	"sample/internal/util"
	"strconv"
)

type UserService struct {
	structRepository i_UserRepository
}

type i_UserRepository interface {
	CreateNewUser(data *model.User) (*model.User, error)
}

func NewUserService(userRepo i_UserRepository) *UserService {
	util.PrintLogWithColor("Enter stockService", "#00ffff")

	callBack := UserService{
		structRepository: userRepo,
	}
	return &callBack
}

func (_target *UserService) CreateNewUser(newUser *model.User) (string, error) {
	// randomNum := rand.Intn(10000)
	// randomNumStr := strconv.Itoa(randomNum)

	// user := &model.User{
	// 	//UserID:    1,
	// 	Account:   "example_account" + randomNumStr,
	// 	Username:  "example_username",
	// 	Password:  "example_password",
	// 	Email:     "example@example.com",
	// 	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// 	UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// }
	newUser, err := _target.structRepository.CreateNewUser(newUser)
	if err != nil {
		return "", err
	}
	userIDStr := strconv.FormatInt(int64(newUser.UserID), 10)
	return userIDStr, nil

}
