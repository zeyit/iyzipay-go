package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveCardManagementPageCardRequest struct {
	*BaseRequest
	PageToken string `json:"pageToken,omitempty"`
}

func NewRetrieveCardManagementPageCardRequest() *RetrieveCardManagementPageCardRequest {
	return &RetrieveCardManagementPageCardRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveCardManagementPageCardRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("token", r.PageToken).
		GetRequestString()
}
