package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error {
	query := fmt.Sprintf(`INSERT INTO %s (
								order
	)
	VALUES ($1)
	RETURNING id
	`, orderHistoryTable)
	err := p.Pool.QueryRow(ctx, query, orderHistory.Order).Scan(&orderHistory.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteOrderHistory(ctx context.Context, orderHistoryID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, orderHistoryTable)

	tag, err := p.Pool.Exec(ctx, query, orderHistoryID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateOrderHistory(ctx context.Context, orderHistory *entity.OrderHistory) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			order = COALESCE($1, order),
		WHERE 
			id = $2
	`, orderHistoryTable)

	tag, err := p.Pool.Exec(ctx, query,
		orderHistory.Order,
		orderHistory.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("order history doesn't exist")
	}

	return nil
}
