package dto

type CountryQuota struct {
	CountryCode   string `json:"countryCode"`
	CountryNumber int    `json:"countryNumber"`
	TotalQuota    int    `json:"totalQuota"`
}
