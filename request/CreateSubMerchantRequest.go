package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CreateSubMerchantRequest struct {
	*BaseRequest
	Name                          string `json:"name,omitempty"`
	Email                         string `json:"email,omitempty"`
	GsmNumber                     string `json:"gsmNumber,omitempty"`
	Address                       string `json:"address,omitempty"`
	Iban                          string `json:"iban,omitempty"`
	TaxOffice                     string `json:"taxOffice,omitempty"`
	ContactName                   string `json:"contactName,omitempty"`
	ContactSurname                string `json:"contactSurname,omitempty"`
	LegalCompanyTitle             string `json:"legalCompanyTitle,omitempty"`
	SubMerchantExternalId         string `json:"subMerchantExternalId,omitempty"`
	IdentityNumber                string `json:"identityNumber,omitempty"`
	TaxNumber                     string `json:"taxNumber,omitempty"`
	SubMerchantType               string `json:"subMerchantType,omitempty"`
	Currency                      string `json:"currency,omitempty"`
	SettlementDescriptionTemplate string `json:"settlementDescriptionTemplate,omitempty"`
	SwiftCode                     string `json:"swiftCode,omitempty"`
}

func NewCreateSubMerchantRequest() *CreateSubMerchantRequest {
	return &CreateSubMerchantRequest{BaseRequest: NewBaseRequest()}
}

func (c *CreateSubMerchantRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("name", c.Name).
		Append("email", c.Email).
		Append("gsmNumber", c.GsmNumber).
		Append("address", c.Address).
		Append("iban", c.Iban).
		Append("taxOffice", c.TaxOffice).
		Append("contactName", c.ContactName).
		Append("contactSurname", c.ContactSurname).
		Append("legalCompanyTitle", c.LegalCompanyTitle).
		Append("swiftCode", c.SwiftCode).
		Append("currency", c.Currency).
		Append("settlementDescriptionTemplate", c.SettlementDescriptionTemplate).
		Append("subMerchantExternalId", c.SubMerchantExternalId).
		Append("identityNumber", c.IdentityNumber).
		Append("taxNumber", c.TaxNumber).
		Append("subMerchantType", c.SubMerchantType).
		GetRequestString()
}
