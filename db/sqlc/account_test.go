package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mushaidul/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams {
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.Balance,account.Balance)
	require.Equal(t,arg.Currency,account.Currency)
	require.Equal(t,arg.Owner,account.Owner)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t,account2)
	require.Equal(t,account1.Balance,account2.Balance)
	require.Equal(t,account1.Currency,account2.Currency)
	require.Equal(t,account1.Owner,account2.Owner)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	updateArg := UpdateAccountParams{
		Balance: util.RandomMoney(),
		ID: account1.ID,
	}

	result,err := testQueries.UpdateAccount(context.Background(),updateArg)

	require.NoError(t,err)
	require.NotEmpty(t,result)
	account2,err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Equal(t,account2.Balance,updateArg.Balance)
	require.NotEqual(t,account1.Balance,account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(),account1.ID)

	require.NoError(t,err)

	account2,err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t,err)
	require.EqualError(t,err, sql.ErrNoRows.Error())
	require.Empty(t,account2)
}

func TestListAccounts(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}
	
	accounts,err := testQueries.ListAccounts(context.Background(),arg)

	require.NoError(t,err)
	require.Len(t,accounts,5)

	for _, account := range accounts {
		require.NotEmpty(t,account)
	}
}
