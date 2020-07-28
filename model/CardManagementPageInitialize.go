package model

type CardManagementPageInitialize struct {
	*IyzipayResource

	ExternalId  string `json:"externalId,omitempty"`
	Token       string `json:"token,omitempty"`
	CardPageUrl string `json:"cardPageUrl,omitempty"`
}

func NewCardManagementPageInitialize() *CardManagementPageInitialize {
	return &CardManagementPageInitialize{IyzipayResource: NewIyzipayResource()}
}
