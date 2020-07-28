package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CardListRequest struct {
	*BaseRequest
	CardUserKey string `json:"cardUserKey,omitempty"`
}

func NewCardListRequest() *CardListRequest {
	return &CardListRequest{BaseRequest: NewBaseRequest()}
}

func (c *CardListRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("cardUserKey", c.CardUserKey).
		GetRequestString()
}
