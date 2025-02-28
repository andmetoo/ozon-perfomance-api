package core

import (
	"errors"
	"github.com/andmetoo/ozon-perfomance-api/ozon/campaign"
	"github.com/andmetoo/ozon-perfomance-api/ozon/statistics"
	"net/http"
	"strings"
	"time"

	"github.com/andmetoo/ozon-perfomance-api/internal/auth"
)

const (
	defaultURI = "https://api-performance.ozon.ru:443/api/client"
)

type ClientOpts struct {
	h *http.Client

	uri string

	clientID string
	apiKey   string

	// Request timeout
	timeout time.Duration
}

type Opts func(c *ClientOpts)

func WithURI(uri string) Opts {
	return func(c *ClientOpts) {
		c.uri = uri
	}
}

func WithClientID(clientID string) Opts {
	return func(c *ClientOpts) {
		c.clientID = clientID
	}
}

func WithApiKey(apiKey string) Opts {
	return func(c *ClientOpts) {
		c.apiKey = apiKey
	}
}

func WithClient(h *http.Client) Opts {
	return func(c *ClientOpts) {
		c.h = h
	}
}

func WithTimeout(dur time.Duration) Opts {
	return func(c *ClientOpts) {
		c.timeout = dur
	}
}

var (
	ErrClientIDRequired = errors.New("clientID is required")
	ErrAPIKeyRequired   = errors.New("apiKey is required")
)

type Client struct {
	campaign   *campaign.Campaign
	statistics *statistics.Statistics
}

func NewClient(opts ...Opts) (*Client, error) {
	c := new(ClientOpts)

	for _, opt := range opts {
		opt(c)
	}

	if c.h == nil {
		c.h = http.DefaultClient
	}

	if c.h.Transport == nil {
		c.h.Transport = http.DefaultTransport
	}

	if strings.EqualFold(c.uri, "") {
		c.uri = defaultURI
	}

	if strings.EqualFold(c.clientID, "") {
		return nil, ErrClientIDRequired
	}

	if strings.EqualFold(c.apiKey, "") {
		return nil, ErrAPIKeyRequired
	}

	if c.h.Timeout == 0 && c.timeout > 0 {
		c.h.Timeout = c.timeout
	}

	c.h.Transport = auth.NewRoundTripper(
		c.h.Transport,
		c.clientID,
		c.apiKey,
	)

	return &Client{
		campaign:   campaign.New(c.h, c.uri),
		statistics: statistics.New(c.h, c.uri),
	}, nil
}

func (c Client) Campaign() *campaign.Campaign {
	return c.campaign
}

func (c Client) Statistics() *statistics.Statistics {
	return c.statistics
}
