package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type PaymentPostAuthRequest struct {
	*BaseRequest
	PaymentId string `json:"paymentId,omitempty"`
	PaidPrice string `json:"paidPrice,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

func NewPaymentPostAuthRequest() *PaymentPostAuthRequest {
	return &PaymentPostAuthRequest{BaseRequest: NewBaseRequest()}
}

func (p *PaymentPostAuthRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(p.BaseRequest.ToPKIRequestString()).
		Append("paymentId", p.PaymentId).
		Append("ip", p.Ip).
		AppendPrice("paidPrice", p.PaidPrice).
		Append("currency", p.Currency).
		GetRequestString()
}
