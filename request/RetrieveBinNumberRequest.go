package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveBinNumberRequest struct {
	*BaseRequest
	BinNumber string `json:"binNumber,omitempty"`
}

func NewRetrieveBinNumberRequest() *RetrieveBinNumberRequest {
	return &RetrieveBinNumberRequest{BaseRequest: NewBaseRequest()}
}
func (b *RetrieveBinNumberRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(b.BaseRequest.ToPKIRequestString()).
		Append("binNumber", b.BinNumber).
		GetRequestString()
}
