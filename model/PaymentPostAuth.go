package model

type PaymentPostAuth struct {
	*PaymentResource
}

func NewPaymentPostAuth() *PaymentPostAuth {
	return &PaymentPostAuth{PaymentResource: NewPaymentResource()}
}
