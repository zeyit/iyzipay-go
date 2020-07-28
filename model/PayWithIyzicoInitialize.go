package model

type PayWithIyzicoInitialize struct {
	*PayWithIyzicoInitializeResource
}

func NewPayWithIyzicoInitialize() *PayWithIyzicoInitialize {
	return &PayWithIyzicoInitialize{PayWithIyzicoInitializeResource: NewPayWithIyzicoInitializeResource()}
}
