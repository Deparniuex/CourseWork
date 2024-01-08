package pgrep

import (
	"context"
	"fmt"
	"ripProject/internal/entity"
)

func (p *Postgres) CreateToken(ctx context.Context, token *entity.Token) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			hash,
			user_id
		)
		VALUES ($1, $2)
	`, tokensTable)
	_, err := p.Pool.Exec(ctx, query, token.Hash, token.UserID)
	if err != nil {
		return err
	}

	return nil
}
