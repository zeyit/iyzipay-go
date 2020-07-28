package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type CardManagementPageInitializeRequest struct {
	*BaseRequest
	AddNewCardEnabled bool   `json:"addNewCardEnabled,omitempty"`
	ValidateNewCard   bool   `json:"validateNewCard,omitempty"`
	ExternalId        string `json:"externalId,omitempty"`
	Email             string `json:"email,omitempty"`
	CardUserKey       string `json:"cardUserKey,omitempty"`
	CallbackUrl       string `json:"callbackUrl,omitempty"`
	DebitCardAllowed  bool   `json:"debitCardAllowed,omitempty"`
}

func NewCardManagementPageInitializeRequest() *CardManagementPageInitializeRequest {
	return &CardManagementPageInitializeRequest{BaseRequest: NewBaseRequest()}
}

func (c *CardManagementPageInitializeRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		AppendSuper(c.BaseRequest.ToPKIRequestString()).
		Append("addNewCardEnabled", c.AddNewCardEnabled).
		Append("validateNewCard", c.ValidateNewCard).
		Append("externalId", c.ExternalId).
		Append("email", c.Email).
		Append("cardUserKey", c.CardUserKey).
		Append("callbackUrl", c.CallbackUrl).
		Append("debitCardAllowed", c.DebitCardAllowed).
		GetRequestString()
}
