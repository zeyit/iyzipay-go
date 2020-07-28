package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrievePaymentRequest struct {
	*BaseRequest
	PaymentId             string `json:"paymentId,omitempty"`
	PaymentConversationId string `json:"paymentConversationId,omitempty"`
}

func NewRetrievePaymentRequest() *RetrievePaymentRequest {
	return &RetrievePaymentRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrievePaymentRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("paymentId", r.PaymentId).
		Append("paymentConversationId", r.PaymentConversationId).
		GetRequestString()
}
