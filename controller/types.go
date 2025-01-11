package controller

import (
	"github.com/gin-gonic/gin"
	"quick_web_golang/lib"
	"quick_web_golang/service"
)

type Controller struct {
	User
}

var Service = new(service.Service)

func GetUid(c *gin.Context) int {
	if claims, ok := c.Get(lib.GinCtxKeyClaims); ok {
		return claims.(*lib.CustomClaims).BaseClaims.Uid
	}
	return 0
}

func GetCid(c *gin.Context) int {
	if claims, ok := c.Get(lib.GinCtxKeyClaims); ok {
		return claims.(*lib.CustomClaims).BaseClaims.CompanyId
	}
	return 0
}
