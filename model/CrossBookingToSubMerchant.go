package model

type CrossBookingToSubMerchant struct {
	*IyzipayResource
}

func NewCrossBookingToSubMerchant() *CrossBookingToSubMerchant {
	return &CrossBookingToSubMerchant{IyzipayResource: NewIyzipayResource()}
}
