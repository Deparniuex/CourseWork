package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateCity(ctx context.Context, city *entity.City) error {
	query := fmt.Sprintf(`INSERT INTO %s (
								city_name
	)
	VALUES ($1)
	RETURNING id
	`, citiesTable)
	err := p.Pool.QueryRow(ctx, query, city.CityName).Scan(&city.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteCity(ctx context.Context, cityID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, citiesTable)

	tag, err := p.Pool.Exec(ctx, query, cityID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateCity(ctx context.Context, city *entity.City) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = COALESCE($1, name),
		WHERE 
			id = $2
	`, citiesTable)

	tag, err := p.Pool.Exec(ctx, query,
		city.CityName,
		city.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("city doesn't exist")
	}

	return nil
}
