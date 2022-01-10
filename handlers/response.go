package handlers

type ClaimRTRresponse struct {
	Links      []Link        `json:"link"`
	Result     []ApiResponse `json:"result"`
	Version    string        `json:"version"`
	Total      int           `json:"total"`
	TotalPages int           `json:"totalPages"`
}

type Link struct {
	Link string `json:"link"`
	Rel  string `json:"rel"`
}

type ApiResponse struct {
	ClaimNbr                              string                  `json:"claimNbr,omitempty"`
	ClaimTypeCd                           string                  `json:"claimTypeCd,omitempty"`
	ClaimTypeCdDesc                       string                  `json:"claimTypeCdDesc,omitempty"`
	EmergencyInd                          string                  `json:"emergencyInd,omitempty"`
	AdmissionDt                           string                  `json:"admissionDt,omitempty"`
	AdmissionSource                       string                  `json:"admissionSource,omitempty"`
	AdmissionTypeCd                       string                  `json:"admissionTypeCd,omitempty"`
	AdmissionTypeDesc                     string                  `json:"admissionTypeDesc,omitempty"`
	BehavioralHealthClaim                 bool                    `json:"behavioralHealthClaim"`
	BillTypeCd                            string                  `json:"billTypeCd,omitempty"`
	BirthWeightCt                         int                     `json:"birthWeightCt,omitempty"`
	NationalDrugCodeBaseMeasureCd         string                  `json:"nationalDrugCodeBaseMeasureCd,omitempty"`
	NationalDrugCodeBaseMeasureDesc       string                  `json:"nationalDrugCodeBaseMeasureDesc,omitempty"`
	NationalDrugCodeNbr                   string                  `json:"nationalDrugCodeNbr,omitempty"`
	NationalDrugCodePrescriptionNbr       string                  `json:"nationalDrugCodePrescriptionNbr,omitempty"`
	NationalDrugCodePrescriptionQualifier string                  `json:"nationalDrugCodePrescriptionQualifier,omitempty"`
	NationalDrugCodePrice                 string                  `json:"nationalDrugCodePrice,omitempty"`
	NationalDrugCodeUnitsCt               string                  `json:"nationalDrugCodeUnitsCt,omitempty"`
	PlaceOfServiceCd                      string                  `json:"placeOfServiceCd,omitempty"`
	PlaceOfServiceDesc                    string                  `json:"placeOfServiceDesc,omitempty"`
	ProcedureDt                           string                  `json:"procedureDt,omitempty"`
	SequenceNbr                           string                  `json:"sequenceNbr,omitempty"`
	SourceProcedureClassCd                string                  `json:"sourceProcedureClassCd,omitempty"`
	SourceProcedureClassDesc              string                  `json:"sourceProcedureClassDesc,omitempty"`
	SourceProcedureQualifierTypeCd        string                  `json:"sourceProcedureQualifierTypeCd,omitempty"`
	SourceProcedureQualifierTypeDesc      string                  `json:"SourceProcedureQualifierTypeCd,omitempty"`
	ServiceLineNbr                        string                  `json:"serviceLineNbr,omitempty"`
	ServiceLineSequenceNbr                string                  `json:"serviceLineSequenceNbr,omitempty"`
	SpecialityProgramInd                  string                  `json:"specilaityProgramInd,omitempty"`
	ThirdPartyStatus                      string                  `json:"thirdPartyStatus,omitempty"`
	DiagnosisRelatedGroupCd               string                  `json:"diagnosisRelatedGroupCd,omitempty"`
	DiagnosisRelatedGroupDesc             string                  `json:"diagnosisRelatedGroupDesc,omitempty"`
	DischargeDispositionCd                string                  `json:"dischargeDispositionCd,omitempty"`
	DischargeDispositionDesc              string                  `json:"dischargeDispositionDesc,omitempty"`
	DischargeDt                           string                  `json:"dischargeDt,omitempty"`
	EpsdtInd                              string                  `json:"epsdtInd,omitempty"`
	ExplanantionCd                        string                  `json:"explanationCd,omitempty"`
	ExplanantionDesc                      string                  `json:"explanationDesc,omitempty"`
	InpatientDaysCt                       int                     `json:"inpatientDaysCt,omitempty"`
	OutcomeStatusCd                       string                  `json:"outcomeStatusCd,omitempty"`
	OutcomeStatusDesc                     string                  `json:"outcomeStatusDesc,omitempty"`
	AdjucatiingMember                     []AdjucatingMember      `json:"adjucatingMember,omitempty"`
	Check                                 []Check                 `json:"check,omitempty"`
	ClaimAffiliation                      []ClaimAffiliation      `json:"claimAffiliation,omitempty"`
	ClaimDiagnosis                        []ClaimDiagnosis        `json:"claimDiagnosis,omitempty"`
	Affiliation                           []Affiliation           `json:"affiliation,omitempty"`
	ClaimServiceDiagnosis                 []ClaimServiceDiagnosis `json:"claimServiceDiagnosis,omitempty"`
	RevenueCode                           RevenueCode             `json:"revenueCode,omitempty"`
	ServiceAffiliation                    []ServiceAffiliation    `json:"serviceAffiliation,omitempty"`
	ServiceHealthProcedure                ServiceHealthProcedure  `json:"serviceHealthProcedure,omitempty"`
	ProcedureCodeModifier                 []ProcedureCodeModifier `json:"procedureCodeModifier,omitempty"`
	Condition                             []Condition             `json:"condition,omitempty"`
	ExternalCauseOfInjury                 []ExternalCauseOfInjury `json:"externalCauseOfInjury,omitempty"`
	HealthProcedure                       []HealthProcedure       `json:"healthProcedure,omitempty"`
	Occurrence                            []Occurrence            `json:"occurrence,omitempty"`
	Provider                              []Provider              `json:"provider,omitempty"`
}

type Provider struct {
	AffiliationNbr    string `json:"affiliationNbr,omitempty"`
	GroupPracticeName string `json:"groupPracticeName,omitempty"`
	ProviderId        string `json:"providerId,omitempty"`
	TypeCd            string `json:"typeCd,omitempty"`
	TypeDesc          string `json:"typeDesc,omitempty"`
}

type Occurrence struct {
	Code        string `json:"code,omitempty"`
	Desc        string `json:"desc,omitempty"`
	SequenceNbr string `json:"sequenceNbr,omitempty"`
	Date        string `json:"date,omitempty"`
}

type HealthProcedure struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ExternalCauseOfInjury struct {
	Code               string `json:"code,omitempty"`
	Desc               string `json:"desc,omitempty"`
	SequenceNbr        string `json:"sequenceNbr,omitempty"`
	PresentOnAdmission string `json:"presentOnAdmission,omitempty"`
}

type Condition struct {
	Code        string `json:"code,omitempty"`
	Desc        string `json:"desc,omitempty"`
	SequenceNbr string `json:"sequenceNbr,omitempty"`
}

type ProcedureCodeModifier struct {
	Code              string `json:"code,omitempty"`
	Desc              string `json:"desc,omitempty"`
	QualifierTypeCd   string `json:"qualifierTypeCd,omitempty"`
	QualifierTypeDesc string `json:"qualifierTypeDesc,omitempty"`
	SequenceNbr       string `json:"sequenceNbr,omitempty"`
}

type ServiceHealthProcedure struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ServiceAffiliation struct {
	ActiveInd                 string `json:"activeInd,omitempty"`
	AddressLine               string `json:"addressLine,omitempty"`
	AffiliationType           string `json:"affiliationType,omitempty"`
	City                      string `json:"city,omitempty"`
	ContractingStatus         string `json:"contractingStatus,omitempty"`
	Country                   string `json:"country,omitempty"`
	IrsNbr                    string `json:"irsNbr,omitempty"`
	MedicaidId                string `json:"medicaidId,omitempty"`
	Npi                       string `json:"npi,omitempty"`
	PostalCd                  string `json:"postalCd,omitempty"`
	ProviderSuppliedFirstName string `json:"providerSuppliedFirstName,omitempty"`
	ProviderSuppliedLastName  string `json:"providerSuppliedLastName,omitempty"`
	State                     string `json:"state,omitempty"`
	Taxonomy                  string `json:"taxonomy,omitempty"`
}

type RevenueCode struct {
	Code string `json:"code,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type AdjucatingMember struct {
	ContractNbr       string `json:"contractNbr,omitempty"`
	ExternalMemberId  string `json:"externalMemberId,omitempty"`
	MemberMedicaidNbr string `json:"memberMedicaidNbr,omitempty"`
	Mpi               string `json:"mpi,omitempty"`
	SourceSystem      string `json:"sourceSystem,omitempty"`
}

type Check struct {
	AddressLine1Text     string `json:"addressLine1Text,omitempty"`
	AddressLine2Text     string `json:"addressLine2Text,omitempty"`
	CityName             string `json:"cityName,omitempty"`
	PaidDt               string `json:"paidDt,omitempty"`
	SequenceNbr          string `json:"sequenceNbr,omitempty"`
	StateCd              string `json:"stateCd,omitempty"`
	StatusCd             string `json:"statusCd,omitempty"`
	StatusDesc           string `json:"statusDesc,omitempty"`
	TaxIdentificationNbr string `json:"taxIdentificationNbr,omitempty"`
	VoidDt               string `json:"voidDt,omitempty"`
	VoidInd              string `json:"voidInd,omitempty"`
	ZipCd                string `json:"zipCd,omitempty"`
}

type ClaimAffiliation struct {
	AddressLine               string `json:"addressLine,omitempty"`
	AffiliationType           string `json:"affiliationType,omitempty"`
	City                      string `json:"city,omitempty"`
	ContractingStatus         string `json:"contractingStatus,omitempty"`
	Country                   string `json:"country,omitempty"`
	IrsNbr                    string `json:"irsNbr,omitempty"`
	MedicaidId                string `json:"medicaidId,omitempty"`
	Name                      string `json:"name,omitempty"`
	Npi                       string `json:"npi,omitempty"`
	PostalCd                  string `json:"postalCd,omitempty"`
	ProviderSuppliedFirstName string `json:"providerSuppliedFirstName,omitempty"`
	ProviderSuppliedLastName  string `json:"providerSuppliedLastName,omitempty"`
	State                     string `json:"state,omitempty"`
	Taxonomy                  string `json:"taxonomy,omitempty"`
}

type ClaimDiagnosis struct {
	Code               string `json:"code,omitempty"`
	Desc               string `json:"desc,omitempty"`
	PresentOnAdmission string `json:"presentOnAdmission,omitempty"`
	SequenceNbr        string `json:"sequenceNbr,omitempty"`
}

type Affiliation struct {
	AffiliationNbr string `json:"affiliationNbr,omitempty"`
	EffectiveDt    string `json:"effectiveDt,omitempty"`
	EndDt          string `json:"endDt,omitempty"`
	IrsNbr         string `json:"irsNbr,omitempty"`
	MedicaidNbr    string `json:"medicaidNbr,omitempty"`
	MedicareNbr    string `json:"medicareNbr,omitempty"`
	Npi            string `json:"npi,omitempty"`
	ProviderNbr    string `json:"providerNbr,omitempty"`
	SpecialityDesc string `json:"specialityDesc,omitempty"`
	Speciality     string `json:"speciality,omitempty"`
	Void           string `json:"void,omitempty"`
}

type ClaimServiceDiagnosis struct {
	Code                 string `json:"code,omitempty"`
	Desc                 string `json:"desc,omitempty"`
	PresentOnAdmission   string `json:"presentOnAdmission,omitempty"`
	SequenceNbr          string `json:"sequenceNbr,omitempty"`
	ClassificationTypeCd string `json:"classificationTypeCd,omitempty"`
}

type AmisysProviderRecord struct {
	ProviderNumber string `json:"providerNumber,omitempty"`
	HealthPlan     string `json:"healthPlan,omitempty"`
}

type Npi struct {
	EffectiveDate string `json:"effectiveDate,omitempty"`
	Number        string `json:"number,omitempty"`
	TermDate      string `json:"termDate,omitempty"`
	TermReason    string `json:"termReason,omitempty"`
	Type          string `json:"type,omitempty"`
}
