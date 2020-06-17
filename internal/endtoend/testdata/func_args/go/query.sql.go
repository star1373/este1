// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const makeIntervalDays = `-- name: MakeIntervalDays :one
SELECT make_interval(days => $1::int)
`

func (q *Queries) MakeIntervalDays(ctx context.Context, dollar_1 int32) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, makeIntervalDays, dollar_1)
	var make_interval interface{}
	err := row.Scan(&make_interval)
	return make_interval, err
}

const makeIntervalMonths = `-- name: MakeIntervalMonths :one
SELECT make_interval(months => $1::int)
`

func (q *Queries) MakeIntervalMonths(ctx context.Context, months int32) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, makeIntervalMonths, months)
	var make_interval interface{}
	err := row.Scan(&make_interval)
	return make_interval, err
}

const makeIntervalSecs = `-- name: MakeIntervalSecs :one
SELECT make_interval(secs => $1)
`

func (q *Queries) MakeIntervalSecs(ctx context.Context, secs interface{}) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, makeIntervalSecs, secs)
	var make_interval interface{}
	err := row.Scan(&make_interval)
	return make_interval, err
}

const plus = `-- name: Plus :one
SELECT plus(b => $2, a => $1)
`

type PlusParams struct {
	B int32
	A int32
}

func (q *Queries) Plus(ctx context.Context, arg PlusParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, plus, arg.B, arg.A)
	var plus int32
	err := row.Scan(&plus)
	return plus, err
}

const tableArgs = `-- name: TableArgs :one
SELECT table_args(x => $1)
`

func (q *Queries) TableArgs(ctx context.Context, x int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, tableArgs, x)
	var table_args int32
	err := row.Scan(&table_args)
	return table_args, err
}