package statistics

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/andmetoo/ozon-perfomance-api/internal/request"
	"github.com/pkg/errors"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Statistics {
	return &Statistics{
		h:   h,
		uri: uri,
	}
}

type Statistics struct {
	h   *http.Client
	uri string
}

func (s *Statistics) ByCampaign(ctx context.Context, req ByCampaignRequest) (*ByCampaignResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ByCampaignRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, s.uri, bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ByCampaignRequest.NewRequest")
	}

	return request.Send[ByCampaignResponse](s.h, r, request.ContentTypeApplicationJson)
}

func (s *Statistics) ReportStatus(ctx context.Context, id string) (*ReportStatusResponse, *http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, s.uri, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ReportStatusRequest.NewRequest")
	}

	r.URL = r.URL.JoinPath(id)

	return request.Send[ReportStatusResponse](s.h, r, request.ContentTypeApplicationJson)
}

func (s *Statistics) GetReport(ctx context.Context, id string) (*GetReportResponse, *http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, s.uri+"/report", nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "GetReportRequest.NewRequest")
	}

	q := r.URL.Query()
	q.Add("UUID", id)
	r.URL.RawQuery = q.Encode()

	return request.Send[GetReportResponse](s.h, r, request.ContentTypeApplicationJson)
}
