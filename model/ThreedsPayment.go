package model

type ThreedsPayment struct {
	*PaymentResource
}

func NewThreedsPayment() *ThreedsPayment {
	return &ThreedsPayment{PaymentResource: NewPaymentResource()}
}
