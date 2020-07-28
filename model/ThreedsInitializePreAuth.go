package model

type ThreedsInitializePreAuth struct {
	*IyzipayResource
	HtmlContent string `json:"threeDSHtmlContent,omitempty"`
}

func NewThreedsInitializePreAuth() *ThreedsInitializePreAuth {
	return &ThreedsInitializePreAuth{IyzipayResource: NewIyzipayResource()}
}
