package model

type BasicPaymentResource struct {
	*IyzipayResource

	Price                        string `json:"price,omitempty"`
	PaidPrice                    string `json:"paidPrice,omitempty"`
	Installment                  *int   `json:"installment,omitempty"`
	Currency                     string `json:"currency,omitempty"`
	PaymentId                    string `json:"paymentId,omitempty"`
	MerchantCommissionRate       string `json:"merchantCommissionRate,omitempty"`
	MerchantCommissionRateAmount string `json:"merchantCommissionRateAmount,omitempty"`
	IyziCommissionFee            string `json:"iyziCommissionFee,omitempty"`
	CardType                     string `json:"cardType,omitempty"`
	CardAssociation              string `json:"cardAssociation,omitempty"`
	CardFamily                   string `json:"cardFamily,omitempty"`
	CardToken                    string `json:"cardToken,omitempty"`
	CardUserKey                  string `json:"cardUserKey,omitempty"`
	BinNumber                    string `json:"binNumber,omitempty"`
	PaymentTransactionId         string `json:"paymentTransactionId,omitempty"`
	AuthCode                     string `json:"authCode,omitempty"`
	ConnectorName                string `json:"connectorName,omitempty"`
	Phase                        string `json:"phase,omitempty"`
}

func NewBasicPaymentResource() *BasicPaymentResource {
	return &BasicPaymentResource{IyzipayResource: NewIyzipayResource()}
}
