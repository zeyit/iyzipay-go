package model

type ThreedsInitialize struct {
	*IyzipayResource
	HtmlContent string `json:"threeDSHtmlContent,omitempty"`
}

func NewThreedsInitialize() *ThreedsInitialize {
	return &ThreedsInitialize{IyzipayResource: NewIyzipayResource()}
}
