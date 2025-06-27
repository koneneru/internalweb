package data

import (
	"authService/internal/validator"
	"context"
	"database/sql"
	"fmt"
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
	Hash      []byte
}

func (u *User) IsAnonimous() bool {
	return u == AnonimousUser
}

func ValidateUsername(v *validator.Validator, un string) {
	v.Check(un != "", "login", "must be provided")
	v.Check(len(un) <= 19, "username", "must not be greater than 19 bytes")
	v.Check(validator.Matches(un, validator.UsernameRX), "username", "must be a valid username")
}

func ValidatePassword(v *validator.Validator, pwd string) {
	v.Check(pwd != "", "password", "must be provided")
	v.Check(len(pwd) >= 8, "password", "must be a least 8 bytes long")
	//v.Check(validator.Matches(pwd, validator.PasswordRX), "password", "must be a valid password")
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) GetByUsername(username, password string) (*User, error) {
	query := `
		SELECT name
		FROM sysusers
		WHERE PWDCOMPARE(?, password) = 1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var str sql.NullString
	err := m.DB.QueryRowContext(ctx, query, username).Scan(&str)
	if err != nil {
		return nil, err
	}
	isMatches := str.String == username
	fmt.Println(isMatches)
	user.Name = str.String
	// NOT IMPLEMENTED YET
	return &user, nil
}
