package company

type CompanyDto struct {
	CompanyId            int64  `json:"company_id,omitempty"`
	CompanyName          string `json:"company_name,omitempty"`
	NumberOfTransactions int64  `json:"transactions,omitempty"`
	Location             string `json:"location,omitempty"`
	LastUpdated          int64  `json:"last_updated,omitempty"`
}

type CompanyDtos struct {
	CompanyDto []*CompanyDto
	Total      int64
}
