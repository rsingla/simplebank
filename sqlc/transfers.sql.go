// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transfers.sql

package sqlc

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO "bank"."transfers" (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING entry_id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (BankTransfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i BankTransfer
	err := row.Scan(
		&i.EntryID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByEntryId = `-- name: GetTransferByEntryId :one
select from_account_id, to_account_id, amount, created_at  from "bank"."transfers" where entry_id = 1
`

type GetTransferByEntryIdRow struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	CreatedAt     sql.NullTime
}

func (q *Queries) GetTransferByEntryId(ctx context.Context) (GetTransferByEntryIdRow, error) {
	row := q.db.QueryRowContext(ctx, getTransferByEntryId)
	var i GetTransferByEntryIdRow
	err := row.Scan(
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
