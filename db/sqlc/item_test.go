package db

import (
	"context"
	"testing"

	"github.com/jasonsemken/todo/util"
	"github.com/stretchr/testify/require"
)

func createRandomItem(t *testing.T) Item {
	arg := util.RandomString(20)

	item, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)

	require.NotZero(t, item.ID)
	require.NotZero(t, item.Item)

	return item
}

func TestCreateItem(t *testing.T) {
	createRandomItem(t)
}

func TestGetItem(t *testing.T) {

	item1 := createRandomItem(t)
	item2, err := testQueries.GetItem(context.Background(), item1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Item, item2.Item)

}

func TestGetItemForUpdate(t *testing.T) {

	item1 := createRandomItem(t)
	item2, err := testQueries.GetItemForUpdate(context.Background(), item1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, item1.ID, item2.ID)
	require.Equal(t, item1.Item, item2.Item)

}

func TestListItem(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomItem(t)
	}

	arg := ListItemParams{
		Limit:  5,
		Offset: 5,
	}

	items, err := testQueries.ListItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, items, 5)

	for _, item := range items {
		require.NotEmpty(t, item)
	}
}

func TestUpdateItem(t *testing.T) {
	item1 := createRandomItem(t)

	arg := UpdateItemParams{
		ID:   item1.ID,
		Item: util.RandomString(30),
	}
	item2, err := testQueries.UpdateItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, item2)

	require.Equal(t, arg.ID, item2.ID)
	require.Equal(t, arg.Item, item2.Item)
}
