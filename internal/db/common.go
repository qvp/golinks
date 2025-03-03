package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

var Func *Repository

func init() {
	dbDsn := "postgres://golinks:golinks@localhost:5432/golinks"
	pool, err := pgxpool.New(context.Background(), dbDsn) // fixme os.Getenv("DATABASE_URL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err) // todo log
		os.Exit(1)
	}
	Func = &Repository{pool: pool}
}

func (rep *Repository) ClosePool() {
	rep.pool.Close()
}
