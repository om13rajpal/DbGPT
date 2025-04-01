package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/om13rajpal/dbgpt/config"
)

var Pool *pgxpool.Pool

func ConnectPostgres() {
	Pool, _ = pgxpool.New(context.Background(), config.POSTGRES_URI)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := Pool.Ping(ctx)

	if err != nil {
		fmt.Println("Unable to connect to postgres database:", err)
		return
	}

	fmt.Println("Connected to postgres database")
}
