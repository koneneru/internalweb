package data

import (
	"context"
	"database/sql"
	"time"
)

type Entity struct {
	Id      string `json:"Id"`
	Name    string `json:"name"`
	Version string `json:"-"`
}

type EntityModel struct {
	DB *sql.DB
}

func (m EntityModel) GetAll() ([]*Entity, error) {
	query := `SELECT *
		FROM table 
		WHERE id=1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	entities := []*Entity{}

	for rows.Next() {
		var entity Entity

		err := rows.Scan(
			&entity.Id,
			&entity.Name,
			&entity.Version,
		)
		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
