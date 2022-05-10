package dbconnector

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO: move to
const url = "postgresql://root:root@localhost:5432/rootdb"

func New(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, url)
}
