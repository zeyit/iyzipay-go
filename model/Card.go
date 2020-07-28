package model

type Card struct {
	*IyzipayResource

	ExternalId      string `json:"externalId,omitempty"`
	Email           string `json:"email,omitempty"`
	CardUserKey     string `json:"cardUserKey,omitempty"`
	CardToken       string `json:"cardToken,omitempty"`
	CardAlias       string `json:"cardAlias,omitempty"`
	BinNumber       string `json:"binNumber,omitempty"`
	CardType        string `json:"cardType,omitempty"`
	CardAssociation string `json:"cardAssociation,omitempty"`
	CardFamily      string `json:"cardFamily,omitempty"`
	CardBankCode    *int64 `json:"cardBankCode,omitempty"`
	CardBankName    string `json:"cardBankName,omitempty"`
}

func NewCard() *Card {
	return &Card{IyzipayResource: NewIyzipayResource()}
}
