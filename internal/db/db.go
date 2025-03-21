package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nimilgp/URLcommentary/internal/dblayer"
)

type DB struct {
	Conn    *pgx.Conn
	Queries *dblayer.Queries
}

func New(dsn string, ctx context.Context) *DB {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Could Not Establish DB Connection")
	}

	queries := dblayer.New(conn)

	return &DB{conn, queries}
}

func (db *DB) Close(ctx context.Context) error {
	return db.Conn.Close(ctx)
}
