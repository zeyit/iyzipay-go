package model

type CheckoutForm struct {
	*PaymentResource
	Token       string `json:"token,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
}

func NewCheckoutForm() *CheckoutForm {
	return &CheckoutForm{PaymentResource: NewPaymentResource()}
}
