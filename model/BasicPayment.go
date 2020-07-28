package model

type BasicPayment struct {
	*BasicPaymentResource
}

func NewBasicPayment() *BasicPayment {
	return &BasicPayment{BasicPaymentResource: NewBasicPaymentResource()}
}
