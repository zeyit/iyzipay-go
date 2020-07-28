package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type RetrieveSubMerchantRequest struct {
	*BaseRequest
	SubMerchantExternalId string `json:"subMerchantExternalId,omitempty"`
}

func NewRetrieveSubMerchantRequest() *RetrieveSubMerchantRequest {
	return &RetrieveSubMerchantRequest{BaseRequest: NewBaseRequest()}
}

func (r *RetrieveSubMerchantRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(r.BaseRequest.ToPKIRequestString()).
		Append("subMerchantExternalId", r.SubMerchantExternalId).
		GetRequestString()
}
