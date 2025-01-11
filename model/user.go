package model

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	Id          int       `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Phone       string    `db:"phone" json:"phone"`
	Avatar      string    `db:"avatar" json:"avatar"`
	Password    string    `db:"password" json:"password"`
	Description string    `db:"description" json:"description"`
	Salt        string    `db:"salt" json:"salt"`
	Lang        string    `db:"lang" json:"lang"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
type UserRepo struct {
	conn *sqlx.DB
}

func NewUserRepo(conn *sqlx.DB) *UserRepo {
	return &UserRepo{
		conn: conn,
	}
}

func (r *UserRepo) GetConn() *sqlx.DB {
	return r.conn
}

func (t *UserRepo) Get(id int) (User, error) {
	var user User
	err := t.conn.Get(&user, "select * from user where id = ?", id)
	return user, err
}

func (t *UserRepo) GetByUsername(username string) (User, error) {
	var user User
	err := t.conn.Get(&user, "select * from user where username = ?", username)
	return user, err
}

func (r *UserRepo) GetByPhone(phone string) (*User, error) {
	var user User
	err := r.conn.Get(&user, "select * from user where phone = ?", phone)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepo) CreateUser(trx *sqlx.Tx, user *User) error {
	q := `
		insert into user (username,phone,avatar,password,description)
		value (?,?,?,?,?)
	`
	if res, err := trx.Exec(
		q,
		user.Username,
		user.Phone,
		user.Avatar,
		user.Password,
		user.Description,
	); err != nil {
		return err
	} else if id, err := res.LastInsertId(); err != nil {
		return err
	} else {
		user.Id = int(id)
	}

	return nil
}

func (r *UserRepo) ExistPhone(companyId int, phone string) (bool, error) {
	var cnt int

	q := `
		select count(1) from user u
		join user_company uc on u.id = uc.user_id
		where uc.company_id=? and u.phone = ? and u.deleted_at is null
	`

	err := r.conn.Get(
		&cnt,
		q,
		companyId,
		phone,
	)
	return cnt > 0, err
}

func (r *UserRepo) UpdateLastLoginAt(id int) error {
	_, err := r.conn.Exec("update user set last_login_at = now() where id = ?", id)
	return err
}
