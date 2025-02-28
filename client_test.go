package core_test

import (
	core "github.com/andmetoo/ozon-perfomance-api"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	clientID = "not_empty_client_id"
	apiKey   = "not_empty_api_key"
)

func TestNewClient_Success(t *testing.T) {
	c, err := core.NewClient(
		core.WithClientID(clientID),
		core.WithApiKey(apiKey),
	)
	require.Nil(t, err)
	require.NotNil(t, c)
}

func TestNewClient_NoClientID(t *testing.T) {
	c, err := core.NewClient(
		core.WithApiKey(apiKey),
	)
	require.Nil(t, c)
	require.NotNil(t, err)
	require.ErrorIs(t, err, core.ErrClientIDRequired)
}

func TestNewClient_NoAPIKey(t *testing.T) {
	c, err := core.NewClient(
		core.WithClientID(clientID),
	)
	require.Nil(t, c)
	require.NotNil(t, err)
	require.ErrorIs(t, err, core.ErrAPIKeyRequired)
}
