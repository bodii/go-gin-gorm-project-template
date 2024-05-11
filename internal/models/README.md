## example

file name: user.go
```go
package models

import (
	"errors"
	"strings"
	"template-project-name/internal/utils"
)

// User -> User struct to save user on database
type User struct {
	Userid      int64   `gorm:"column:userid" json:"userid" form:"userid"`
	Username    string  `gorm:"column:username" json:"username" form:"username"`
	Password    string  `gorm:"column:password" json:"password" form:"password"`
	Nickname    string  `gorm:"column:nickname" json:"nickname" form:"nickname"`
	Email       string  `gorm:"column:email" json:"email" form:"email"`
	Phone       string  `gorm:"column:phone" json:"phone" form:"phone"`
}

// TableName -> returns the table name of User Model
func (user *User) TableName() string {
	return "users"
}

type CheckUser struct {
	Userid string `form:"userid" binding:"required"`
}

// UserLogin -> Request Binding for User Login
type UserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserLoginResponse struct {
	Userid    int64   `gorm:"column:userid" json:"userid" form:"userid"`
	Nickname  string  `gorm:"column:nickname" json:"nickname" form:"nickname"`
}

// UserRegister -> Request Binding for User Register
type UserRegister struct {
	Username    string `form:"username" json:"username" binding:"required"`
	Phone       string `form:"phone" json:"phone" binding:"required"`
	Nickname    string `form:"nickname" json:"nickname" binding:"required"`
    Password    string `form:"password" json:"password" binding:"required"`
    Salt        string `form:"salt" json:"salt"`
    Status      int    `form:"status" json:"status"`
}

// ResponseMap -> response map method of User
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["userid"] = user.Userid
	resp["username"] = user.Username
	resp["nickname"] = user.Nickname
	return resp
}

// 当前用户是否可用
func (user *User) IsAvailable() (bool, error) {
	return user.Status == 1, errors.New("当前用户不可用")
}

// 登录
func (user *User) Login(login UserLogin) (bool, error) {
	if user.Username == "" {
		return false, errors.New("数据库未查询到username字段的值")
	}

	if user.Username != login.Username {
		return false, nil
	}

	return user.CheckLoginPassword(login.Password)
}

// 验证密码
func (user *User) CheckLoginPassword(inputPassword string) (bool, error) {
	if user.Salt == "" {
		return false, errors.New("数据库未查询到salt字段的值")
	}

	if user.Password == "" {
		return false, errors.New("数据库未查询到password字段的值")
	}

	encodeInputPwd := utils.MD5EncodePasswd(inputPassword, user.Salt)

	result := user.Password == encodeInputPwd

	// log.Printf("input: %s, p: %s", encodeInputPwd, user.Password)

	return result, nil
}

// user to userOtherType
func (u *User) ToLoginResponse() *UserLoginResponse {
	return &UserLoginResponse{
		Userid:    u.Userid,
		Nickname:  u.Nickname,
	}
}

```