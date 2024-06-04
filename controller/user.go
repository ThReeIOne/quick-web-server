package controller

import "net/http"

type User struct{}

func (*User) Login(c *gin.Context) {
	input := struct {
		Phone    string `json:"phone" binding:"required"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&input); err != nil {
		badRequest(c)
		return
	}

	var user *model.User
	var uc *model.UserCompany
	var err error

	if input.Password != "" {
		// 验证密码
		if lib.IsPhone(input.Phone) {
			user, err = Service.GetByPhone(input.Phone)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusInternalServerError, lib.InternalServerError)
				return
			}
		} else {
			user, err = Service.GetByUsername(input.Phone)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusInternalServerError, lib.InternalServerError)
				return
			}
		}
		if user == nil || user.Password != lib.Password(input.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
			return
		}

	} else if input.Code != "" {
		// 验证短信验证码
		if valid, err := provider.Sms.ValidateSmsCode(input.Phone, input.Code); err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, lib.InternalServerError)
			return
		} else if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"message": "验证码不正确，请重新输入"})
			return
		}

		user, err = Service.GetByPhone(input.Phone)
		if err != nil {
			log.Error(err)
			internalServerError(c)
			return
		}
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户不存在"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请输入密码或验证码"})
		return
	}

	uc, err = Service.GetLastLoginCompany(user.Id)
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}
	if user == nil || uc == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "用户名或密码错误"})
		return
	}

	company, err := Service.Company.Get(uc.CompanyId)
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}

	// 生成 JWT Token
	jwt := lib.NewJWT()
	token, err := jwt.CreateToken(jwt.CreateClaims(lib.BaseClaims{
		Uid:             user.Id,
		CompanyId:       uc.CompanyId,
		DepartmentId:    uc.DepartmentId,
		Phone:           user.Phone,
		CompanyObjectId: company.ObjectId,
	}))
	if err != nil {
		log.Error(err)
		internalServerError(c)
		return
	}

	go Service.UpdateLastLoginAt(user.Id)

	c.JSON(http.StatusOK, LoginInfo{
		Id:              user.Id,
		Phone:           user.Phone,
		Name:            user.Name,
		Avatar:          user.Avatar,
		CompanyId:       uc.CompanyId,
		DepartmentId:    uc.DepartmentId,
		CompanyObjectId: company.ObjectId,
		Token:           token,
	})
}
