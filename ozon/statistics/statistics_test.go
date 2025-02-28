package statistics_test

import (
	"bytes"
	"context"
	"github.com/andmetoo/ozon-perfomance-api/internal/auth"
	"github.com/andmetoo/ozon-perfomance-api/internal/test"
	"github.com/andmetoo/ozon-perfomance-api/ozon/statistics"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestByCampaign_Success(t *testing.T) {
	c := statistics.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-performance.ozon.ru:443/api/client/statistics", r.URL.String())
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))
					require.Equal(t, `{"campaigns":["string"],"from":"2019-08-24T14:15:22Z","to":"2019-09-24T14:15:22Z","dateFrom":"string","dateTo":"string","groupBy":"NO_GROUP_BY"}`, test.Body(t, r))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  "UUID": "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1",
						  "vendor": true
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-performance.ozon.ru:443/api/client/statistics",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.ByCampaign(context.Background(), statistics.ByCampaignRequest{
		Campaigns: []string{"string"},
		From:      time.Date(2019, 8, 24, 14, 15, 22, 0, time.UTC),
		To:        time.Date(2019, 9, 24, 14, 15, 22, 0, time.UTC),
		DateFrom:  "string",
		DateTo:    "string",
		GroupBy:   "NO_GROUP_BY",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &statistics.ByCampaignResponse{
		UUID:   "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1",
		Vendor: true,
	}, resp)
}

func TestReportStatus_Success(t *testing.T) {
	c := statistics.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-performance.ozon.ru:443/api/client/statistics/0c159c60-ab92-46d9-9a6b-d225dbf5c7b1", r.URL.String())
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
						  "UUID": "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1",
						  "state": "OK",
						  "createdAt": "2020-03-23T19:30:58.264653Z",
						  "updatedAt": "2020-04-23T19:30:58.264653Z",
						  "request": {
							"attributionDays": "",
							"campaignId": 1,
							"campaigns": [
							  12558
							],
							"dateFrom": "2020-03-01",
							"dateTo": "2020-03-20",
							"from": "2020-03-01T00:00:00Z",
							"groupBy": "DATE",
							"objects": [],
							"to": "2020-03-20T00:00:00Z"
						  },
						  "link": "string",
						  "kind": "STATS"
						}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-performance.ozon.ru:443/api/client/statistics",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.ReportStatus(context.Background(), "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1")
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &statistics.ReportStatusResponse{
		UUID:      "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1",
		State:     "OK",
		CreatedAt: time.Date(2020, 3, 23, 19, 30, 58, 264653000, time.UTC),
		UpdatedAt: time.Date(2020, 4, 23, 19, 30, 58, 264653000, time.UTC),
		Request: statistics.ReportStatusResponseRequest{
			AttributionDays: "",
			CampaignID:      1,
			Campaigns:       []uint64{12558},
			DateFrom:        "2020-03-01",
			DateTo:          "2020-03-20",
			From:            time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			GroupBy:         "DATE",
			To:              time.Date(2020, 3, 20, 0, 0, 0, 0, time.UTC),
		},
		Link: "string",
		Kind: "STATS",
	}, resp)
}

func TestGetReport_Success(t *testing.T) {
	c := statistics.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-performance.ozon.ru:443/api/client/statistics/report?UUID=0c159c60-ab92-46d9-9a6b-d225dbf5c7b1", r.URL.String())
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
									"contentType": "string"
								}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-performance.ozon.ru:443/api/client/statistics",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.GetReport(context.Background(), "0c159c60-ab92-46d9-9a6b-d225dbf5c7b1")
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &statistics.GetReportResponse{
		ContentType: "string",
	}, resp)
}
