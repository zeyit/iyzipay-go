package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type ThreedsPaymentRequest struct {
	*BaseRequest
	PaymentId        string `json:"paymentId,omitempty"`
	ConversationData string `json:"conversationData,omitempty"`
}

func NewThreedsPaymentRequest() *ThreedsPaymentRequest {
	return &ThreedsPaymentRequest{BaseRequest: NewBaseRequest()}
}

func (t *ThreedsPaymentRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(t.BaseRequest.ToPKIRequestString()).
		Append("paymentId", t.PaymentId).
		Append("conversationData", t.ConversationData).
		GetRequestString()
}
