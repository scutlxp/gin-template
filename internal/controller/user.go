package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

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
}
