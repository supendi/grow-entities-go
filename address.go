package entities

//Country represent country
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Province represent province
type Province struct {
	ID        string `json:"id"`
	CountryID string `json:"countryId"`
	Name      string `json:"name"`
}

//City represent city
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"provinceId"`
	Name       string `json:"name"`
}

//District represent district
type District struct {
	ID     string `json:"id"`
	CityID string `json:"cityId"`
	Name   string `json:"name"`
}

//SubDistrict represent Subdistrict
type SubDistrict struct {
	ID         string `json:"id"`
	DistrictID string `json:"districtId"`
	Name       string `json:"name"`
}

//Village represent village
type Village struct {
	ID            string `json:"id"`
	SubDistrictID string `json:"subDistrictId"`
	Name          string `json:"name"`
}

//Address represent address
type Address struct {
	ID              string  `json:"id"`
	Name            *string `json:"name"`
	SubDistrictID   string  `json:"subDistrictId"`
	SubDistrictName string  `json:"subDistrictName"`
	DistrictID      string  `json:"districtId"`
	DistrictName    string  `json:"districtName"`
	CityID          string  `json:"cityId"`
	CityName        string  `json:"cityName"`
	ProvinceID      string  `json:"provinceId"`
	ProvinceName    string  `json:"provinceName"`
	CountryID       string  `json:"countryId"`
	CountryName     string  `json:"countryName"`
}
