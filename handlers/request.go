package handlers

type ProviderRtrRequest struct {
	Npi          string `json:"npi"`
	Tin          string `json:"tin"`
	PpgId        string `json:"ppgId"`
	AmisysNumber string `json:"amisysNumber"`
}
