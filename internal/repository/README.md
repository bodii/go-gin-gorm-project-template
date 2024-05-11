## example 

file name: user.go
```go
package repository

import (
	"template-project-name/internal/models"
	"template-project-name/internal/types"
	"template-project-name/internal/utils"
)

// UserRepository -> UserRepository resposible for accessing database
type UserRepository struct {
	db *types.MysqlDBT
}

// NewUserRepository -> creates a instance on UserRepository
func NewUserRepository(db *types.MysqlDBT) UserRepository {
	return UserRepository{
		db: db,
	}
}

// CreateUser -> method for saving user to database
func (u UserRepository) CreateUser(user models.UserRegister) error {

	var dbUser models.User
	dbUser.Password = user.Password
	return u.db.Create(&dbUser).Error
}

// LoginUser -> method for returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	password := user.Password

	err := u.db.Where("username = ?", user.Username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := utils.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}

func (u UserRepository) CheckUserExist(user models.CheckUser) (bool, error) {

	var dbUser models.User

	err := u.db.Where("userid = ?", user.Userid).First(&dbUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

```