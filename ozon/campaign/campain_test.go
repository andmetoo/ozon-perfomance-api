package campaign_test

import (
	"bytes"
	"context"
	"github.com/andmetoo/ozon-perfomance-api/internal/auth"
	"github.com/andmetoo/ozon-perfomance-api/internal/test"
	"github.com/andmetoo/ozon-perfomance-api/ozon/campaign"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestList_Success(t *testing.T) {
	c := campaign.New(
		test.NewTestClient(
			auth.NewRoundTripper(
				test.RoundTripFunc(func(r *http.Request) *http.Response {
					require.Equal(t, "https://api-performance.ozon.ru:443/api/client/campaign?advObjectType=SKU&campaignIds=52374&state=CAMPAIGN_STATE_UNKNOWN", r.URL.String())
					require.Equal(t, test.ApiKey, r.Header.Get(auth.APIKeyHeader))
					require.Equal(t, test.ClientID, r.Header.Get(auth.ClientIDHeader))

					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(bytes.NewBufferString(`{
							  "list": {
								"id": 48949,
								"paymentType": "CPM",
								"title": "Баннерная кампания",
								"state": "CAMPAIGN_STATE_RUNNING",
								"advObjectType": "BANNER",
								"fromDate": "2019-10-07",
								"toDate": "2021-10-07",
								"dailyBudget": 504000000,
								"placement": [
								  "PLACEMENT_PDP"
								],
								"budget": 50000000,
								"createdAt": "2019-10-07T06:28:44.055042Z",
								"updatedAt": "2020-10-01T06:28:44.055042Z",
								"productAutopilotStrategy": "NO_AUTO_STRATEGY",
								"productCampaignMode": "PRODUCT_CAMPAIGN_MODE_AUTO"
							  }
							}`)),
					}
				}),
				test.ClientID,
				test.ApiKey,
			),
		),
		"https://api-performance.ozon.ru:443/api/client/campaign",
	)
	require.NotNil(t, c)

	resp, httpResp, err := c.List(context.Background(), campaign.ListParams{
		CampaignIDs:   []uint64{52374},
		AdvObjectType: "SKU",
		State:         "CAMPAIGN_STATE_UNKNOWN",
	})
	require.Nil(t, err)
	require.NotNil(t, httpResp)
	require.Equal(t, httpResp.StatusCode, http.StatusOK)
	require.EqualValues(t, &campaign.ListResponse{
		List: campaign.ListResponseList{
			ID:                       48949,
			PaymentType:              "CPM",
			Title:                    "Баннерная кампания",
			State:                    "CAMPAIGN_STATE_RUNNING",
			AdvObjectType:            "BANNER",
			FromDate:                 "2019-10-07",
			ToDate:                   "2021-10-07",
			DailyBudget:              504000000,
			Placement:                []string{"PLACEMENT_PDP"},
			Budget:                   50000000,
			CreatedAt:                time.Date(2019, 10, 7, 6, 28, 44, 55042000, time.UTC),
			UpdatedAt:                time.Date(2020, 10, 1, 6, 28, 44, 55042000, time.UTC),
			ProductAutopilotStrategy: "NO_AUTO_STRATEGY",
			ProductCampaignMode:      "PRODUCT_CAMPAIGN_MODE_AUTO",
		},
	}, resp)
}
