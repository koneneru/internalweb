package data

import (
	"atsinfoService/internal/validator"
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
	TrunkGroup int
}

type GatewayModel struct {
	DB *sql.DB
}

func ValidateGateway(v *validator.Validator, g Gateway) {
	v.Check(g.Name != "", "name", "must be provided")
	v.Check(len(g.Name) <= 250, "name", "must not be greater than 250 bytes long")

	v.Check(g.Number != "", "number", "must be provided")
	v.Check(len(g.Number) <= 12, "number", "must not  be greater than 12 bytes")

	v.Check(g.Digit <= 32_767, "digit", "must not be greater than 32 767")
	v.Check(g.Trunk <= 32_767, "trunk", "must not be greater than 32 767")
	v.Check(g.TrunkGroup <= 32_767, "trunkgroup", "must not be greater than 32 767")
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
			TrunkGroup: int(input.groupTrunk.Int16),
		}

		gateways = append(gateways, &gateway)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gateways, err
}

func (m GatewayModel) Insert(g *Gateway) error {
	query := `ISERT INTO ats_Trunks (TRUNK, GROUPTRUNK, DIGIT, NAME, NUMBER)
		VALUES (?,?,?,?,?);
		SELECT @@IDENTITY AS NewID`
	args := []any{g.Trunk, g.TrunkGroup, g.Digit, g.Name, g.Number}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return nil
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&g.Id)
}
