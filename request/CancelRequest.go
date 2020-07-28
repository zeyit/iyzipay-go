package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CancelRequest struct {
	*BaseRequest
	PaymentId   string `json:"paymentId,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewCancelRequest() *CancelRequest {
	return &CancelRequest{BaseRequest: NewBaseRequest()}
}

func (c *CancelRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("paymentId", c.PaymentId).
		Append("ip", c.Ip).
		Append("reason", c.Reason).
		Append("description", c.Description).
		GetRequestString()
}
