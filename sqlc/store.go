package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions.
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store.
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

var txKey = struct{}{}

// execTx executes a function within a database transaction.
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx error: %v; rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult contains the result of the transfer transaction.
type TransferTxResult struct {
	Transfer    BankTransfer `json:"transfer"`
	FromAccount BankAccount  `json:"from_account"`
	ToAccount   BankAccount  `json:"to_account"`
	FromEntry   BankEntry    `json:"from_entry"`
	ToEntry     BankEntry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to another.
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	transferDetails := CreateTransferParams{
		FromAccountID: arg.FromAccountID,
		ToAccountID:   arg.ToAccountID,
		Amount:        arg.Amount,
	}

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "Create Tansfer")

		result.Transfer, err = q.CreateTransfer(ctx, transferDetails)

		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "Get Account for Update1")
		accountInfo, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)

		fmt.Println(txName, "Update Account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance:   accountInfo.Balance - arg.Amount,
			AccountID: arg.FromAccountID,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "Get Account for Update2")
		accountInfo, err = q.GetAccountForUpdate(ctx, arg.ToAccountID)

		fmt.Println(txName, "Update Account 2")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			Balance:   accountInfo.Balance + arg.Amount,
			AccountID: arg.ToAccountID,
		})

		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
