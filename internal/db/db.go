package db

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type client struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *client {
	return &client{pool: pool}
}
