## example

file name: user.go
```go
package handlers

import (
	"template-project-name/internal/models"
	"template-project-name/internal/service"
	"template-project-name/internal/types"
	"template-project-name/internal/utils"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// userController struct
type userController struct {
	service service.UserService
	Redis   *types.RedisDBT
	Log     *types.AppLogT
}

// NewUserController : userController
func NewUserController() *userController {
	return &userController{
		service: service.NewUserService(),
	}
}

// Login is a user login to userController
// http :8080/api/v1/user/login username=aa password=bb
func (u *userController) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	dbUser, err := u.service.UserLogin(user)
	u.Log.Info("user info: %#v", dbUser)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	// jwt generale token
	jwtService := utils.NewJwtGeneraleService()
	dbUser.Token = jwtService.GenerateToken(dbUser.Userid)

	utils.OkJSON(c, dbUser)
}

func (u *userController) CheckUser(c *gin.Context) {

	// see: internal/server/middlewares/jwt.go 
	// context set userid
	userid := c.Get("userid").(uint64);
	if userid < 1 {
		utils.ErrorJSON(c, http.StatusUnauthorized, err.Error())
		return
	}

	UserLoginRes, err := u.service.CheckUser(models.CheckUser{Userid: userid})
	if err != nil {
		utils.ErrorJSON(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.OkJSON(c, UserLoginRes)
}
```