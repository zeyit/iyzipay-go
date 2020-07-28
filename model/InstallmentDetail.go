package model

type InstallmentDetail struct {
	BinNumber         string              `json:"binNumber,omitempty"`
	Price             string              `json:"price,omitempty"`
	CardType          string              `json:"cardType,omitempty"`
	CardAssociation   string              `json:"cardAssociation,omitempty"`
	CardFamilyName    string              `json:"cardFamilyName,omitempty"`
	Force3Ds          *int                `json:"force3Ds,omitempty"`
	BankCode          *int64              `json:"bankCode,omitempty"`
	BankName          string              `json:"bankName,omitempty"`
	ForceCvc          *int                `json:"forceCvc,omitempty"`
	InstallmentPrices []*InstallmentPrice `json:"installmentPrices,omitempty"`
}

func NewInstallmentDetail() *InstallmentDetail {
	return &InstallmentDetail{}
}
