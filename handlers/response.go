package handlers

type ProviderRtrResponse struct {
	Metadata Metadata      `json:"metadata"`
	Result   []ApiResponse `json:"result"`
}

type Metadata struct {
	Page       string `json:"page"`
	PageCount  string `json:"pageCount"`
	TotalCount string `json:"totalCount"`
}

type ApiResponse struct {
	PracticeEmail        string                 `json:"practiceEmail"`
	Prescriber           string                 `json:"prescriber"`
	ProvId               string                 `json:"provId"`
	ProvName             string                 `json:"provName"`
	Taxonomies           string                 `json:"taxonomies"`
	Tin                  string                 `json:"tin"`
	IsTeleHealth         bool                   `json:"isTeleHealth"`
	Locations            Location               `json:"location"`
	Speciality           []Speciality           `json:"speciality"`
	Address              []Address              `json:"address"`
	AmisysGroupRecord    []AmisysGroupRecord    `json:"amisysGroupRecord"`
	AmisysProviderRecord []AmisysProviderRecord `json:"amisysProviderRecord"`
	Npi                  []Npi                  `json:"Npi"`
}

type Location struct {
	Phones Phone `json:"Phones"`
}

type Phone struct {
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
}

type PhoneNumber struct {
	AreaCode    string `json:"areaCode"`
	CountryCode string `json:"countryCode"`
	Exchange    string `json:"exchange"`
	Extension   string `json:"extension"`
	Number      string `json:"number"`
	PhoneId     string `json:"phoneId"`
	Type        string `json:"type"`
}

type Speciality struct {
	AtypicalIndicator string `json:"atypicalIndicator"`
	Description       string `json:"description"`
	SpecialityId      string `json:"specialityId"`
}

type Address struct {
	AddressId    string `json:"addressId"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	CityId       string `json:"cityId"`
	Country      string `json:"country"`
	County       string `json:"county"`
	CountyId     string `json:"countyId"`
	EndDate      string `json:"endDate"`
	GeoCoords    string `json:"geoCoords"`
	StartDate    string `json:"startDate"`
	State        string `json:"state"`
	StateCode    string `json:"stateCode"`
	Type         string `json:"type"`
	ZipCode      string `json:"zipCode"`
	ZipPlus4     string `json:"zipPlus4"`
}

type AmisysGroupRecord struct {
	GroupNumber string `json:"groupNumber"`
	HealthPlan  string `json:"healthPlan"`
}

type AmisysProviderRecord struct {
	ProviderNumber string `json:"providerNumber"`
	HealthPlan     string `json:"healthPlan"`
}

type Npi struct {
	EffectiveDate string `json:"effectiveDate"`
	Number        string `json:"number"`
	TermDate      string `json:"termDate"`
	TermReason    string `json:"termReason"`
	Type          string `json:"type"`
}
