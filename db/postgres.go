package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect(dsn string) {
	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ failed to connect to Postgres: %v", err)
	}

	if err := Pool.Ping(context.Background()); err != nil {
		log.Fatalf("❌ cannot ping Postgres: %v", err)
	}

	log.Println("✅ connected to Postgres")
}
