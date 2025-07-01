package data

import (
	"context"
	"database/sql"
	"time"
)

type Gateway struct {
	Id         int64
	Name       string
	Number     string
	Digit      int
	Trunk      int
	GroupTrunk int
}

type GatewayModel struct {
	DB *sql.DB
}

func (m GatewayModel) GetAll() ([]*Gateway, error) {
	query := `SET DATEFORMAT ymd;
		SELECT ID, TRUNK, GROUPTRNk, DIGIT, NAME, NUMBER
		FROM ats_Trunks
		ORDER BY ID`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	gateways := []*Gateway{}

	for rows.Next() {
		var input struct {
			id         int64
			name       sql.NullString
			number     sql.NullString
			digit      sql.NullInt16
			trunk      sql.NullInt16
			groupTrunk sql.NullInt16
		}

		err := rows.Scan(
			&input.id,
			&input.trunk,
			&input.groupTrunk,
			&input.digit,
			&input.name,
			&input.number,
		)
		if err != nil {
			return nil, err
		}

		input.name.String, _ = Win1251ToUTF8(input.name.String)
		input.number.String, _ = Win1251ToUTF8(input.number.String)

		var gateway = Gateway{
			Id:         input.id,
			Name:       input.name.String,
			Number:     input.number.String,
			Digit:      int(input.digit.Int16),
			Trunk:      int(input.trunk.Int16),
			GroupTrunk: int(input.groupTrunk.Int16),
		}

		gateways = append(gateways, &gateway)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gateways, err
}
