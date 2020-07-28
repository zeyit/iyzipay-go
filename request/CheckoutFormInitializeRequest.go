package request

import (
	. "github.com/zeyit/iyzipay-go/model"
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CheckoutFormInitializeRequest struct {
	*BaseRequest
	Price               string        `json:"price,omitempty"`
	PaidPrice           string        `json:"paidPrice,omitempty"`
	BasketId            string        `json:"basketId,omitempty"`
	PaymentGroup        string        `json:"paymentGroup,omitempty"`
	PaymentSource       string        `json:"paymentSource,omitempty"`
	Currency            string        `json:"currency,omitempty"`
	Buyer               *Buyer        `json:"buyer,omitempty"`
	ShippingAddress     *Address      `json:"shippingAddress,omitempty"`
	BillingAddress      *Address      `json:"billingAddress,omitempty"`
	BasketItems         []*BasketItem `json:"basketItems,omitempty"`
	CallbackUrl         string        `json:"callbackUrl,omitempty"`
	ForceThreeDS        *int          `json:"forceThreeDS,omitempty"`
	CardUserKey         string        `json:"cardUserKey,omitempty"`
	PosOrderId          string        `json:"oosOrderId,omitempty"`
	EnabledInstallments []int         `json:"enabledInstallments,omitempty"`
}

func NewCheckoutFormInitializeRequest() *CheckoutFormInitializeRequest {
	return &CheckoutFormInitializeRequest{BaseRequest: NewBaseRequest()}
}

func (c *CheckoutFormInitializeRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", c.Price).
		Append("basketId", c.BasketId).
		Append("paymentGroup", c.PaymentGroup).
		Append("buyer", c.Buyer).
		Append("shippingAddress", c.ShippingAddress).
		Append("billingAddress", c.BillingAddress).
		AppendList("basketItems", c.BasketItems).
		Append("callbackUrl", c.CallbackUrl).
		Append("paymentSource", c.PaymentSource).
		Append("currency", c.Currency).
		Append("posOrderId", c.PosOrderId).
		AppendPrice("paidPrice", c.PaidPrice).
		Append("forceThreeDS", c.ForceThreeDS).
		Append("cardUserKey", c.CardUserKey).
		AppendList("enabledInstallments", c.EnabledInstallments).
		GetRequestString()
}
