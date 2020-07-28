package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type PaymentItemRequest struct {
	*BaseRequest
	SubMerchantKey       string `json:"subMerchantKey,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	SubMerchantPrice     string `json:"subMerchantPrice,omitempty"`
}

func NewPaymentItemRequest() *PaymentItemRequest {
	return &PaymentItemRequest{BaseRequest: NewBaseRequest()}
}

func (p *PaymentItemRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(p.BaseRequest.ToPKIRequestString()).
		Append("subMerchantKey", p.SubMerchantKey).
		Append("paymentTransactionId", p.PaymentTransactionId).
		Append("subMerchantPrice", p.SubMerchantPrice).
		GetRequestString()
}
