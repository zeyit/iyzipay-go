package request

import (
	. "github.com/zeyit/iyzipay-go/model"

	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

var single_installment = 1

type BasicPaymentRequest struct {
	*BaseRequest

	Price         string       `json:"price,omitempty"`
	PaidPrice     string       `json:"paidPrice,omitempty"`
	Installment   *int         `json:"installment,omitempty"`
	BuyerEmail    string       `json:"buyerEmail,omitempty"`
	BuyerId       string       `json:"buyerId,omitempty"`
	BuyerIp       string       `json:"buyerIp,omitempty"`
	PosOrderId    string       `json:"posOrderId,omitempty"`
	PaymentCard   *PaymentCard `json:"paymentCard,omitempty"`
	Currency      string       `json:"currency,omitempty"`
	ConnectorName string       `json:"connectorName,omitempty"`
	CallbackUrl   string       `json:"callbackUrl,omitempty"`
}

func NewBasicPaymentRequest() *BasicPaymentRequest {
	return &BasicPaymentRequest{BaseRequest: NewBaseRequest(), Installment: &single_installment}

}

func (b *BasicPaymentRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(b.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", b.Price).
		AppendPrice("paidPrice", b.PaidPrice).
		Append("installment", b.Installment).
		Append("buyerEmail", b.BuyerEmail).
		Append("buyerId", b.BuyerId).
		Append("buyerIp", b.BuyerIp).
		Append("posOrderId", b.PosOrderId).
		Append("paymentCard", b.PaymentCard).
		Append("currency", b.Currency).
		Append("connectorName", b.ConnectorName).
		Append("callbackUrl", b.CallbackUrl).
		GetRequestString()
}
