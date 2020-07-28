package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type UpdateSubMerchantRequest struct {
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
	SubMerchantKey                string `json:"subMerchantKey,omitempty"`
	IdentityNumber                string `json:"identityNumber,omitempty"`
	TaxNumber                     string `json:"taxNumber,omitempty"`
	Currency                      string `json:"currency,omitempty"`
	SettlementDescriptionTemplate string `json:"settlementDescriptionTemplate,omitempty"`
	SwiftCode                     string `json:"swiftCode,omitempty"`
}

func NewUpdateSubMerchantRequest() *UpdateSubMerchantRequest {
	return &UpdateSubMerchantRequest{BaseRequest: NewBaseRequest()}
}

func (u *UpdateSubMerchantRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(u.BaseRequest.ToPKIRequestString()).
		Append("name", u.Name).
		Append("email", u.Email).
		Append("gsmNumber", u.GsmNumber).
		Append("address", u.Address).
		Append("iban", u.Iban).
		Append("taxOffice", u.TaxOffice).
		Append("contactName", u.ContactName).
		Append("contactSurname", u.ContactSurname).
		Append("legalCompanyTitle", u.LegalCompanyTitle).
		Append("swiftCode", u.SwiftCode).
		Append("currency", u.Currency).
		Append("settlementDescriptionTemplate", u.SettlementDescriptionTemplate).
		Append("subMerchantKey", u.SubMerchantKey).
		Append("identityNumber", u.IdentityNumber).
		Append("taxNumber", u.TaxNumber).
		GetRequestString()
}
