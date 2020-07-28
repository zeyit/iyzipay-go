package request

import cnt "github.com/zeyit/iyzipay-go/requestconvert"

type CrossBookingRequest struct {
	*BaseRequest
	SubMerchantKey string `json:"subMerchantKey,omitempty"`
	Price          string `json:"price,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Currency       string `json:"currency,omitempty"`
}

func NewCrossBookingRequest() *CrossBookingRequest {
	return &CrossBookingRequest{BaseRequest: NewBaseRequest()}
}

func (c *CrossBookingRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("subMerchantKey", c.SubMerchantKey).
		AppendPrice("price", c.Price).
		Append("reason", c.Reason).
		Append("currency", c.Currency).
		GetRequestString()
}
