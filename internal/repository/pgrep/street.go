package pgrep

import (
	"context"
	"errors"
	"fmt"
	"ripProject/internal/entity"
	"ripProject/internal/repository"
)

func (p *Postgres) CreateStreet(ctx context.Context, street *entity.Street) error {
	query := fmt.Sprintf(`INSERT INTO %s (
								name
	)
	VALUES ($1)
	RETURNING id
	`, streetsTable)
	err := p.Pool.QueryRow(ctx, query, street.Name).Scan(&street.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteStreet(ctx context.Context, streetID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, streetsTable)

	tag, err := p.Pool.Exec(ctx, query, streetID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return repository.ErrRecordNotFound
	}

	return nil
}

func (p *Postgres) UpdateStreet(ctx context.Context, street *entity.Street) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			name = COALESCE($1, name),
		WHERE 
			id = $2
	`, streetsTable)

	tag, err := p.Pool.Exec(ctx, query,
		street.Name,
		street.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("street doesn't exist")
	}

	return nil
}
