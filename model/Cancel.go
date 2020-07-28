package model

type Cancel struct {
	*IyzipayResource

	PaymentId     string `json:"paymentId,omitempty"`
	Price         string `json:"price,omitempty"`
	Currency      string `json:"currency,omitempty"`
	ConnectorName string `json:"connectorName,omitempty"`
	AuthCode      string `json:"authCode,omitempty"`
	HostReference string `json:"hostReference,omitempty"`
}

func NewCancel() *Cancel {
	return &Cancel{IyzipayResource: NewIyzipayResource()}
}
