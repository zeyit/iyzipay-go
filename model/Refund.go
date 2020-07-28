package model

type Refund struct {
	*IyzipayResource

	PaymentId            string `json:"paymentId,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	Price                string `json:"price,omitempty"`
	Currency             string `json:"currency,omitempty"`
	ConnectorName        string `json:"connectorName,omitempty"`
	AuthCode             string `json:"authCode,omitempty"`
	HostReference        string `json:"hostReference,omitempty"`
}

func NewRefund() *Refund {
	return &Refund{IyzipayResource: NewIyzipayResource()}
}
