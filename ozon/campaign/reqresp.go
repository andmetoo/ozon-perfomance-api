package campaign

import "time"

type ListParams struct {
	CampaignIDs   []uint64
	AdvObjectType string
	State         string
}

type ListResponseList struct {
	ID                       uint64    `json:"id"`
	PaymentType              string    `json:"paymentType"`
	Title                    string    `json:"title"`
	State                    string    `json:"state"`
	AdvObjectType            string    `json:"advObjectType"`
	FromDate                 string    `json:"fromDate"`
	ToDate                   string    `json:"toDate"`
	DailyBudget              uint64    `json:"dailyBudget"`
	Placement                []string  `json:"placement"`
	Budget                   uint64    `json:"budget"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
	ProductAutopilotStrategy string    `json:"productAutopilotStrategy"`
	ProductCampaignMode      string    `json:"productCampaignMode"`
}

type ListResponse struct {
	List ListResponseList `json:"list"`
}
