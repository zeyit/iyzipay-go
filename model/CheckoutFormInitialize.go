package model

type CheckoutFormInitialize struct {
	*CheckoutFormInitializeResource
}

func NewCheckoutFormInitialize() *CheckoutFormInitialize {
	return &CheckoutFormInitialize{
		CheckoutFormInitializeResource: NewCheckoutFormInitializeResource(),
	}
}
