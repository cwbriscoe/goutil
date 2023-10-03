// Copyright 2023 Christopher Briscoe.  All rights reserved.
package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgConnInfo stores postgres database connection info
type PgConnInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// GetPgPool connects to a postgres db and returns a pool of connections
func GetPgPool(info PgConnInfo) (*pgxpool.Pool, error) {
	str := "postgresql://" +
		info.Host + ":" +
		info.Port + "/" +
		info.Name + "?user=" +
		info.User + "&password=" +
		info.Pass
	pool, err := pgxpool.New(context.Background(), str)
	return pool, err
}

// GetPgConn connects to a postgres db and returns a connection
func GetPgConn(info PgConnInfo) (*pgx.Conn, error) {
	str := "postgresql://" +
		info.Host + ":" +
		info.Port + "/" +
		info.Name + "?user=" +
		info.User + "&password=" +
		info.Pass
	conn, err := pgx.Connect(context.Background(), str)
	return conn, err
}
