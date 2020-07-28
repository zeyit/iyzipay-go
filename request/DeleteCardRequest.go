package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type DeleteCardRequest struct {
	*BaseRequest
	CardUserKey string `json:"cardUserKey,omitempty"`
	CardToken   string `json:"cardToken,omitempty"`
}

func NewDeleteCardRequest() *DeleteCardRequest {
	return &DeleteCardRequest{BaseRequest: NewBaseRequest()}
}

func (c *DeleteCardRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("cardUserKey", c.CardUserKey).
		Append("cardToken", c.CardToken).
		GetRequestString()
}
