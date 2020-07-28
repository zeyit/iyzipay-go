package model

import cnt "github.com/zeyit/iyzipay-go/requestconvert"

type InitialConsumer struct {
	Name        string           `json:"name,omitempty"`
	Surname     string           `json:"surname,omitempty"`
	Email       string           `json:"email,omitempty"`
	GsmNumber   string           `json:"gsmNumber,omitempty"`
	AddressList []*IyziupAddress `json:"addressList,omitempty"`
}

func NewInitialConsumer() *InitialConsumer {
	return &InitialConsumer{}
}

func (c *InitialConsumer) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("name", c.Name).
		Append("surname", c.Surname).
		Append("email", c.Email).
		Append("gsmNumber", c.GsmNumber).
		Append("addressList", c.AddressList).
		GetRequestString()
}
