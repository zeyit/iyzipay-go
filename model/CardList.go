package model

type CardList struct {
	*IyzipayResource

	CardUserKey string  `json:"cardUserKey,omitempty"`
	CardDetails []*Card `json:"cardDetails,omitempty"`
}

func NewCardList() *CardList {
	return &CardList{IyzipayResource: NewIyzipayResource()}
}
