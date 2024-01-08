package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateMenu(ctx context.Context, menu *entity.Menu) error {
	query := fmt.Sprintf(`INSERT INTO %s (
		name
		restaurant  
	)
	VALUES ($1, $2)
	RETURNING id
	`, menuTable)
	err := p.Pool.QueryRow(ctx, query, menu.Name, menu.Restaurant).Scan(&menu.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteMenu(ctx context.Context, menuID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, menuTable)

	tag, err := p.Pool.Exec(ctx, query, menuID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateMenu(ctx context.Context, menu *entity.Menu) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = COALESCE($1, name),
			restaurant = COALESCE($2, restaurant)
		WHERE 
			id = $3
	`, menuTable)

	tag, err := p.Pool.Exec(ctx, query,
		menu.Name,
		menu.Restaurant,
		menu.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("menu doesn't exist")
	}

	return nil
}
