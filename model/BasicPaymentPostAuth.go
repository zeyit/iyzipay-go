package model

type BasicPaymentPostAuth struct {
	*BasicPaymentResource
}

func NewBasicPaymentPostAuth() *BasicPaymentPostAuth {
	return &BasicPaymentPostAuth{BasicPaymentResource: NewBasicPaymentResource()}
}
