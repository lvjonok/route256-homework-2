package db

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type Client struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Client {
	return &Client{pool: pool}
}
