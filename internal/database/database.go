package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

func Connect(databaseURL string) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	return &Database{
		Conn: conn,
	}, nil
}

func (db *Database) Close() {
	if db.Conn != nil {
		db.Conn.Close(context.Background())
	}
}
