package model

type BasicThreedsPayment struct {
	*BasicPaymentResource
}

func NewBasicThreedsPayment() *BasicThreedsPayment {
	return &BasicThreedsPayment{BasicPaymentResource: NewBasicPaymentResource()}
}
