package model

type BkmInitialize struct {
	*IyzipayResource
	HtmlContent string `json:"htmlContent,omitempty"`
	Token       string `json:"token,omitempty"`
}

func NewBkmInitialize() *BkmInitialize {
	return &BkmInitialize{IyzipayResource: NewIyzipayResource()}
}
