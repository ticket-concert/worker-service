package response

import (
	"worker-service/internal/pkg/constants"
)

type Province struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProvinceResp struct {
	CollectionData []Province
	MetaData       constants.MetaData
}

type City struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

type CityResp struct {
	CollectionData []City
	MetaData       constants.MetaData
}

type District struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CityId       string `json:"city_id"`
	CityName     string `json:"city_name"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

type DistrictResp struct {
	CollectionData []District
	MetaData       constants.MetaData
}

type SubDistrict struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	DistrictId   string `json:"district_id"`
	DistrictName string `json:"district_name"`
	CityId       string `json:"city_id"`
	CityName     string `json:"city_name"`
	ProvinceId   string `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

type SubDistrictResp struct {
	CollectionData []SubDistrict
	MetaData       constants.MetaData
}
