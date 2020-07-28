package model

type PayWithIyzicoInitializeResource struct {
	*IyzipayResource
	Token                string `json:"token,omitempty"`
	CheckoutFormContent  string `json:"checkoutFormContent,omitempty"`
	TokenExpireTime      *int64 `json:"tokenExpireTime,omitempty"`
	PayWithIyzicoPageUrl string `json:"payWithIyzicoPageUrl,omitempty"`
}

func NewPayWithIyzicoInitializeResource() *PayWithIyzicoInitializeResource {
	return &PayWithIyzicoInitializeResource{IyzipayResource: NewIyzipayResource()}
}
