package model

import (
	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

type OrderItem struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Category1       string `json:"category1,omitempty"`
	Category2       string `json:"category2,omitempty"`
	ItemType        string `json:"itemType,omitempty"`
	ItemUrl         string `json:"itemUrl,omitempty"`
	ItemDescription string `json:"itemDescription,omitempty"`
	Price           string `json:"price,omitempty"`
}

func NewOrderItem() *OrderItem {
	return &OrderItem{}
}
func (o *OrderItem) ToPKIRequestString() string {
	return cnt.NewStrRequestBuilder().
		Append("id", o.Id).
		Append("name", o.Name).
		Append("category1", o.Category1).
		Append("category2", o.Category2).
		Append("itemType", o.ItemType).
		Append("itemUrl", o.ItemUrl).
		Append("itemDescription", o.ItemDescription).
		AppendPrice("price", o.Price).
		GetRequestString()
}
