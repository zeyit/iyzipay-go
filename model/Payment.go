package model

type Payment struct {
	*PaymentResource
}

func NewPayment() *Payment {
	return &Payment{PaymentResource: NewPaymentResource()}
}
