package auth

import "net/http"

const (
	ClientIDHeader = "Client-Id"
	APIKeyHeader   = "Api-Key"
)

var _ http.RoundTripper = (*RoundTripper)(nil)

func NewRoundTripper(r http.RoundTripper, clientID, apiKey string) *RoundTripper {
	return &RoundTripper{r: r, clientID: clientID, apiKey: apiKey}
}

type RoundTripper struct {
	r        http.RoundTripper
	clientID string
	apiKey   string
}

func (a *RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set(ClientIDHeader, a.clientID)
	r.Header.Set(APIKeyHeader, a.apiKey)

	return a.r.RoundTrip(r)
}
