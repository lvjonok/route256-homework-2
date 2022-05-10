package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	m "gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func Test_client_CreateProblem(t *testing.T) {
	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		problem m.Problem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client{
				pool: tt.fields.pool,
			}
			if err := c.CreateProblem(tt.args.ctx, tt.args.problem); (err != nil) != tt.wantErr {
				t.Errorf("client.CreateProblem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
