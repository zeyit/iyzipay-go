package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type ApprovalRequest struct {
	*BaseRequest
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
}

func NewApprovalRequest() *ApprovalRequest {
	return &ApprovalRequest{BaseRequest: &BaseRequest{}}
}

func (a *ApprovalRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(a.BaseRequest.ToPKIRequestString()).
		Append("paymentTransactionId", a.PaymentTransactionId).
		GetRequestString()
}
