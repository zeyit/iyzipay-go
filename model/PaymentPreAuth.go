package model

type PaymentPreAuth struct {
	*PaymentResource
}

func NewPaymentPreAuth() *PaymentPreAuth {
	return &PaymentPreAuth{PaymentResource: NewPaymentResource()}
}
