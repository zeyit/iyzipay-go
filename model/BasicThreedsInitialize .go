package model

type BasicThreedsInitialize struct {
	*IyzipayResource
	HtmlContent string `json:"threeDSHtmlContent,omitempty"`
}

func NewBasicThreedsInitialize() *BasicThreedsInitialize {
	return &BasicThreedsInitialize{IyzipayResource: NewIyzipayResource()}
}
