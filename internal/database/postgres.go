package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/om13rajpal/dbgpt/config"
)

var Pool *pgxpool.Pool

func ConnectPostgres() {
	Pool, _ = pgxpool.New(context.Background(), config.POSTGRES_URI)

	fmt.Println("Connected to postgres database")
}
