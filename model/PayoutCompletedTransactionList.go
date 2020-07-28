package model

type PayoutCompletedTransactionList struct {
	*IyzipayResource
	PayoutCompletedTransactions []*PayoutCompletedTransaction `json:"payoutCompletedTransactions,omitempty"`
}

func NewPayoutCompletedTransactionList() *PayoutCompletedTransactionList {
	return &PayoutCompletedTransactionList{IyzipayResource: NewIyzipayResource()}
}
