package model

type BouncedBankTransferList struct {
	*IyzipayResource
	BankTransfers []*BankTransfer `josn:"bouncedRows,omitempty"`
}

func NewBouncedBankTransferList() *BouncedBankTransferList {
	return &BouncedBankTransferList{IyzipayResource: NewIyzipayResource()}
}
