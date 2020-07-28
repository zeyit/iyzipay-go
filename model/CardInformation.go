package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CardInformation struct {
	CardAlias      string `json:"cardAlias,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	CardHolderName string `json:"cardHolderName,omitempty"`
}

func NewCardInformation() *CardInformation {
	return &CardInformation{}
}

func (c *CardInformation) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("cardAlias", c.CardAlias).
		Append("cardNumber", c.CardNumber).
		Append("expireYear", c.ExpireYear).
		Append("expireMonth", c.ExpireMonth).
		Append("cardHolderName", c.CardHolderName).
		GetRequestString()
}
