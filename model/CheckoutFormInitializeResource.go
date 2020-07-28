package model

type CheckoutFormInitializeResource struct {
	*IyzipayResource
	Token               string `json:"token,omitempty"`
	CheckoutFormContent string `json:"checkoutFormContent,omitempty"`
	TokenExpireTime     *int64 `json:"tokenExpireTime,omitempty"`
	PaymentPageUrl      string `json:"paymentPageUrl,omitempty"`
}

func NewCheckoutFormInitializeResource() *CheckoutFormInitializeResource {
	return &CheckoutFormInitializeResource{IyzipayResource: NewIyzipayResource()}
}
