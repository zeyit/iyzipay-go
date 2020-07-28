package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveCheckoutFormRequest struct {
	*BaseRequest
	Token string `json:"token,omitempty"`
}

func NewRetrieveCheckoutFormRequest() *RetrieveCheckoutFormRequest {
	return &RetrieveCheckoutFormRequest{BaseRequest: NewBaseRequest()}
}

func (c *RetrieveCheckoutFormRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("token", c.Token).
		GetRequestString()
}
