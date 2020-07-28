package request

import (
	. "github.com/zeyit/iyzipay-go/model"
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CreatePaymentRequest struct {
	*BaseRequest
	Price           string        `json:"price,omitempty"`
	PaidPrice       string        `json:"paidPrice,omitempty"`
	Installment     *int          `json:"installment,omitempty"`
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

func NewCreatePaymentRequest() *CreatePaymentRequest {
	return &CreatePaymentRequest{BaseRequest: NewBaseRequest()}
}

func (c *CreatePaymentRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", c.Price).
		AppendPrice("paidPrice", c.PaidPrice).
		Append("installment", c.Installment).
		Append("paymentChannel", c.PaymentChannel).
		Append("basketId", c.BasketId).
		Append("paymentGroup", c.PaymentGroup).
		Append("paymentCard", c.PaymentCard).
		Append("buyer", c.Buyer).
		Append("shippingAddress", c.ShippingAddress).
		Append("billingAddress", c.BillingAddress).
		AppendList("basketItems", c.BasketItems).
		Append("paymentSource", c.PaymentSource).
		Append("currency", c.Currency).
		Append("posOrderId", c.PosOrderId).
		Append("connectorName", c.ConnectorName).
		Append("callbackUrl", c.CallbackUrl).
		GetRequestString()
}
