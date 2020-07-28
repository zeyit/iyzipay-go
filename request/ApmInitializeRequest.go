package request

import (
	. "github.com/zeyit/iyzipay-go/model"
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type ApmInitializeRequest struct {
	*BaseRequest
	Price                   string        `json:"price,omitempty"`
	PaidPrice               string        `json:"paidPrice,omitempty"`
	PaymentChannel          string        `json:"paymentChannel,omitempty"`
	PaymentGroup            string        `json:"paymentGroup,omitempty"`
	PaymentSource           string        `json:"paymentSource,omitempty"`
	Currency                string        `json:"currency,omitempty"`
	BasketId                string        `json:"basketId,omitempty"`
	MerchantOrderId         string        `json:"merchantOrderId,omitempty"`
	CountryCode             string        `json:"countryCode,omitempty"`
	AccountHolderName       string        `json:"accountHolderName,omitempty"`
	MerchantCallbackUrl     string        `json:"merchantCallbackUrl,omitempty"`
	MerchantErrorUrl        string        `json:"merchantErrorUrl,omitempty"`
	MerchantNotificationUrl string        `json:"merchantNotificationUrl,omitempty"`
	ApmType                 string        `json:"apmType,omitempty"`
	Buyer                   *Buyer        `json:"buyer,omitempty"`
	ShippingAddress         *Address      `json:"shippingAddress,omitempty"`
	BillingAddress          *Address      `json:"billingAddress,omitempty"`
	BasketItems             []*BasketItem `json:"basketItems,omitempty"`
}

func NewApmInitializeRequest() *ApmInitializeRequest {
	return &ApmInitializeRequest{BaseRequest: &BaseRequest{}}
}

func (a *ApmInitializeRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(a.BaseRequest.ToPKIRequestString()).
		AppendPrice("price", a.Price).
		AppendPrice("paidPrice", a.PaidPrice).
		Append("paymentChannel", a.PaymentChannel).
		Append("paymentGroup", a.PaymentGroup).
		Append("paymentSource", a.PaymentSource).
		Append("currency", a.Currency).
		Append("merchantOrderId", a.MerchantOrderId).
		Append("countryCode", a.CountryCode).
		Append("accountHolderName", a.AccountHolderName).
		Append("merchantCallbackUrl", a.MerchantCallbackUrl).
		Append("merchantErrorUrl", a.MerchantErrorUrl).
		Append("merchantNotificationUrl", a.MerchantNotificationUrl).
		Append("apmType", a.ApmType).
		Append("basketId", a.BasketId).
		Append("buyer", a.Buyer).
		Append("shippingAddress", a.ShippingAddress).
		Append("billingAddress", a.BillingAddress).
		AppendList("basketItems", a.BasketItems).
		GetRequestString()
}
