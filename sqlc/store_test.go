package sqlc

import (
	"context"
	"fmt"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bxcodec/faker/v4"
)

func TestTransferTx(t *testing.T) {

	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	// create two accounts
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := TransferTxParams{
		FromAccountID: fromAccount.AccountID,
		ToAccountID:   toAccount.AccountID,
		Amount:        util.RandomInt(10, 100),
	}

	n := 10

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i)

		ctx := context.WithValue(context.Background(), txKey, txName)

		go func() {

			result, err := store.TransferTx(ctx, arg)
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, fromAccount.AccountID)
		require.Equal(t, transfer.ToAccountID, toAccount.AccountID)
		require.Equal(t, transfer.Amount, arg.Amount)

		diff1 := fromAccount.Balance - result.FromAccount.Balance
		diff2 := result.ToAccount.Balance - toAccount.Balance
		fmt.Println(diff1, diff2)
		require.Equal(t, diff1, diff2)

		require.True(t, diff1 >= 0)

	}

}

func createRandomAccount(t *testing.T) BankAccount {
	arg := CreateAccountParams{
		Owner:    faker.Name(),
		Balance:  util.RandomInt(1000, 2000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
