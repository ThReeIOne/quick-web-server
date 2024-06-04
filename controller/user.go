package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick_web_golang/log"
	"quick_web_golang/model"
	"time"
)

type User struct{}

func (*User) Login(c *gin.Context) {
	input := struct {
		Phone string `json:"phone" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
		return
	}
	phone := input.Phone
	valid, err := provider.Sms.ValidateSmsCode(phone, input.Code)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
		return
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	user, err := Service.GetByPhone(phone)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
		return
	}

	if user == nil {
		err = Service.User.CreateUser(phone)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
			return
		}
		user, err = Service.GetByPhone(phone)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
			return
		}
	}

	jwt := lib.NewJWT()
	claims := jwt.CreateClaims(lib.BaseClaims{ID: user.Id, Phone: user.Phone})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
		return
	}

	loginTime := time.Now()
	go Service.UpdateLastLoginAt(user.Id, loginTime)
	go Service.Activity.Experience(user.Id, model.TaskType(""), 0)

	c.JSON(http.StatusOK, struct {
		Id          int        `json:"id"`
		Phone       string     `json:"phone"`
		Balance     int        `json:"balance"`
		Token       string     `json:"token"`
		LastLoginAt *time.Time `json:"lastLoginAt"`
	}{
		Id:          user.Id,
		Phone:       user.Phone,
		Balance:     user.Balance,
		Token:       token,
		LastLoginAt: &loginTime,
	})
}

func (*User) GetByToken(c *gin.Context) {
	user, err := Service.User.Get(GetUid(c))
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}

	c.JSON(http.StatusOK, struct {
		Id          int        `json:"id"`
		Phone       string     `json:"phone"`
		Balance     int        `json:"balance"`
		LastLoginAt *time.Time `json:"lastLoginAt"`
	}{
		Id:          user.Id,
		Phone:       user.Phone,
		Balance:     user.Balance,
		LastLoginAt: user.LastLoginAt,
	})
}

func (*User) GetUserGuidanceCount(c *gin.Context) {
	stage := c.Query("stage")
	if len(stage) == 0 {
		badRequest(c)
		return
	}
	count, err := Service.GetGuidanceCount(GetUid(c), stage)
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}
	c.JSON(http.StatusOK, struct {
		Count int `json:"count"`
	}{Count: count})
}

// NOTE: 容易被滥用，应和前端对好 stage
func (*User) UpdateUserGuidanceCount(c *gin.Context) {
	input := struct {
		Stage string `json:"stage" binding:"required"`
		Count int    `json:"count" binding:"required"`
	}{}
	if err := c.BindJSON(&input); err != nil {
		internalServerError(c)
		return
	}
	err := Service.UpdateGuidanceCount(GetUid(c), input.Stage, input.Count)
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}
	c.JSON(http.StatusOK, lib.OK)
}
