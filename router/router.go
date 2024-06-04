package router

import (
	"cemeta-resource/config"
	"cemeta-resource/controller"
	"cemeta-resource/lib"
	"cemeta-resource/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Api ApiGroup
}

var Controller = new(controller.Controller)

func New() *gin.Engine {
	if !lib.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// middleware
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(
		middleware.Cors(),
	)

	r := new(Router)

	{
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("OK, Version: %s", config.Get(config.Version)))
		})
		apiGroup := router.Group(config.Get(config.ApiPathPrefix))
		apiGroup.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("OK, Version: %s", config.Get(config.Version)))
		})
		r.Api.UserRouter.Init(apiGroup)
	}

	return router
}
