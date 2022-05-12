package dbconnector

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New(ctx context.Context, url string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, url)
}
