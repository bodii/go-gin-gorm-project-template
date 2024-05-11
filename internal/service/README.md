## example

file name: user.go
```go
package service

import (
	"errors"
	"log/slog"
	"template-project-name/internal/bootstrap"
	"template-project-name/internal/models"
	"template-project-name/internal/types"
)

// UserService UserService struct
type UserService struct {
	DB    *types.MysqlDBT
	Redis *types.RedisDBT
	Log   *slog.Logger
}

// NewUserService : get injected user repo
func NewUserService() UserService {
	db, _ := bootstrap.GetMysqlDB("dbname")
	// redis, _ := bootstrap.GetRedis("one") // config in 'internal/config/redis.toml' file

	return UserService{
		DB: db,
	}
}

// Login -> Gets validated user
func (u UserService) UserLogin(userLogin models.UserLogin) (*models.UserLoginResponse, error) {
	user := &models.User{}
	response := &models.UserLoginResponse{}
	result := u.DB.Where("username = ?", userLogin.Username).First(user)
	if result.Error != nil {
		return response, result.Error // errors.New("用户信息不存在")
	}

	// 当前用户是否可用
	if ok, err := user.IsAvailable(); !ok {
		return response, err
	}

	// 验证密码是否正确
	ok, err := user.Login(userLogin)
	if err != nil {
		return response, errors.New("查询用户信息有误")
	}

	if !ok {
		return response, errors.New("登录失败")
	}

	return response, nil
}

func (u UserService) CheckUser(userInfo models.CheckUser) (*models.UserLoginResponse, error) {
	user := &models.User{}
	response := &models.UserLoginResponse{}
	result := u.DB.Where("userid = ?", userInfo.Userid).First(user)

	// 当前用户是否可用
	if ok, err := user.IsAvailable(); !ok {
		return response, err
	}

	return user.ToLoginResponse(), result.Error
}
```