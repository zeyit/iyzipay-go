package model

type BasicPaymentPreAuth struct {
	*BasicPaymentResource
}

func NewBasicPaymentPreAuth() *BasicPaymentPreAuth {
	return &BasicPaymentPreAuth{BasicPaymentResource: NewBasicPaymentResource()}
}
