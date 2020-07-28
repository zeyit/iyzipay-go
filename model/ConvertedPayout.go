package model

type ConvertedPayout struct {
	PaidPrice                     string `json:"paidPrice,omitempty"`
	IyziCommissionRateAmount      string `json:"iyziCommissionRateAmount,omitempty"`
	IyziCommissionFee             string `json:"iyziCommissionFee,omitempty"`
	BlockageRateAmountMerchant    string `json:"blockageRateAmountMerchant,omitempty"`
	BlockageRateAmountSubMerchant string `json:"blockageRateAmountSubMerchant,omitempty"`
	SubMerchantPayoutAmount       string `json:"subMerchantPayoutAmount,omitempty"`
	MerchantPayoutAmount          string `json:"merchantPayoutAmount,omitempty"`
	IyziConversionRate            string `json:"iyziConversionRate,omitempty"`
	IyziConversionRateAmount      string `json:"iyziConversionRateAmount,omitempty"`
	Currency                      string `json:"currency,omitempty"`
}

func NewConvertedPayout() *ConvertedPayout {
	return &ConvertedPayout{}
}
