// Copyright 2023 Christopher Briscoe.  All rights reserved.
package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config stores postgres database connection info
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

// NewPgPool connects to a postgres db and returns a pool of connections
func NewPgPool(cfg *Config) (*pgxpool.Pool, error) {
	str := "postgresql://" +
		cfg.Host + ":" +
		cfg.Port + "/" +
		cfg.Name + "?user=" +
		cfg.User + "&password=" +
		cfg.Pass
	pool, err := pgxpool.New(context.Background(), str)
	return pool, err
}

// NewPgConn connects to a postgres db and returns a connection
func NewPgConn(cfg *Config) (*pgx.Conn, error) {
	str := "postgresql://" +
		cfg.Host + ":" +
		cfg.Port + "/" +
		cfg.Name + "?user=" +
		cfg.User + "&password=" +
		cfg.Pass
	conn, err := pgx.Connect(context.Background(), str)
	return conn, err
}
