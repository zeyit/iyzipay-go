package model

type BankTransfer struct {
	SubMerchantKey             string `json:"subMerchantKey,omitempty"`
	Iban                       string `json:"iban,omitempty"`
	ContactName                string `json:"contactName,omitempty"`
	ContactSurname             string `json:"contactSurname,omitempty"`
	LegalCompanyTitle          string `json:"legalCompanyTitle,omitempty"`
	MarketplaceSubMerchantType string `json:"marketplaceSubmerchantType,omitempty"`
}

func NewBankTransfer() *BankTransfer {
	return &BankTransfer{}
}
