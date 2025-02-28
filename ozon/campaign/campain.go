package campaign

import (
	"context"
	"github.com/andmetoo/ozon-perfomance-api/internal/request"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

func New(
	h *http.Client,
	uri string,
) *Campaign {
	return &Campaign{
		h:   h,
		uri: uri,
	}
}

type Campaign struct {
	h   *http.Client
	uri string
}

func (c *Campaign) List(ctx context.Context, params ListParams) (*ListResponse, *http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.uri, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "CampaignRequest.NewRequest")
	}

	q := r.URL.Query()

	campaignIDsString := make([]string, 0, len(params.CampaignIDs))

	for _, id := range params.CampaignIDs {
		campaignIDsString = append(campaignIDsString, strconv.FormatUint(id, 10))
	}

	q.Add("campaignIds", strings.Join(campaignIDsString, ","))
	q.Add("advObjectType", params.AdvObjectType)
	q.Add("state", params.State)
	r.URL.RawQuery = q.Encode()

	return request.Send[ListResponse](c.h, r, request.ContentTypeApplicationJson)
}
