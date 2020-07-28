package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RefundRequest struct {
	*BaseRequest
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	Price                string `json:"price,omitempty"`
	Ip                   string `json:"ip,omitempty"`
	Currency             string `json:"currency,omitempty"`
	Reason               string `json:"reason,omitempty"`
	Description          string `json:"description,omitempty"`
}

func NewRefundRequest() *RefundRequest {
	return &RefundRequest{BaseRequest: NewBaseRequest()}
}

func (r *RefundRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("paymentTransactionId", r.PaymentTransactionId).
		AppendPrice("price", r.Price).
		Append("ip", r.Ip).
		Append("currency", r.Currency).
		Append("reason", r.Reason).
		Append("description", r.Description).
		GetRequestString()
}
