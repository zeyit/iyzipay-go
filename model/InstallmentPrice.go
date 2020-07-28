package model

type InstallmentPrice struct {
	Price             string `json:"InstallmentPrice,omitempty"`
	TotalPrice        string `json:"totalPrice,omitempty"`
	InstallmentNumber *int   `json:"installmentNumber,omitempty"`
}

func NewInstallmentPrice() *InstallmentPrice {
	return &InstallmentPrice{}
}
