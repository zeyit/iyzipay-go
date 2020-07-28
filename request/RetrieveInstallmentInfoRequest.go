package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveInstallmentInfoRequest struct {
	*BaseRequest
	BinNumber string `json:"binNumber,omitempty"`
	Price     string `json:"price,omitempty"`
}

func NewRetrieveInstallmentInfoRequest() *RetrieveInstallmentInfoRequest {
	return &RetrieveInstallmentInfoRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveInstallmentInfoRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("binNumber", r.BinNumber).
		AppendPrice("price", r.Price).
		GetRequestString()
}
