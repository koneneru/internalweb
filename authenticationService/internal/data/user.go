package data

import (
	"authService/internal/validator"
	"database/sql"
	"time"
)

var AnonimousUser = &User{}

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type password struct {
	PlainText *string
	hash      []byte
}

func (u *User) IsAnonimous() bool {
	return u == AnonimousUser
}

func ValidateUsername(v *validator.Validator, un string) {
	v.Check(un != "", "login", "must be provided")
	v.Check(validator.Matches(un, validator.UsernameRX), "username", "must be a valid username")
}

func ValidatePassword(v *validator.Validator, pwd string) {
	v.Check(pwd != "", "password", "must be provided")
	v.Check(validator.Matches(pwd, validator.PasswordRX1), "password", "must be a valid password")
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) GetByUsername(username, password string) bool {
	query := ``
	// NOT IMPLEMENTED YET
	return true
}
