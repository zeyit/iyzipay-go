package model

type BasicThreedsInitializePreAuth struct {
	*IyzipayResource
	HtmlContent string `json:"threeDSHtmlContent,omitempty"`
}

func NewBasicThreedsInitializePreAuth() *BasicThreedsInitializePreAuth {
	return &BasicThreedsInitializePreAuth{IyzipayResource: NewIyzipayResource()}
}
