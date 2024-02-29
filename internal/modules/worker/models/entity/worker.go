package entity

import "time"

type BankTicket struct {
	TicketNumber  string    `json:"ticketNumber" bson:"ticketNumber"`
	SeatNumber    int       `json:"seatNumber" bson:"seatNumber"`
	IsUsed        bool      `json:"isUsed" bson:"isUsed"`
	UserId        string    `json:"userId" bson:"userId"`
	QueueId       string    `json:"queueId" bson:"queueId"`
	TicketId      string    `json:"ticketId" bson:"ticketId"`
	EventId       string    `json:"eventId" bson:"eventId"`
	CountryCode   string    `json:"countryCode" bson:"countryCode"`
	Price         int       `json:"price" bson:"price"`
	TicketType    string    `json:"ticketType" bson:"ticketType"`
	PaymentStatus string    `json:"paymentStatus" bson:"paymentStatus"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Country struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	City  string `json:"city" bson:"city"`
	Place string `json:"place" bson:"place"`
}

type TicketDetail struct {
	TicketId       string    `json:"ticketId" bson:"ticketId"`
	EventId        string    `json:"eventId" bson:"eventId"`
	TicketType     string    `json:"ticketType" bson:"ticketType"`
	TicketPrice    int       `json:"ticketPrice" bson:"ticketPrice"`
	TotalQuota     int       `json:"totalQuota" bson:"totalQuota"`
	TotalRemaining int       `json:"totalRemaining" bson:"totalRemaining"`
	ContinentName  string    `json:"continentName" bson:"continentName"`
	ContinentCode  string    `json:"continentCode" bson:"continentCode"`
	Country        Country   `json:"country" bson:"country"`
	Tag            string    `json:"tag" bson:"tag"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
}

type VaNumber struct {
	Bank     string `json:"bank" bson:"bank"`
	VaNumber string `json:"vaNumber" bson:"vaNumber"`
}

type Ticket struct {
	TicketNumber string `json:"ticketNumber" bson:"ticketNumber"`
	EventId      string `json:"eventId" bson:"eventId"`
	TicketType   string `json:"ticketType" bson:"ticketType"`
	SeatNumber   int    `json:"seatNumber" bson:"seatNumber"`
	CountryCode  string `json:"countryCode" bson:"countryCode"`
	TicketId     string `json:"ticketId" bson:"ticketId"`
}

type Payment struct {
	TransactionID     string     `json:"transactionId" bson:"transactionId"`
	StatusCode        string     `json:"statusCode" bson:"statusCode"`
	GrossAmount       string     `json:"grossAmount" bson:"grossAmount"`
	PaymentType       string     `json:"paymentType" bson:"paymentType"`
	TransactionStatus string     `json:"transactionStatus" bson:"transactionStatus"`
	FraudStatus       string     `json:"fraudStatus" bson:"fraudStatus"`
	StatusMessage     string     `json:"statusMessage" bson:"statusMessage"`
	MerchantID        string     `json:"merchantId" bson:"merchantId"`
	PermataVaNumber   string     `json:"permataVaNumber" bson:"permataVaNumber"`
	VaNumbers         []VaNumber `json:"vaNumbers" bson:"vaNumbers"`
	PaymentAmounts    []string   `json:"paymentAmounts" bson:"paymentAmounts"`
	TransactionTime   string     `json:"transactionTime" bson:"transactionTime"`
}

type PaymentHistory struct {
	PaymentId      string    `json:"paymentId" bson:"paymentId"`
	UserId         string    `json:"userId" bson:"userId"`
	Ticket         *Ticket   `json:"ticket" bson:"ticket"`
	Payment        *Payment  `json:"payment" bson:"payment"`
	IsValidPayment bool      `json:"isValidPayment" bson:"isValidPayment"`
	ExpiryTime     time.Time `json:"expiryTime" bson:"expiryTime"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
}

type OnlineTicketConfig struct {
	Tag         string        `json:"tag" bson:"tag"`
	TotalQuota  int           `json:"totalQuota" bson:"totalQuota"`
	CountryList []CountryList `json:"countryList" bson:"countryList"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt" bson:"updatedAt"`
	CreatedBy   string        `json:"createdBy" bson:"createdBy"`
	UpdatedBy   string        `json:"updatedBy" bson:"updatedBy"`
}

type CountryList struct {
	CountryNumber int    `json:"countryNumber" bson:"countryNumber"`
	Percentage    int    `json:"percentage" bson:"percentage"`
	CountryCode   string `json:"countryCode" bson:"countryCode"`
}

type AggregateTotalTicket struct {
	Id                   string `json:"_id" bson:"_id"`
	CountryName          string `json:"countryName" bson:"countryName"`
	TotalAvailableTicket int    `json:"totalAvailableTicket" bson:"totalAvailableTicket"`
	TotalTicket          int    `json:"totalTicket" bson:"totalTicket"`
}
