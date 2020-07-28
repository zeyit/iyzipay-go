package model

type PaymentItem struct {
	*IyzipayResource
	ItemId                        string           `json:"itemId,omitempty"`
	PaymentTransactionId          string           `json:"paymentTransactionId,omitempty"`
	TransactionStatus             *int             `json:"transactionStatus,omitempty"`
	Price                         string           `json:"price,omitempty"`
	PaidPrice                     string           `json:"paidPrice,omitempty"`
	MerchantCommissionRate        string           `json:"merchantCommissionRate,omitempty"`
	MerchantCommissionRateAmount  string           `json:"merchantCommissionRateAmount,omitempty"`
	IyziCommissionRateAmount      string           `json:"iyziCommissionRateAmount,omitempty"`
	IyziCommissionFee             string           `json:"iyziCommissionFee,omitempty"`
	BlockageRate                  string           `json:"blockageRate,omitempty"`
	BlockageRateAmountMerchant    string           `json:"blockageRateAmountMerchant,omitempty"`
	BlockageRateAmountSubMerchant string           `json:"blockageRateAmountSubMerchant,omitempty"`
	BlockageResolvedDate          string           `json:"blockageResolvedDate,omitempty"`
	SubMerchantKey                string           `json:"subMerchantKey,omitempty"`
	SubMerchantPrice              string           `json:"subMerchantPrice,omitempty"`
	SubMerchantPayoutRate         string           `json:"subMerchantPayoutRate,omitempty"`
	SubMerchantPayoutAmount       string           `json:"subMerchantPayoutAmount,omitempty"`
	MerchantPayoutAmount          string           `json:"merchantPayoutAmount,omitempty"`
	ConvertedPayout               *ConvertedPayout `json:"convertedPayout,omitempty"`
}

func NewPaymentItem() *PaymentItem {
	return &PaymentItem{IyzipayResource: NewIyzipayResource()}
}
