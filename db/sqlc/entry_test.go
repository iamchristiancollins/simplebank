package db

import (
	"context"
	"testing"

	"github.com/iamchristiancollins/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams {
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
		
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}


func TestCreateEntry (t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry (t *testing.T) {
	// TODO
}

func TestListEntries (t *testing.T) {
	// TODO
}