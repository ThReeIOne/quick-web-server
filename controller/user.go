package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick_web_golang/lib"
	"quick_web_golang/log"
	"quick_web_golang/model"
	"regexp"
)

//Note：这里涉及到登录注册的，但是相关参数和流程目前不完整，根据具体业务进行更改就好
//如果采用http，涉及RPC的部分可以去掉

type User struct{}

func (*User) Create(c *gin.Context) {
	input := struct {
		DepartmentId int    `json:"departmentId" binding:"required"`
		RoleId       int    `json:"roleId" binding:"required"`
		Name         string `json:"name" binding:"required"`
		Phone        string `json:"phone" binding:"required"`
		Avatar       string `json:"avatar"`
		Description  string `json:"description"`
	}{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := lib.Verify(input, CommonVerify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	cid := GetCid(c)
	if exist, err := Service.UserService.ExistPhone(cid, input.Phone); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, lib.InternalServerError)
		return
	} else if exist {
		c.JSON(http.StatusBadRequest, gin.H{"message": "手机号已存在"})
		return
	}

	if err := Service.UserService.CreateUser(&model.User{
		Phone:       input.Phone,
		Avatar:      input.Avatar,
		Password:    lib.MD5(lib.Token(6)),
		Description: input.Description,
	}); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, lib.InternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (*User) Login(c *gin.Context) {
	input := struct {
		Phone    string `json:"phone" binding:"required"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// 使用正则表达式匹配手机号格式
	phonePattern := regexp.MustCompile(`^1[3456789]\d{9}$`)
	if !phonePattern.MatchString(input.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "手机号格式不正确"})
		return
	}

	var user *model.User
	var err error

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "用户名或密码错误"})
		return
	}

	// 生成 JWT Token
	jwt := lib.NewJWT()
	token, err := jwt.CreateToken(jwt.CreateClaims(lib.BaseClaims{
		Uid: user.Id,
		//CompanyId:    uc.CompanyId,
		//DepartmentId: uc.DepartmentId,
		Phone: user.Phone,
	}))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, lib.InternalServerError)
		return
	}

	go Service.UpdateLastLoginAt(user.Id)

	c.JSON(http.StatusOK, struct {
		Id           int    `json:"id"`
		Phone        string `json:"phone"`
		CompanyId    int    `json:"companyId"`
		DepartmentId int    `json:"departmentId"`
		Token        string `json:"token"`
	}{
		Id:    user.Id,
		Phone: user.Phone,
		Token: token,
	})
}

//func (*User) GetByToken(c *gin.Context) {
//	user, err := Service.UserService.Get(GetUid(c))
//	if err != nil {
//		log.Error(err)
//		c.JSON(http.StatusInternalServerError, lib.InternalServerError)
//		return
//	}
//
//	c.JSON(http.StatusOK, struct {
//		Id          int        `json:"id"`
//		Phone       string     `json:"phone"`
//		Balance     int        `json:"balance"`
//		LastLoginAt *time.Time `json:"lastLoginAt"`
//	}{
//		Id:          user.Id,
//		Phone:       user.Phone,
//		Balance:     user.Balance,
//		LastLoginAt: user.LastLoginAt,
//	})
//}

//func (*User) GetUserGuidanceCount(c *gin.Context) {
//	stage := c.Query("stage")
//	if len(stage) == 0 {
//		badRequest(c)
//		return
//	}
//	count, err := Service.GetGuidanceCount(GetUid(c), stage)
//	if err != nil {
//		log.Error(err)
//		internalServerError(c)
//		return
//	}
//	c.JSON(http.StatusOK, struct {
//		Count int `json:"count"`
//	}{Count: count})
//}

// NOTE: 容易被滥用，应和前端对好 stage
//func (*User) UpdateUserGuidanceCount(c *gin.Context) {
//	input := struct {
//		Stage string `json:"stage" binding:"required"`
//		Count int    `json:"count" binding:"required"`
//	}{}
//	if err := c.BindJSON(&input); err != nil {
//		internalServerError(c)
//		return
//	}
//	err := Service.UpdateGuidanceCount(GetUid(c), input.Stage, input.Count)
//	if err != nil {
//		log.Error(err)
//		internalServerError(c)
//		return
//	}
//	c.JSON(http.StatusOK, lib.OK)
//}
