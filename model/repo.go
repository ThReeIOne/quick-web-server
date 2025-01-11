package model

import "github.com/jmoiron/sqlx"

type Repo struct {
	UserRepo *UserRepo
}

var Repos = &Repo{}

func NewRepo(conn *sqlx.DB) *Repo {
	return &Repo{
		UserRepo: NewUserRepo(conn),
	}
}
