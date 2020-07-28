package request

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BaseRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

func NewBaseRequest() *BaseRequest {
	return &BaseRequest{}
}

func (b *BaseRequest) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("locale", b.Locale).
		Append("conversationId", b.ConversationId).
		GetRequestString()

}
