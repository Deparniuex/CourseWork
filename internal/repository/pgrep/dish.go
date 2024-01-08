package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateDish(ctx context.Context, dish *entity.Dish) error {
	query := fmt.Sprintf(`INSERT INTO %s (
		name,  
		weight,
		price, 
		menu  
	)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`, dishesTable)
	err := p.Pool.QueryRow(ctx, query, dish.Name, dish.Weight, dish.Price, dish.Menu).Scan(&dish.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteDish(ctx context.Context, dishID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, dishesTable)

	tag, err := p.Pool.Exec(ctx, query, dishID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateDish(ctx context.Context, dish *entity.Dish) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = COALESCE($1, name),
			weight = COALESCE($2, weight),
			price = COALESCE($3, price),
			menu = COALESCE($4, menu)
		WHERE 
			id = $5
	`, dishesTable)

	tag, err := p.Pool.Exec(ctx, query,
		dish.Name,
		dish.Weight,
		dish.Price,
		dish.Menu,
		dish.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("dish doesn't exist")
	}

	return nil
}
