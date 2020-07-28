package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveBkmRequest struct {
	*BaseRequest
	Token string `json:"token,omitempty"`
}

func NewRetrieveBkmRequest() *RetrieveBkmRequest {
	return &RetrieveBkmRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveBkmRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("token", r.Token).
		GetRequestString()
}
