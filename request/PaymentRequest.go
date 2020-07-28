package request

import (
	. "github.com/zeyit/iyzipay-go/model"

	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type PaymentRequest struct {
	*BaseRequest
	Price           string        `json:"price,omitempty"`
	PaidPrice       string        `json:"paidPrice,omitempty"`
	Installment     int           `json:"installment,omitempty"`
	PaymentChannel  string        `json:"paymentChannel,omitempty"`
	BasketId        string        `json:"basketId,omitempty"`
	PaymentGroup    string        `json:"paymentGroup,omitempty"`
	PaymentCard     *PaymentCard  `json:"paymentCard,omitempty"`
	Buyer           *Buyer        `json:"buyer,omitempty"`
	ShippingAddress *Address      `json:"shippingAddress,omitempty"`
	BillingAddress  *Address      `json:"billingAddress,omitempty"`
	BasketItems     []*BasketItem `json:"basketItems,omitempty"`
	PaymentSource   string        `json:"paymentSource,omitempty"`
	CallbackUrl     string        `json:"callbackUrl,omitempty"`
	PosOrderId      string        `json:"posOrderId,omitempty"`
	ConnectorName   string        `json:"connectorName,omitempty"`
	Currency        string        `json:"currency,omitempty"`
}

func NewPaymentRequest() *PaymentRequest {
	return &PaymentRequest{BaseRequest: NewBaseRequest()}
}

func (p *PaymentRequest) ToPKIRequestString() string {

	return cnt.NewStrRequestBuilder().
		AppendSuper(p.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", p.Price).
		AppendPrice("paidPrice", p.PaidPrice).
		Append("installment", p.Installment).
		Append("paymentChannel", p.PaymentChannel).
		Append("basketId", p.BasketId).
		Append("paymentGroup", p.PaymentGroup).
		Append("paymentCard", p.PaymentCard).
		Append("buyer", p.Buyer).
		Append("shippingAddress", p.ShippingAddress).
		Append("billingAddress", p.BillingAddress).
		AppendList("basketItems", p.BasketItems).
		Append("paymentSource", p.PaymentSource).
		Append("currency", p.Currency).
		Append("posOrderId", p.PosOrderId).
		Append("connectorName", p.ConnectorName).
		Append("callbackUrl", p.CallbackUrl).
		GetRequestString()
}
