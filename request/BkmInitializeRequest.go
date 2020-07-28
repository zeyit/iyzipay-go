package request

import (
	. "github.com/zeyit/iyzipay-go/model"

	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BkmInitializeRequest struct {
	*BaseRequest
	Price           string        `json:"price,omitempty"`
	BasketId        string        `json:"basketId,omitempty"`
	PaymentGroup    string        `json:"paymentGroup,omitempty"`
	PaymentSource   string        `json:"paymentSource,omitempty"`
	Buyer           *Buyer        `json:"buyer,omitempty"`
	ShippingAddress *Address      `json:"shippingAddress,omitempty"`
	BillingAddress  *Address      `json:"billingAddress,omitempty"`
	BasketItems     []*BasketItem `json:"basketItems,omitempty"`
	CallbackUrl     string        `json:"callbackUrl,omitempty"`
}

func NewBkmInitializeRequest() *BkmInitializeRequest {
	return &BkmInitializeRequest{BaseRequest: NewBaseRequest()}
}

func (b *BkmInitializeRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(b.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", b.Price).
		Append("basketId", b.BasketId).
		Append("paymentGroup", b.PaymentGroup).
		Append("buyer", b.Buyer).
		Append("shippingAddress", b.ShippingAddress).
		Append("billingAddress", b.BillingAddress).
		AppendList("basketItems", b.BasketItems).
		Append("callbackUrl", b.CallbackUrl).
		Append("paymentSource", b.PaymentSource).
		GetRequestString()
}
