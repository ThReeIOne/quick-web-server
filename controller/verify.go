package controller

import "quick_web_golang/lib"

var (
	CommonVerify = lib.Rules{
		"ImageUrl":     {lib.RegexpMatch("^(http|https)://[a-zA-Z0-9]+(.[a-zA-Z0-9]+)+([a-zA-Z0-9-._?,'+/\\~:#[]@!$&*])*(.png|.jpg|.jpeg)$")},
		"Radio":        {lib.RegexpMatch("^([1-9][0-9]*):([1-9][0-9]*)$")},
		"DepartmentId": {lib.Gt("0")},
		"RoleId":       {lib.Gt("0")},
		"PId":          {lib.Gt("0")},
		"Name":         {lib.NotEmpty()},
		"Phone":        {lib.RegexpMatch("^1[3456789]\\d{9}$")},
	}
	PageVerify = lib.Rules{
		"Size": {lib.Gt("0"), lib.Le("500")},
		"Page": {lib.Ge("0")},
	}
)
