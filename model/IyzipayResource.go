package model

type IyzipayResource struct {
	Status         string `json:"status,omitempty"`
	ErrorCode      string `json:"errorCode,omitempty"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
	ErrorGroup     string `json:"errorGroup,omitempty"`
	Locale         string `json:"locale,omitempty"`
	SystemTime     int64  `json:"systemTime,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

func NewIyzipayResource() *IyzipayResource {
	return &IyzipayResource{}
}
