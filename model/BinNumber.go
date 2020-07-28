package model

type BinNumber struct {
	*IyzipayResource

	Bin             string `json:"binNumber,omitempty"`
	CardType        string `json:"cardType,omitempty"`
	CardAssociation string `json:"cardAssociation,omitempty"`
	CardFamily      string `json:"cardFamily,omitempty"`
	BankName        string `json:"bankName,omitempty"`
	BankCode        int64  `json:"bankCode,omitempty"`
}

func NewBinNumber() *BinNumber {
	return &BinNumber{IyzipayResource: NewIyzipayResource()}
}
