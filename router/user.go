package router

type UserRouter struct{}

func (*UserRouter) Init(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", Controller.Login)
		userGroup.GET("/channel/token", middleware.CheckChannel(), Controller.GetChannelToken)
	}
}
