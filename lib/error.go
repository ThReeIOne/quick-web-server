package lib

import "github.com/gin-gonic/gin"

var (
	OK                  = gin.H{"message": "成功"}
	Unauthorized        = gin.H{"message": "未登录或非法访问"}
	InternalServerError = gin.H{"message": "系统错误"}
	BadRequest          = gin.H{"message": "参数错误"}
	EmptyResponse       = gin.H{}
)
