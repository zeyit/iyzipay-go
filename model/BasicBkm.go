package model

type BasicBkm struct {
	*BasicPaymentResource
	Token         string `json:"token,omitempty"`
	CallbackUrl   string `json:"callbackUrl,omitempty"`
	PaymentStatus string `json:"paymentStatus,omitempty"`
}

func NewBasicBkm() *BasicBkm {
	return &BasicBkm{BasicPaymentResource: NewBasicPaymentResource()}
}
