package handlers

type ClaimRTRrequest struct {
	ClaimNbr                       string `json:"claimNbr"`
	CheckPayToNbr                  string `json:"checkPayToNbr"`
	SourceSystem                   string `json:"sourceSystem"`
	PlanId                         string `json:"planId"`
	MemberExternalId               string `json:"memberExternalId"`
	CoverageEffectiveDate          string `json:"coverageEffectiveDate"`
	FieldName                      string `json:"fieldName"`
	Fields                         string `json:"fields"`
	Page                           int    `json:"page"`
	PageSize                       int    `json:"pageSize"`
	Sort                           string `json:"sort"`
	Direction                      string `json:"direction"`
	FieldNameFrom                  string `json:"fieldNameFrom"`
	FieldNameTo                    string `json:"fieldNameTo"`
	ProviderCorrespondenceOperator string `json:"providerCorrespondenceOperator"`
}
