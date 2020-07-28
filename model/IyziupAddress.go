package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type IyziupAddress struct {
	Alias        string `json:"alias,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	ZipCode      string `json:"zipCode,omitempty"`
	ContactName  string `json:"contactName,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
}

func NewIyziupAddress() *IyziupAddress {
	return &IyziupAddress{}
}

func (i *IyziupAddress) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("alias", i.Alias).
		Append("addressLine1", i.AddressLine1).
		Append("addressLine2", i.AddressLine2).
		Append("zipCode", i.ZipCode).
		Append("contactName", i.ContactName).
		Append("city", i.City).
		Append("country", i.Country).
		GetRequestString()
}
