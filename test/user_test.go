package test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	patch := gomonkey.ApplyFunc(time.Now, func() time.Time { return now })
	defer patch.Reset()

	type params struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	type res struct {
		Id           int    `json:"id"`
		Phone        string `json:"phone"`
		CompanyId    int    `json:"companyId"`
		DepartmentId int    `json:"departmentId"`
		Token        string `json:"token"`
	}
	tests := []struct {
		name   string
		method string
		path   string
		in     params
		want   res
		status int
		sql    []string
	}{
		{
			name:   "正常登陆",
			method: "POST",
			path:   "/user/login",
			in: params{
				Phone: user.Phone,
				Code:  code,
			},
			want: res{
				Id:           1,
				Phone:        user.Phone,
				CompanyId:    userCompany.CompanyId,
				DepartmentId: userCompany.DepartmentId,
				Token:        token,
			},
			status: http.StatusOK,
			sql:    []string{},
		},
		{
			name:   "无效的验证码",
			method: "POST",
			path:   "/user/login",
			in: params{
				Phone: user.Phone,
				Code:  "654321",
			},
			want:   res{},
			status: http.StatusBadRequest,
			sql:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup()
			before(c, tt.method, tt.path, tt.in, tt.sql)
			server.User.Login(c)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data res
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, data, tt.want)
		})
	}
}

func TestListUser(t *testing.T) {
	type params struct {
		DepartmentId int    `json:"departmentId"`
		Name         string `json:"name"`
		Page         int    `json:"page"`
		Size         int    `json:"size"`
	}
	type platform struct {
		Authorized bool   `json:"authorized"`
		OpenId     string `json:"openId"`
		Platform   string `json:"platform"`
	}
	type user struct {
		Id           int         `json:"id"`
		Name         string      `json:"name"`
		Phone        string      `json:"phone"`
		Avatar       string      `json:"avatar"`
		CompanyId    int         `json:"companyId"`
		DepartmentId int         `json:"departmentId"`
		Platform     []*platform `json:"platform"`
	}
	type res struct {
		Page  int     `json:"page"`
		Size  int     `json:"size"`
		Total int     `json:"total"`
		List  []*user `json:"list"`
	}
	var tests = []struct {
		name   string
		method string
		path   string
		in     params
		want   res
		status int
		sql    []string
	}{
		{
			name:   "正常获取",
			method: "GET",
			path:   "/user",
			in: params{
				DepartmentId: 1,
				Name:         "TEST",
				Page:         0,
				Size:         1,
			},
			want: res{
				Page:  0,
				Size:  1,
				Total: 1,
				List: []*user{
					{
						Id:           1,
						Name:         "TEST",
						Phone:        "13688889999",
						Avatar:       "https://www.cemeta.cn/avatar.png",
						CompanyId:    1,
						DepartmentId: 1,
						Platform: []*platform{
							{
								Authorized: true,
								OpenId:     "abc",
								Platform:   "douyin",
							},
						},
					},
				},
			},
			status: http.StatusOK,
			sql:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup()
			before(c, tt.method, tt.path, tt.in, tt.sql)
			server.User.List(c)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data res
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, data, tt.want)
		})
	}
}

func TestCreateUser(t *testing.T) {
	type params struct {
		DepartmentId int    `json:"departmentId"`
		RoleId       int    `json:"roleId"`
		Name         string `json:"name"`
		Phone        string `json:"phone"`
		Avatar       string `json:"avatar"`
		Description  string `json:"description"`
	}
	type res struct {
		Message string `json:"message"`
	}
	var tests = []struct {
		name   string
		method string
		path   string
		in     params
		want   res
		status int
		sql    []string
	}{
		{
			name:   "正常创建",
			method: "POST",
			path:   "/user",
			in: params{
				DepartmentId: 1,
				RoleId:       1,
				Name:         "Tests",
				Phone:        "13699998888",
				Avatar:       "https://www.cenbo.cn/xxx.jpg",
				Description:  "test",
			},
			want: res{
				Message: "ok",
			},
			status: http.StatusOK,
			sql:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup()
			before(c, tt.method, tt.path, tt.in, tt.sql)
			server.User.Create(c)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data res
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, data, tt.want)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type params struct {
		Id           int    `json:"id"`
		DepartmentId int    `json:"departmentId" binding:"departmentId"`
		RoleId       int    `json:"roleId" binding:"roleId"`
		Phone        string `json:"phone" binding:"phone"`
		Avatar       string `json:"avatar" binding:"avatar"`
		Description  string `json:"description" binding:"description"`
	}
	type res struct {
		Message string `json:"message"`
	}
	var tests = []struct {
		name   string
		method string
		path   string
		in     params
		want   res
		status int
		sql    []string
	}{
		{
			name:   "正常修改",
			method: "PUT",
			path:   "/user/:id",
			in: params{
				Id:           1,
				DepartmentId: 1,
				RoleId:       1,
				Phone:        "13699998881",
				Avatar:       "https://www.cenbo.cn/xxx.jpg",
				Description:  "testt",
			},
			want: res{
				Message: "ok",
			},
			status: http.StatusOK,
			sql:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup()
			before(c, tt.method, tt.path, tt.in, tt.sql)
			server.User.Update(c)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data res
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, data, tt.want)
		})
	}
}

func TestRemoveUser(t *testing.T) {
	type params struct {
		Id int `json:"id"`
	}
	type res struct {
		Message string `json:"message"`
	}
	var tests = []struct {
		name   string
		method string
		path   string
		in     params
		want   res
		status int
		sql    []string
	}{
		{
			name:   "正常移除",
			method: "PUT",
			path:   "/user/:id",
			in: params{
				Id: 1,
			},
			want: res{
				Message: "ok",
			},
			status: http.StatusOK,
			sql:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := setup()
			before(c, tt.method, tt.path, tt.in, tt.sql)
			server.User.Remove(c)

			assert.Nil(t, c.Err())
			assert.Equal(t, tt.status, w.Code)

			var data res
			assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &data))
			assert.Equal(t, data, tt.want)
		})
	}
}
