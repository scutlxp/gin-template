package v1

import (
	"gin-project/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(g *gin.RouterGroup) {
	userController := controller.NewUserController()

	// todo 添加路由
	g.GET("/users", userController.GetUsers)
}
