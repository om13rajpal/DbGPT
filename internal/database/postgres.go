package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/om13rajpal/dbgpt/config"
)

func ConnectPostgres() {
	pool, err := pgxpool.New(context.Background(), config.POSTGRES_URI)

	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	fmt.Println("Connected to postgres database")
	defer pool.Close()
}
