package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Settings struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (s *Settings) parseToDSN() string {
	q := url.Values{}
	q.Add("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(s.user, s.password),
		Host:     fmt.Sprintf("%s:%s", s.host, s.port),
		Path:     s.dbname,
		RawQuery: q.Encode(),
	}
	return u.String()
}

func ConnectDB(options ...Option) (*pgxpool.Pool, error) {
	s := new(Settings)
	for _, option := range options {
		option(s)
	}
	config, err := pgxpool.ParseConfig(s.parseToDSN())
	if err != nil {
		return nil, err
	}
	conn, err := pgxpool.NewWithConfig(context.Background(), config)

	return conn, nil
}
