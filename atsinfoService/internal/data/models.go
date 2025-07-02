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
	Calls       CallModel
	Gateways    GatewayModel
	Subscribers SubscriberModel
}

func NewModels(dbs, dbd *sql.DB) Models {
	return Models{
		Calls:       CallModel{DB: dbs},
		Gateways:    GatewayModel{DB: dbs},
		Subscribers: SubscriberModel{DB: dbd},
	}
}
