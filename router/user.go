package router

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (*UserRouter) Init(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", Controller.Login)
	}
}
