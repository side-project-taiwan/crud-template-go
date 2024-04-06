package repository

import (
	"sample/internal/repository/model"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	Database *gorm.DB
}

func NewUserRepositoryGorm(_database *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{
		Database: _database,
	}
}
func (r *UserRepositoryGorm) CreateNewUser(user *model.User) (int32, error) {
	result := r.Database.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}

	// 獲取創建後的 ID
	userID := user.UserID

	return userID, nil
}

// func (r *UserRepositoryGorm) GetUserByID(userID int32) (*model.User, error) {
// 	user := &model.User{}
// 	result := r.Database.First(user, userID)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return user, nil
// }

// func (r *UserRepositoryGorm) UpdateUser(user *model.User) error {
// 	result := r.Database.Save(user)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (r *UserRepositoryGorm) DeleteUser(userID int32) error {
// 	result := r.Database.Delete(&model.User{}, userID)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
