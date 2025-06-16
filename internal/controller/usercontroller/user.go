package usercontroller

import (
	"gin-project/internal/service/userservice"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var service = userservice.Get()

// GetUsers
// @Summary 获取所有用户
// @Description 获取系统中的所有用户
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := service.GetUsers(ctx)
	if err != nil {
		//
	}
	ctx.JSON(200, users)
}
