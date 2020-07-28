package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveTransactionsRequest struct {
	*BaseRequest
	Date string `json:"date,omitempty"`
}

func NewRetrieveTransactionsRequest() *RetrieveTransactionsRequest {
	return &RetrieveTransactionsRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveTransactionsRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("date", r.Date).
		GetRequestString()
}
