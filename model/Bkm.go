package model

type Bkm struct {
	*PaymentResource

	Token       string `json:"token,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
}

func NewBkm() *Bkm {
	return &Bkm{PaymentResource: NewPaymentResource()}
}
