package sqrl

import (
	"context"
	"database/sql"
)

// ResultBuilder wraps the Insert/Update/Delete builder to return sql.Result methods.
type ResultBuilder interface {
	RunWith(db ExecerContext) ResultBuilder
	Exec() (int64, error)
	ExecContext(ctx context.Context) (int64, error)
}

type resultBuilder struct {
	db       ExecerContext
	sq       Sqlizer
	callback func(sql.Result) (int64, error)
}

func (b *resultBuilder) RunWith(db ExecerContext) ResultBuilder {
	b.db = db
	return b
}

func (b *resultBuilder) Exec() (int64, error) {
	return b.ExecContext(context.Background())
}

func (b *resultBuilder) ExecContext(ctx context.Context) (int64, error) {
	res, err := ExecWithContext(ctx, b.db, b.sq)
	if err != nil {
		return 0, err
	}

	c, err := b.callback(res)
	if err != nil {
		return 0, err
	}

	return c, nil
}
