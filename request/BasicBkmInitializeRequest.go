package request

import (
	. "github.com/zeyit/iyzipay-go/model"
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BasicBkmInitializeRequest struct {
	*BaseRequest
	ConnectorName      string            `json:"connectorName,omitempty"`
	Price              string            `json:"price,omitempty"`
	CallbackUrl        string            `json:"callbackUrl,omitempty"`
	BuyerEmail         string            `json:"buyerEmail,omitempty"`
	BuyerId            string            `json:"buyerId,omitempty"`
	BuyerIp            string            `json:"buyerIp,omitempty"`
	PosOrderId         string            `json:"posOrderId,omitempty"`
	InstallmentDetails []*BkmInstallment `json:"installmentDetails,omitempty"`
}

func NewBasicBkmInitializeRequest() *BasicBkmInitializeRequest {
	return &BasicBkmInitializeRequest{BaseRequest: NewBaseRequest()}
}

func (b *BasicBkmInitializeRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(b.BaseRequest.ToPKIRequestString()).
		Append("connectorName", b.ConnectorName).
		AppendPrice("price", b.Price).
		Append("callbackUrl", b.CallbackUrl).
		Append("buyerEmail", b.BuyerEmail).
		Append("buyerId", b.BuyerId).
		Append("buyerIp", b.BuyerIp).
		Append("posOrderId", b.PosOrderId).
		AppendList("installmentDetails", b.InstallmentDetails).
		GetRequestString()
}
