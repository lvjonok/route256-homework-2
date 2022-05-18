package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/config"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/dbconnector"
)

func Prepare(t *testing.T) (*db.Client, context.Context) {
	cfg, err := config.New("../../config.yaml")
	require.NoError(t, err)

	ctx := context.Background()
	adp, err := dbconnector.New(ctx, cfg.Database.Url)
	require.NoError(t, err)

	// truncate tables for preparation
	_, err = adp.Exec(ctx, "TRUNCATE TABLE problems CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE submissions CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE categories CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE images CASCADE;")
	require.NoError(t, err)

	client := db.New(adp)

	return client, ctx
}
