package statistics

import "time"

type ByCampaignRequest struct {
	Campaigns []string  `json:"campaigns"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	DateFrom  string    `json:"dateFrom"`
	DateTo    string    `json:"dateTo"`
	GroupBy   string    `json:"groupBy"`
}

type ByCampaignResponse struct {
	UUID   string `json:"UUID"`
	Vendor bool   `json:"vendor"`
}

type ReportStatusResponseRequest struct {
	AttributionDays string    `json:"attributionDays"`
	CampaignID      uint64    `json:"campaignId"`
	Campaigns       []uint64  `json:"campaigns"`
	DateFrom        string    `json:"dateFrom"`
	DateTo          string    `json:"dateTo"`
	From            time.Time `json:"from"`
	GroupBy         string    `json:"groupBy"`
	To              time.Time `json:"to"`
}

type ReportStatusResponse struct {
	UUID      string                      `json:"UUID"`
	State     string                      `json:"state"`
	CreatedAt time.Time                   `json:"createdAt"`
	UpdatedAt time.Time                   `json:"updatedAt"`
	Request   ReportStatusResponseRequest `json:"request"`
	Link      string                      `json:"link"`
	Kind      string                      `json:"kind"`
}

type GetReportResponse struct {
	ContentType string `json:"contentType"`
	Link        string `json:"contentDisposition"`
}
