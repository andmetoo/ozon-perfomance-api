package test

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	ApiKey   = "my-api-key"
	ClientID = "my-client-id"
)

var _ http.RoundTripper = (*RoundTripFunc)(nil)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Body(t *testing.T, r *http.Request) string {
	body, err := io.ReadAll(r.Body)
	require.Nil(t, err)

	return string(body)
}

func TimeFromString(t *testing.T, format, datetime string) time.Time {
	dt, err := time.Parse(format, datetime)
	require.Nil(t, err)

	return dt
}
