package model

type PaymentResource struct {
	*IyzipayResource

	Price                        int    `json:"price,omitempty"`
	PaidPrice                    int    `json:"paidPrice,omitempty"`
	Installment                  *int   `json:"installment,omitempty"`
	Currency                     string `json:"currency,omitempty"`
	PaymentId                    string `json:"paymentId,omitempty"`
	PaymentStatus                string `json:"paymentStatus,omitempty"`
	FraudStatus                  *int   `json:"fraudStatus,omitempty"`
	MerchantCommissionRate       string `json:"merchantCommissionRate,omitempty"`
	MerchantCommissionRateAmount string `json:"merchantCommissionRateAmount,omitempty"`
	IyziCommissionRateAmount     string `json:"iyziCommissionRateAmount,omitempty"`
	IyziCommissionFee            string `json:"iyziCommissionFee,omitempty"`
	CardType                     string `json:"cardType,omitempty"`
	CardAssociation              string `json:"cardAssociation,omitempty"`
	CardFamily                   string `json:"cardFamily,omitempty"`
	CardToken                    string `json:"cardToken,omitempty"`
	CardUserKey                  string `json:"cardUserKey,omitempty"`
	BinNumber                    string `json:"binNumber,omitempty"`
	LastFourDigits               string `json:"lastFourDigits,omitempty"`
	BasketId                     string `json:"basketId,omitempty"`

	PaymentItems  []*PaymentItem `json:"itemTransactions,omitempty"`
	ConnectorName string         `json:"connectorName,omitempty"`
	AuthCode      string         `json:"authCode,omitempty"`
	HostReference string         `json:"hostReference,omitempty"`
	Phase         string         `json:"phase,omitempty"`
}

func NewPaymentResource() *PaymentResource {
	return &PaymentResource{IyzipayResource: NewIyzipayResource()}
}
