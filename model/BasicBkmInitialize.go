package model

type BasicBkmInitialize struct {
	*IyzipayResource
	HtmlContent string `json:"htmlContent,omitempty"`
	Token       string `json:"token,omitempty"`
}

func NewBasicBkmInitialize() *BasicBkmInitialize {
	return &BasicBkmInitialize{IyzipayResource: NewIyzipayResource()}
}
