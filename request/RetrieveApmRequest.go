package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveApmRequest struct {
	*BaseRequest
	PaymentId string `json:"paymentId,omitempty"`
}

func NewRetrieveApmRequest() *RetrieveApmRequest {
	return &RetrieveApmRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveApmRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("paymentId", r.PaymentId).
		GetRequestString()
}
