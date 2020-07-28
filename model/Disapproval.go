package model

type Disapproval struct {
	*IyzipayResource
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
}

func NewDisapproval() *Disapproval {
	return &Disapproval{IyzipayResource: NewIyzipayResource()}
}
