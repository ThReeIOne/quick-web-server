package network

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"quick_web_golang/config"
	"quick_web_golang/log"
	"quick_web_golang/router"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 30 * time.Minute
	s.WriteTimeout = 30 * time.Minute
	s.MaxHeaderBytes = 1 << 20
	return s
}

func Run() {
	r := router.New()
	s := initServer(config.Get(config.GatewayAddress), r)
	if e := s.ListenAndServe(); e != nil {
		log.Error(e.Error())
	}
}
