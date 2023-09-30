package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Batch holds sql to be executed at once
type Batch struct {
	db    *pgxpool.Pool
	batch *pgx.Batch
	ctx   context.Context
	cnt   int
}

// NewBatch returns a new Batch object
func NewBatch(ctx context.Context, db *pgxpool.Pool) *Batch {
	b := &Batch{}
	b.db = db
	b.batch = &pgx.Batch{}
	b.ctx = ctx
	return b
}

// Queue adds an sql statement to the batch
func (b *Batch) Queue(query string, arguments ...any) {
	b.batch.Queue(query, arguments...)
	b.cnt++
}

// Exec executes all of the sql in the batch and returns the first error
func (b *Batch) Exec() (int64, error) {
	response := b.db.SendBatch(b.ctx, b.batch)
	defer response.Close()

	var cnt int64
	for i := 0; i < b.cnt; i++ {
		tag, err := response.Exec()
		if err != nil {
			return 0, err
		}
		cnt += tag.RowsAffected()
	}

	return cnt, nil
}
