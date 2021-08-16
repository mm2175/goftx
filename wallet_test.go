package goftx

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/grishinsana/goftx/models"
)

func TestWallet_GetBalances(t *testing.T) {
	_ = godotenv.Load()

	ftx := New(
		WithAuth(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")),
	)
	err := ftx.SetServerTimeDiff()
	require.NoError(t, err)

	balances, err := ftx.GetBalances()
	require.Nil(t, err)
	t.Logf("%+v", balances)
	require.NotNil(t, balances)
}

func TestWallet_Withdraw(t *testing.T) {
	_ = godotenv.Load()

	ftx := New(
		WithAuth(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")),
	)

	payload := models.CreateWithdrawPayload{
		Coin:    os.Getenv("WITHDRAW_COIN"),
		Size:    2,
		Address: os.Getenv("WITHDRAW_ADDR"),
		Tag:     os.Getenv("WITHDRAW_TAG"),
		Method:  os.Getenv("WITHDRAW_NETWORK"),
	}
	res, err := ftx.Withdraw(context.Background(), &payload)

	require.Nil(t, err)
	t.Logf("%+v", res)
	require.NotNil(t, res)
}

func TestWallet_WithdrawHistory(t *testing.T) {
	_ = godotenv.Load()

	ftx := New(
		WithAuth(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")),
	)

	res, err := ftx.WithdrawHistory(context.Background(), 1628611200, 1628697540)
	require.Nil(t, err)
	assert.Equal(t, 4, len(res))
	t.Logf("%+v", res)
}
