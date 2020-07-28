package model

type ApmResource struct {
	*IyzipayResource
	RedirectUrl                  string         `json:"redirectUrl,omitempty"`
	Price                        string         `json:"price,omitempty"`
	PaidPrice                    string         `json:"paidPrice,omitempty"`
	PaymentId                    string         `json:"paymentId,omitempty"`
	MerchantCommissionRate       string         `json:"merchantCommissionRate,omitempty"`
	MerchantCommissionRateAmount string         `json:"merchantCommissionRateAmount,omitempty"`
	IyziCommissionRateAmount     string         `json:"iyziCommissionRateAmount,omitempty"`
	IyziCommissionFee            string         `json:"iyziCommissionFee,omitempty"`
	BasketId                     string         `json:"basketId,omitempty"`
	Currency                     string         `json:"currency,omitempty"`
	PaymentItems                 []*PaymentItem `json:"itemTransactions,omitempty"`
	Phase                        string         `json:"phase,omitempty"`
	AccountHolderName            string         `json:"accountHolderName,omitempty"`
	AccountNumber                string         `json:"accountNumber,omitempty"`
	BankName                     string         `json:"bankName,omitempty"`
	BankCode                     string         `json:"bankCode,omitempty"`
	Bic                          string         `json:"bic,omitempty"`
	PaymentPurpose               string         `json:"paymentPurpose,omitempty"`
	Iban                         string         `json:"iban,omitempty"`
	CountryCode                  string         `json:"countryCode,omitempty"`
	Apm                          string         `json:"amp,omitempty"`
	MobilePhone                  string         `json:"mobilePhone,omitempty"`
	PaymentStatus                string         `json:"paymentStatus,omitempty"`
}

func NewApmResource() *ApmResource {
	return &ApmResource{IyzipayResource: NewIyzipayResource()}
}
