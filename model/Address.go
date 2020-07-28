package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type Address struct {
	Description string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}

func NewAddress() *Address {
	return &Address{}
}

func (a *Address) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().Append("address", a.Description).
		Append("zipCode", a.ZipCode).
		Append("contactName", a.ContactName).
		Append("city", a.City).
		Append("country", a.Country).
		GetRequestString()

}
