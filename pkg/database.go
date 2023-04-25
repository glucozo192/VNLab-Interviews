package pkg

import (
	"context"
)

type Table interface {
	TableName() string
}

type CommandTag interface {
	RowsAffected() int64
}
type Rows interface {
	Scan(dest ...interface{}) (err error)
	Next() bool
	Close()
}

type Row interface {
	Scan(dest ...interface{}) (err error)
}

type QueryExecer interface {
	Exec(sql string, arguments ...interface{}) (CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
}
