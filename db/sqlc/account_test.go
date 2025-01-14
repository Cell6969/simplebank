package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Cell6969/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	ctx := context.Background()
	args := CreateAccountParams{
		Name:     utils.RandomName(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Name, account.Name)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	// delete account first
	ctx := context.Background()
	_ = testQueries.DeleteAllAccount(ctx)

	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	// delete account first
	_ = testQueries.DeleteAllAccount(ctx)

	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(ctx, account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	ctx := context.Background()
	// delete account first
	_ = testQueries.DeleteAllAccount(ctx)

	account1 := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	ctx := context.Background()
	// delete account first
	_ = testQueries.DeleteAllAccount(ctx)

	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(ctx, account1.ID)
	require.NoError(t, err)

	// test to get account again
	account2, err := testQueries.GetAccount(ctx, account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	ctx := context.Background()
	// delete account first
	_ = testQueries.DeleteAllAccount(ctx)

	for i := 0; i <= 10; i++ {
		createRandomAccount(t)
	}

	args := GetListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.GetListAccount(ctx, args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
