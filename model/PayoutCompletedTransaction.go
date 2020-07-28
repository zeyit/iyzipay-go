package model

type PayoutCompletedTransaction struct {
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	PayoutAmount         string `json:"payoutAmount,omitempty"`
	PayoutType           string `json:"payoutType,omitempty"`
	SubMerchantKey       string `json:"subMerchantKey,omitempty"`
	Currency             string `json:"currency,omitempty"`
}

func NewPayoutCompletedTransaction() *PayoutCompletedTransaction {
	return &PayoutCompletedTransaction{}
}
