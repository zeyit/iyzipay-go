package iyzipay

import (
	"testing"

	. "github.com/zeyit/iyzipay-go/model"
	request "github.com/zeyit/iyzipay-go/request"
)

type testRequest struct{}

func (t *testRequest) ToPKIRequestString() string {
	return "[data=value]"
}

func TestgenerateHash(t *testing.T) {
	expectedHash := "Cy84UuLZpfGhI7oaPD0Ckx1M0mo="
	generatedHash := generateHash("apiKey", "secretKey", "random", &testRequest{})
	if expectedHash != generatedHash {
		t.Error("Geçersi hash")
	}
}

func TestgenerateHash2(t *testing.T) {
	req := request.NewCreatePaymentRequest()

	req.Locale = LocaleTR
	req.ConversationId = "16"
	req.Price = "68.00"
	req.PaidPrice = "68.00"
	req.Currency = CurrencyTRY
	req.BasketId = "16"
	req.BasketId = "16"
	req.PaymentGroup = PaymentGroupLISTING
	req.CallbackUrl = "https://945f2c80b344.ngrok.io/cart/paymentinfo/16"

	card := NewPaymentCard()
	card.CardHolderName = "John Doe"
	card.CardNumber = "5528790000000008"
	card.ExpireMonth = "12"
	card.ExpireYear = "2030"
	card.Cvc = "123"
	registerCard := 0
	card.RegisterCard = &registerCard
	req.PaymentCard = card
	buyer := NewBuyer()
	buyer.Id = "16"
	buyer.Name = "aa"
	buyer.Surname = "bb"
	buyer.GsmNumber = "+905555555555"
	buyer.Email = "asd@qwe.com"
	buyer.RegistrationAddress = "İnönü mah Naeçiçeği sokak No:7  Ataşehir/İstanbul"
	buyer.Ip = "85.107.193.143"
	buyer.City = "İstanbul"
	buyer.Country = "Turkey"
	req.Buyer = buyer
	req.BillingAddress = NewAddress()
	req.BillingAddress.ContactName = "aa  bb"
	req.BillingAddress.City = "İstanbul"
	req.BillingAddress.Country = "Turkey"
	req.ShippingAddress = req.BillingAddress

	basketItem := NewBasketItem()
	basketItem.Id = "19"
	basketItem.Name = "Pure Pineapple"
	basketItem.Category1 = "Kadın"
	basketItem.ItemType = "PHYSICAL"
	basketItem.Price = "68.00"
	req.BasketItems = append(req.BasketItems, basketItem)

	expectedHash := "p9dd+WTGLE9aDUIvkzKK9ZkyJxo="
	generatedHash := generateHash("sandbox-8VtewP2jUBIohUMXClKlslCAxE4uK3NQ",
		"sandbox-cH1vJbVcdSFk9pvdzlALLpyoopQjdgB6", "270720201146255585", req)
	if expectedHash != generatedHash {
		t.Error("Geçersi hash")
	}
}
