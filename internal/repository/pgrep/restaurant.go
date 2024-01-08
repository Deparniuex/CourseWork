package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error {
	query := fmt.Sprintf(`INSERT INTO %s (
								name,
								city,
								street
	)
	VALUES ($1, $2, $3)
	RETURNING id
	`, restaurantsTable)
	err := p.Pool.QueryRow(ctx, query, restaurant.Name, restaurant.City, restaurant.Street).Scan(&restaurant.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteRestaurant(ctx context.Context, restaurantID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, restaurantsTable)

	tag, err := p.Pool.Exec(ctx, query, restaurantID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateRestaurant(ctx context.Context, restaurant *entity.Restaurant) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = COALESCE($1, name),
			city = COALESCE($2, city),
			street = COALESCE($3, street),
		WHERE 
			id = $4
	`, restaurantsTable)

	tag, err := p.Pool.Exec(ctx, query,
		restaurant.Name,
		restaurant.City,
		restaurant.Street,
		restaurant.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("restaurant doesn't exist")
	}

	return nil
}
