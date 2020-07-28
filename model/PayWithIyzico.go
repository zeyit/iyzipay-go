package model

type PayWithIyzico struct {
	*PaymentResource
	Token       string `json:"token,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
}

func NewPayWithIyzico() *PayWithIyzico {
	return &PayWithIyzico{PaymentResource: NewPaymentResource()}
}
