package request

type CreateTicketRequest struct {
	TicketType  string `json:"ticketType" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
	Counter     int    `json:"counter"`
	Prefix      string `json:"prefix" validate:"required"`
}

type CreateTicketReq struct {
	TicketId string `json:"ticketId"`
	EventId  string `json:"eventId"`
}

type CreateOnlineTicketReq struct {
	Tag         string `json:"tag" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

type UpdateBankTicketRequest struct {
	TicketNumber string `json:"ticketNumber"`
	Price        int    `json:"price"`
}

type UpdateOnlineTicketConfigReq struct {
	Tag           string `json:"tag"`
	CountryNumber int    `json:"countryNumber"`
	CountryCode   string `json:"countryCode"`
}

type CountryList struct {
	CountryNumber int    `json:"countryNumber"`
	Percentage    int    `json:"percentage"`
	CountryCode   string `json:"countryCode"`
}

type UpdateTicketDetailReq struct {
	Tag            string `json:"tag"`
	TicketType     string `json:"ticketType"`
	CountryCode    string `json:"countryCode"`
	TotalQuota     int    `json:"totalQuota"`
	TotalRemaining int    `json:"totalRemaining"`
}

type UpdateTicketDetailByIdReq struct {
	TicketId       string `json:"ticketId"`
	TotalRemaining int    `json:"totalRemaining"`
}

type TicketDetailByTagReq struct {
	Tag         string `json:"tag"`
	TicketType  string `json:"ticketType"`
	CountryCode string `json:"countryCode"`
}
