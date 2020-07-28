package model

type CardManagementPageCard struct {
	*IyzipayResource
	ExternalId string `json:"externalId,omitempty"`

	CardUserKey string  `json:"cardUserKey,omitempty"`
	CardDetails []*Card `json:"cardDetails,omitempty"`
}

func NewCardManagementPageCard() *CardManagementPageCard {
	return &CardManagementPageCard{IyzipayResource: NewIyzipayResource()}
}
