package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    faker.Name(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

}

func TestDeleteAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	Account, err := testQueries.DeleteAccount(context.Background(), account.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, Account)

}

func TestGetAccountByOwnerName(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	Account, err := testQueries.GetAccountByOwnerName(context.Background(), arg.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, Account)

}

func TestGetAccountByAccountId(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	Account, err := testQueries.GetAccountByAccountId(context.Background(), account.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, Account)

}

func TestListAccounts(t *testing.T) {

	Account, err := testQueries.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, Account)

}
