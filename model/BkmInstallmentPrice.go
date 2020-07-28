package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BkmInstallmentPrice struct {
	InstallmentNumber *int   `json:"installmentNumber,omitempty"`
	TotalPrice        string `json:"totalPrice,omitempty"`
}

func NewBkmInstallmentPrice() *BkmInstallmentPrice {
	return &BkmInstallmentPrice{}
}

func (b *BkmInstallmentPrice) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("installmentNumber", b.InstallmentNumber).
		AppendPrice("totalPrice", b.TotalPrice).
		GetRequestString()
}
