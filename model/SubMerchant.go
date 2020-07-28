package model

type SubMerchant struct {
	*IyzipayResource
	Name                          string `json:"name,omitempty"`
	Email                         string `json:"email,omitempty"`
	GsmNumber                     string `json:"gsmNumber,omitempty"`
	Address                       string `json:"address,omitempty"`
	Iban                          string `json:"iban,omitempty"`
	SwiftCode                     string `json:"swiftCode,omitempty"`
	Currency                      string `json:"currency,omitempty"`
	TaxOffice                     string `json:"taxOffice,omitempty"`
	ContactName                   string `json:"contactName,omitempty"`
	ContactSurname                string `json:"contactSurname,omitempty"`
	LegalCompanyTitle             string `json:"legalCompanyTitle,omitempty"`
	SubMerchantExternalId         string `json:"subMerchantExternalId,omitempty"`
	IdentityNumber                string `json:"identityNumber,omitempty"`
	TaxNumber                     string `json:"taxNumber,omitempty"`
	SubMerchantType               string `json:"subMerchantType,omitempty"`
	SubMerchantKey                string `json:"subMerchantKey,omitempty"`
	SettlementDescriptionTemplate string `json:"settlementDescriptionTemplate,omitempty"`
}

func NewSubMerchant() *SubMerchant {
	return &SubMerchant{IyzipayResource: NewIyzipayResource()}
}
