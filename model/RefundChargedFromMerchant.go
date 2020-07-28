package model

type RefundChargedFromMerchant struct {
	*IyzipayResource
	PaymentId            string `json:"paymentId,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	Price                string `json:"price,omitempty"`
	AuthCode             string `json:"authCode,omitempty"`
	HostReference        string `json:"hostReference,omitempty"`
}

func NewRefundChargedFromMerchant() *RefundChargedFromMerchant {
	return &RefundChargedFromMerchant{IyzipayResource: NewIyzipayResource()}
}
