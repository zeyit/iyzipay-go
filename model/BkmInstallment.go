package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BkmInstallment struct {
	BankId            *int64                 `json:"bankId,omitempty"`
	InstallmentPrices []*BkmInstallmentPrice `json:"installmentPrices,omitempty"`
}

func NewBkmInstallment() *BkmInstallment {
	return &BkmInstallment{}
}

func (b *BkmInstallment) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("bankId", b.BankId).
		AppendList("installmentPrices", b.InstallmentPrices).
		GetRequestString()
}
