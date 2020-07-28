package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type PaymentCard struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	Cvc            string `json:"cvc,omitempty"`
	RegisterCard   *int   `json:"registerCard,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardToken      string `json:"cardToken,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

func NewPaymentCard() *PaymentCard {
	return &PaymentCard{}
}

func (p *PaymentCard) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("cardHolderName", p.CardHolderName).
		Append("cardNumber", p.CardNumber).
		Append("expireYear", p.ExpireYear).
		Append("expireMonth", p.ExpireMonth).
		Append("cvc", p.Cvc).
		Append("registerCard", p.RegisterCard).
		Append("cardAlias", p.CardAlias).
		Append("cardToken", p.CardToken).
		Append("cardUserKey", p.CardUserKey).
		GetRequestString()
}
