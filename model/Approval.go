package model

type Approval struct {
	*IyzipayResource
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
}

func NewApproval() *Approval {
	return &Approval{IyzipayResource: NewIyzipayResource()}
}
