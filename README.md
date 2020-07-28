# iyzipay-go  [![Build Status](https://travis-ci.org/zeyit/iyzipay-go.svg?branch=master)](https://travis-ci.org/github/zeyit/iyzipay-go)

You can sign up for an iyzico account at https://iyzico.com

# Installation
  ```bash
$ go get github.com/zeyit/iyzipay-go
```
# Usage
  ```golang


	req := request.NewCreatePaymentRequest()

	req.Locale = iyzipay.LocaleTR
	req.ConversationId = "123456789"
	req.Price = "1"
	req.PaidPrice = "1.2"
	req.Currency = iyzipay.CurrencyTRY
	installment := 1
	req.Installment = &installment
	req.BasketId = "B67832"
	req.PaymentChannel = iyzipay.PaymentChannelWEB
	req.PaymentGroup = iyzipay.PaymentGroupPRODUCT
	req.CallbackUrl = "https://merchant-callback.com"

	card := model.NewPaymentCard()
	card.CardHolderName = "John Doe"
	card.CardNumber = "5528790000000008"
	card.ExpireMonth = "12"
	card.ExpireYear = "2030"
	card.Cvc = "123"
	registerCard := 0
	card.RegisterCard = &registerCard
	req.PaymentCard = card

	buyer := model.NewBuyer()
	buyer.Id = "BY789"
	buyer.Name = "John"
	buyer.Surname = "Doe"
	buyer.GsmNumber = "+905350000000"
	buyer.Email = "email@email.com"
	buyer.IdentityNumber = "74300864791"
	buyer.LastLoginDate = "2015-10-05 12:43:35"
	buyer.RegistrationDate = "2013-04-21 15:12:09"
	buyer.RegistrationAddress = "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1"
	buyer.Ip = "85.34.78.112"
	buyer.City = "Istanbul"
	buyer.Country = "Turkey"
	buyer.ZipCode = "34732"
	req.Buyer = buyer

	shippingAddress := model.NewAddress()
	shippingAddress.ContactName = "Jane Doe"
	shippingAddress.City = "Istanbul"
	shippingAddress.Country = "Turkey"
	shippingAddress.Description = "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1"
	shippingAddress.ZipCode = "34742"
	req.ShippingAddress = shippingAddress

	billingAddress := model.NewAddress()
	billingAddress.ContactName = "Jane Doe"
	billingAddress.City = "Istanbul"
	billingAddress.Country = "Turkey"
	billingAddress.Description = "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1"
	billingAddress.ZipCode = "34742"
	req.BillingAddress = billingAddress

	firstBasketItem := model.NewBasketItem()
	firstBasketItem.Id = "BI101"
	firstBasketItem.Name = "Binocular"
	firstBasketItem.Category1 = "Collectibles"
	firstBasketItem.Category2 = "Accessories"
	firstBasketItem.ItemType = iyzipay.BasketItemTypePHYSICAL
	firstBasketItem.Price = "0.3"

	secondBasketItem := model.NewBasketItem()
	secondBasketItem.Id = "BI102"
	secondBasketItem.Name = "Game code"
	secondBasketItem.Category1 = "Game"
	secondBasketItem.Category2 = "Online Game Items"
	secondBasketItem.ItemType = iyzipay.BasketItemTypeVIRTUAL
	secondBasketItem.Price = "0.5"

	req.BasketItems = append(req.BasketItems, firstBasketItem, secondBasketItem)

	options := &iyzipay.Options{
		APIKey:    "your api key",
		SecretKey: "your secret key",
		BaseURL:   "https://sandbox-api.iyzipay.com",
	}

	threedsInitialize, _ := iyzipay.NewService().ThreedsInitialize(req, options)
  ```


# Mock test cards

Test cards that can be used to simulate a *successful* payment:

Card Number      | Bank                       | Card Type
-----------      | ----                       | ---------
5890040000000016 | Akbank                     | Master Card (Debit)  
5526080000000006 | Akbank                     | Master Card (Credit)  
4766620000000001 | Denizbank                  | Visa (Debit)  
4603450000000000 | Denizbank                  | Visa (Credit)
4729150000000005 | Denizbank Bonus            | Visa (Credit)  
4987490000000002 | Finansbank                 | Visa (Debit)  
5311570000000005 | Finansbank                 | Master Card (Credit)  
9792020000000001 | Finansbank                 | Troy (Debit)  
9792030000000000 | Finansbank                 | Troy (Credit)  
5170410000000004 | Garanti Bankası            | Master Card (Debit)  
5400360000000003 | Garanti Bankası            | Master Card (Credit)  
374427000000003  | Garanti Bankası            | American Express  
4475050000000003 | Halkbank                   | Visa (Debit)  
5528790000000008 | Halkbank                   | Master Card (Credit)  
4059030000000009 | HSBC Bank                  | Visa (Debit)  
5504720000000003 | HSBC Bank                  | Master Card (Credit)  
5892830000000000 | Türkiye İş Bankası         | Master Card (Debit)  
4543590000000006 | Türkiye İş Bankası         | Visa (Credit)  
4910050000000006 | Vakıfbank                  | Visa (Debit)  
4157920000000002 | Vakıfbank                  | Visa (Credit)  
5168880000000002 | Yapı ve Kredi Bankası      | Master Card (Debit)  
5451030000000000 | Yapı ve Kredi Bankası      | Master Card (Credit)  

*Cross border* test cards:

Card Number      | Country
-----------      | -------
4054180000000007 | Non-Turkish (Debit)
5400010000000004 | Non-Turkish (Credit)    

Test cards to get specific *error* codes:

Card Number       | Description
-----------       | -----------
5406670000000009  | Success but cannot be cancelled, refund or post auth
4111111111111129  | Not sufficient funds
4129111111111111  | Do not honour
4128111111111112  | Invalid transaction
4127111111111113  | Lost card
4126111111111114  | Stolen card
4125111111111115  | Expired card
4124111111111116  | Invalid cvc2
4123111111111117  | Not permitted to card holder
4122111111111118  | Not permitted to terminal
4121111111111119  | Fraud suspect
4120111111111110  | Pickup card
4130111111111118  | General error
4131111111111117  | Success but mdStatus is 0
4141111111111115  | Success but mdStatus is 4
4151111111111112  | 3dsecure initialize failed
