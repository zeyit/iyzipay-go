package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrievePayWithIyzicoRequest struct {
	*BaseRequest
	Token string `json:"token,omitempty"`
}

func NewRetrievePayWithIyzicoRequest() *RetrievePayWithIyzicoRequest {
	return &RetrievePayWithIyzicoRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrievePayWithIyzicoRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("token", r.Token).
		GetRequestString()
}
