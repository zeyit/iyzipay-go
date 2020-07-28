package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type BasketItem struct {
	Id               string `json:"id,omitempty"`
	Price            string `json:"price,omitempty"`
	Name             string `json:"name,omitempty"`
	Category1        string `json:"category1,omitempty"`
	Category2        string `json:"category2,omitempty"`
	ItemType         string `json:"itemType,omitempty"`
	SubMerchantKey   string `json:"subMerchantKey,omitempty"`
	SubMerchantPrice string `json:"subMerchantPrice,omitempty"`
}

func NewBasketItem() *BasketItem {
	return &BasketItem{}
}

func (b *BasketItem) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().Append("id", b.Id).
		AppendPrice("price", b.Price).
		Append("name", b.Name).
		Append("category1", b.Category1).
		Append("category2", b.Category2).
		Append("itemType", b.ItemType).
		Append("subMerchantKey", b.SubMerchantKey).
		AppendPrice("subMerchantPrice", b.SubMerchantPrice).
		GetRequestString()
}
