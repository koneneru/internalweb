package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Calls CallModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Calls: CallModel{DB: db},
	}
}
