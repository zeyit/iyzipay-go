package model

type CrossBookingFromSubMerchant struct {
	*IyzipayResource
}

func NewCrossBookingFromSubMerchant() *CrossBookingFromSubMerchant {
	return &CrossBookingFromSubMerchant{IyzipayResource: NewIyzipayResource()}
}
