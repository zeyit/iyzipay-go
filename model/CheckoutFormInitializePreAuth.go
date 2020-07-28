package model

type CheckoutFormInitializePreAuth struct {
	*CheckoutFormInitializeResource
}

func NewCheckoutFormInitializePreAuth() *CheckoutFormInitializePreAuth {
	return &CheckoutFormInitializePreAuth{
		CheckoutFormInitializeResource: NewCheckoutFormInitializeResource(),
	}
}
