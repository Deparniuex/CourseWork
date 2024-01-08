package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateOrder(ctx context.Context, order *entity.Order) error {
	query := fmt.Sprintf(`INSERT INTO %s (
								restaurant
								dish
	)
	VALUES ($1, $2)
	RETURNING id
	`, ordersTable)
	err := p.Pool.QueryRow(ctx, query, order.Restaurant, order.Dish).Scan(&order.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteOrder(ctx context.Context, orderID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, ordersTable)

	tag, err := p.Pool.Exec(ctx, query, orderID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateOrder(ctx context.Context, order *entity.Order) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			restaurant = COALESCE($1, restaurant),
			dish = COALESCE($2, dish),
		WHERE 
			id = $3
	`, ordersTable)

	tag, err := p.Pool.Exec(ctx, query,
		order.Restaurant,
		order.Dish,
		order.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("order doesn't exist")
	}

	return nil
}
