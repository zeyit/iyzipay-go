package request

import (
	. "github.com/zeyit/iyzipay-go/model"

	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CreateCardRequest struct {
	*BaseRequest
	ExternalId  string           `json:"externalId,omitempty"`
	Email       string           `json:"email,omitempty"`
	CardUserKey string           `json:"cardUserKey,omitempty"`
	Card        *CardInformation `json:"cart,omitempty"`
}

func NewCreateCardRequest() *CreateCardRequest {
	return &CreateCardRequest{BaseRequest: NewBaseRequest()}
}

func (c *CreateCardRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("externalId", c.ExternalId).
		Append("email", c.Email).
		Append("cardUserKey", c.CardUserKey).
		Append("card", c.Card).
		GetRequestString()
}
