package pgrep

import (
	"ripProject/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	restaurantsTable  = "restaurants"
	citiesTable       = "cities"
	dishesTable       = "dishes"
	menuTable         = "menu"
	orderHistoryTable = "order_history"
	ordersTable       = "orders"
	promosTable       = "promos"
	streetsTable      = "streets"
	tokensTable       = "tokens"
	typeOfDishesTable = "type_of_dish"
	typeOfMenuTable   = "type_of_menu"
	usersTable        = "users"
)

type Postgres struct {
	Pool   *pgxpool.Pool
	Config *config.Config
}

func New(pool *pgxpool.Pool, config *config.Config) *Postgres {
	return &Postgres{
		Pool:   pool,
		Config: config,
	}
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
