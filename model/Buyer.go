package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type Buyer struct {
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Surname             string `json:"surname,omitempty"`
	IdentityNumber      string `json:"identityNumber,omitempty"`
	Email               string `json:"email,omitempty"`
	GsmNumber           string `json:"gsmNumber,omitempty"`
	RegistrationDate    string `json:"registrationDate,omitempty"`
	LastLoginDate       string `json:"lastLoginDate,omitempty"`
	RegistrationAddress string `json:"registrationAddress,omitempty"`
	City                string `json:"city,omitempty"`
	Country             string `json:"country,omitempty"`
	ZipCode             string `json:"zipCode,omitempty"`
	Ip                  string `json:"ip,omitempty"`
}

func NewBuyer() *Buyer {
	return &Buyer{}
}

func (b *Buyer) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().Append("id", b.Id).
		Append("name", b.Name).
		Append("surname", b.Surname).
		Append("identityNumber", b.IdentityNumber).
		Append("email", b.Email).
		Append("gsmNumber", b.GsmNumber).
		Append("registrationDate", b.RegistrationDate).
		Append("lastLoginDate", b.LastLoginDate).
		Append("registrationAddress", b.RegistrationAddress).
		Append("city", b.City).
		Append("country", b.Country).
		Append("zipCode", b.ZipCode).
		Append("ip", b.Ip).
		GetRequestString()

}
