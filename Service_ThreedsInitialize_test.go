package iyzipay

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/zeyit/iyzipay-go/model"
	request "github.com/zeyit/iyzipay-go/request"
)

var options = &Options{
	APIKey:    "sandbox-8VtewP2jUBIohUMXClKlslCAxE4uK3NQ",
	SecretKey: "sandbox-cH1vJbVcdSFk9pvdzlALLpyoopQjdgB6",
	BaseURL:   "https://sandbox-api.iyzipay.com",
}

func PrintResponse(i interface{}) {
	if json, err := json.Marshal(i); err == nil {
		fmt.Println("Json:", string(json))
	} else {
		fmt.Println("Hata:", err)
	}
}

func TestThreedsInitialize(t *testing.T) {
	req := request.NewCreatePaymentRequest()

	req.Locale = LocaleTR
	req.ConversationId = "16"
	req.Price = "69.00"
	req.PaidPrice = "69.00"
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
	buyer.RegistrationAddress = "Test Adress  Ataşehir/İstanbul"
	buyer.Ip = "85.107.193.143"
	buyer.City = "İstanbul"
	buyer.Country = "Turkey"
	buyer.IdentityNumber = "00016"
	req.Buyer = buyer
	req.BillingAddress = NewAddress()
	req.BillingAddress.ContactName = "aa  bb"
	req.BillingAddress.City = "İstanbul"
	req.BillingAddress.Country = "Turkey"
	req.BillingAddress.Description = "snd ndmnsdnsd"
	req.ShippingAddress = req.BillingAddress

	basketItem := NewBasketItem()
	basketItem.Id = "19"
	basketItem.Name = "Pure Pineapple"
	basketItem.Category1 = "Kadın"
	basketItem.ItemType = "PHYSICAL"
	basketItem.Price = "68.00"

	basketItem2 := NewBasketItem()
	basketItem2.Id = "19"
	basketItem2.Name = "Pure Pineapple2"
	basketItem2.Category1 = "Kadın"
	basketItem2.ItemType = "PHYSICAL"
	basketItem2.Price = "1.00"
	req.BasketItems = append(req.BasketItems, basketItem, basketItem2)

	threedsInitialize, _ := NewService().ThreedsInitialize(req, options)
	PrintResponse(threedsInitialize)

	assert := assert.New(t)

	assert.Equal(StatusSUCCESS, threedsInitialize.Status)
	assert.Equal("16", threedsInitialize.ConversationId)

	assert.NotEmpty(threedsInitialize.SystemTime)
	assert.Empty(threedsInitialize.ErrorCode)
	assert.Empty(threedsInitialize.ErrorMessage)
	assert.Empty(threedsInitialize.ErrorGroup)

}
