package iyzipay

import (
	"fmt"

	model "github.com/zeyit/iyzipay-go/model"
	request "github.com/zeyit/iyzipay-go/request"
)

type Service struct {
	client IClient
}

func NewService() *Service {
	return &Service{client: &restClient{}}
}

func NewServiceWithClient(client IClient) *Service {
	return &Service{client: client}
}

/*******************************************************************/

func (s *Service) APITest(options *Options) (*model.IyzipayResource, error) {
	m := model.NewIyzipayResource()
	headers := make(map[string]string)
	err := s.client.Init().Get(options.BaseURL+"/payment/test", headers, m)
	return m, err
}

func (s *Service) AmpCreate(request *request.ApmInitializeRequest, options *Options) (*model.Amp, error) {
	m := model.NewAmp()
	err := s.client.Init().Post(options.BaseURL+"/payment/apm/initialize", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) AmpRetrieve(request *request.RetrieveApmRequest, options *Options) (*model.Amp, error) {
	m := model.NewAmp()
	err := s.client.Init().Post(options.BaseURL+"/payment/apm/retrieve", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) Approval(request *request.ApprovalRequest, options *Options) (*model.Approval, error) {
	m := model.NewApproval()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/item/approve", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BasicBkm(request *request.RetrieveBkmRequest, options *Options) (*model.BasicBkm, error) {
	m := model.NewBasicBkm()
	err := s.client.Init().Post(options.BaseURL+"/payment/bkm/auth/detail/basic", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BasicBkmInitialize(request *request.BasicBkmInitializeRequest, options *Options) (*model.BasicBkmInitialize, error) {
	m := model.NewBasicBkmInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/bkm/initialize/basic", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) BasicPayment(request *request.BasicPaymentRequest, options *Options) (*model.BasicPayment, error) {
	m := model.NewBasicPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/auth/basic", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BasicPaymentPostAuth(request *request.PaymentPostAuthRequest, options *Options) (*model.BasicPaymentPostAuth, error) {
	m := model.NewBasicPaymentPostAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/postauth/basic", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BasicPaymentPreAuth(request *request.BasicPaymentRequest, options *Options) (*model.BasicPaymentPreAuth, error) {
	m := model.NewBasicPaymentPreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/preauth/basic", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BasicThreedsInitialize(request *request.BasicPaymentRequest, options *Options) (*model.BasicThreedsInitialize, error) {
	m := model.NewBasicThreedsInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/initialize/basic", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) BasicThreedsInitializePreAuth(request *request.BasicPaymentRequest, options *Options) (*model.BasicThreedsInitializePreAuth, error) {
	m := model.NewBasicThreedsInitializePreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/initialize/preauth/basic", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) BasicThreedsPayment(request *request.BasicPaymentRequest, options *Options) (*model.BasicThreedsPayment, error) {
	m := model.NewBasicThreedsPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/auth/basic", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BinNumber(request *request.RetrieveBinNumberRequest, options *Options) (*model.BinNumber, error) {
	m := model.NewBinNumber()
	err := s.client.Init().Post(options.BaseURL+"/payment/bin/check", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) Bkm(request *request.RetrieveBkmRequest, options *Options) (*model.Bkm, error) {
	m := model.NewBkm()
	err := s.client.Init().Post(options.BaseURL+"/payment/bkm/auth/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) BkmInitialize(request *request.BkmInitializeRequest, options *Options) (*model.BkmInitialize, error) {
	m := model.NewBkmInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/bkm/initialize", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) BouncedBankTransferList(request *request.RetrieveTransactionsRequest, options *Options) (*model.BouncedBankTransferList, error) {
	m := model.NewBouncedBankTransferList()
	err := s.client.Init().Post(options.BaseURL+"/reporting/settlement/bounced", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) Cancel(request *request.CancelRequest, options *Options) (*model.Cancel, error) {
	m := model.NewCancel()
	err := s.client.Init().Post(options.BaseURL+"/payment/cancel", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CardCreate(request *request.CreateCardRequest, options *Options) (*model.Card, error) {
	m := model.NewCard()
	err := s.client.Init().Post(options.BaseURL+"/cardstorage/card", getHeaders(request, options), request, m)
	return m, err
}
func (s *Service) CardDelete(request *request.DeleteCardRequest, options *Options) (*model.Card, error) {
	m := model.NewCard()
	err := s.client.Init().Delete(options.BaseURL+"/cardstorage/card", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CardList(request *request.CardListRequest, options *Options) (*model.CardList, error) {
	m := model.NewCardList()
	err := s.client.Init().Post(options.BaseURL+"/cardstorage/cards", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CardManagementPageCard(request *request.RetrieveCardManagementPageCardRequest, options *Options) (*model.CardManagementPageCard, error) {
	url := fmt.Sprintf("%s/v1/card-management/pages/%s/cards?locale=%s&conversationId=%s",
		options.BaseURL, request.PageToken, request.Locale, request.ConversationId)

	m := model.NewCardManagementPageCard()
	err := s.client.Init().Post(url, getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CardManagementPageInitialize(request *request.CardManagementPageInitializeRequest, options *Options) (*model.CardManagementPageInitialize, error) {
	m := model.NewCardManagementPageInitialize()
	err := s.client.Init().Post(options.BaseURL+"/v1/card-management/pages", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CheckoutForm(request *request.RetrieveCheckoutFormRequest, options *Options) (*model.CheckoutForm, error) {
	m := model.NewCheckoutForm()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/checkoutform/auth/ecom/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CheckoutFormInitialize(request *request.CheckoutFormInitializeRequest, options *Options) (*model.CheckoutFormInitialize, error) {
	m := model.NewCheckoutFormInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/checkoutform/initialize/auth/ecom",
		getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CheckoutFormInitializePreAuth(request *request.CheckoutFormInitializeRequest, options *Options) (*model.CheckoutFormInitializePreAuth, error) {
	m := model.NewCheckoutFormInitializePreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/checkoutform/initialize/preauth/ecom",
		getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CrossBookingFromSubMerchant(request *request.CrossBookingRequest, options *Options) (*model.CrossBookingFromSubMerchant, error) {
	m := model.NewCrossBookingFromSubMerchant()
	err := s.client.Init().Post(options.BaseURL+"/crossbooking/receive", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) CrossBookingToSubMerchant(request *request.CrossBookingRequest, options *Options) (*model.CrossBookingToSubMerchant, error) {
	m := model.NewCrossBookingToSubMerchant()
	err := s.client.Init().Post(options.BaseURL+"/crossbooking/send", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) Disapproval(request *request.ApprovalRequest, options *Options) (*model.Disapproval, error) {
	m := model.NewDisapproval()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/item/disapprove", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) InstallmentInfo(request *request.RetrieveInstallmentInfoRequest, options *Options) (*model.InstallmentInfo, error) {
	m := model.NewInstallmentInfo()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/installment", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentCreate(request *request.CreatePaymentRequest, options *Options) (*model.Payment, error) {
	m := model.NewPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/auth", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentRetrieve(request *request.RetrievePaymentRequest, options *Options) (*model.Payment, error) {
	m := model.NewPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentItemUpdate(request *request.PaymentItemRequest, options *Options) (*model.PaymentItem, error) {
	m := model.NewPaymentItem()
	err := s.client.Init().Put(options.BaseURL+"/payment/item", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentPostAuth(request *request.PaymentPostAuthRequest, options *Options) (*model.PaymentPostAuth, error) {
	m := model.NewPaymentPostAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/postauth", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentPreAuthCreate(request *request.CreatePaymentRequest, options *Options) (*model.PaymentPreAuth, error) {
	m := model.NewPaymentPreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/preauth", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PaymentPreAuthRetrieve(request *request.RetrievePaymentRequest, options *Options) (*model.PaymentPreAuth, error) {
	m := model.NewPaymentPreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PayoutCompletedTransactionList(request *request.RetrieveTransactionsRequest, options *Options) (*model.PayoutCompletedTransactionList, error) {
	m := model.NewPayoutCompletedTransactionList()
	err := s.client.Init().Post(options.BaseURL+"/reporting/settlement/payoutcompleted", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PayWithIyzico(request *request.RetrievePayWithIyzicoRequest, options *Options) (*model.PayWithIyzico, error) {
	m := model.NewPayWithIyzico()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/checkoutform/auth/ecom/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) PayWithIyzicoInitialize(request *request.CreatePayWithIyzicoInitializeRequest, options *Options) (*model.PayWithIyzicoInitialize, error) {
	m := model.NewPayWithIyzicoInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/pay-with-iyzico/initialize", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) Refund(request *request.RefundRequest, options *Options) (*model.Refund, error) {
	m := model.NewRefund()
	err := s.client.Init().Post(options.BaseURL+"/payment/refund", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) RefundChargedFromMerchant(request *request.RefundRequest, options *Options) (*model.RefundChargedFromMerchant, error) {
	m := model.NewRefundChargedFromMerchant()
	err := s.client.Init().Post(options.BaseURL+"/payment/iyzipos/refund/merchant/charge", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) SubMerchantCreate(request *request.CreateSubMerchantRequest, options *Options) (*model.SubMerchant, error) {
	m := model.NewSubMerchant()
	err := s.client.Init().Post(options.BaseURL+"/onboarding/submerchant", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) SubMerchantUpdate(request *request.UpdateSubMerchantRequest, options *Options) (*model.SubMerchant, error) {
	m := model.NewSubMerchant()
	err := s.client.Init().Post(options.BaseURL+"/onboarding/submerchant", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) SubMerchantRetrieve(request *request.RetrieveSubMerchantRequest, options *Options) (*model.SubMerchant, error) {
	m := model.NewSubMerchant()
	err := s.client.Init().Post(options.BaseURL+"/onboarding/submerchant/detail", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) ThreedsInitialize(request *request.CreatePaymentRequest, options *Options) (*model.ThreedsInitialize, error) {
	m := model.NewThreedsInitialize()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/initialize", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) ThreedsInitializePreAuth(request *request.CreatePaymentRequest, options *Options) (*model.ThreedsInitializePreAuth, error) {
	m := model.NewThreedsInitializePreAuth()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/initialize/preauth", getHeaders(request, options), request, m)
	if m.HtmlContent != "" {
		m.HtmlContent, err = decodeString(m.HtmlContent)
	}
	return m, err
}

func (s *Service) ThreedsPaymentCreate(request *request.ThreedsPaymentRequest, options *Options) (*model.ThreedsPayment, error) {
	m := model.NewThreedsPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/3dsecure/auth", getHeaders(request, options), request, m)
	return m, err
}

func (s *Service) ThreedsPaymentRetrieve(request *request.PaymentRequest, options *Options) (*model.ThreedsPayment, error) {
	m := model.NewThreedsPayment()
	err := s.client.Init().Post(options.BaseURL+"/payment/detail", getHeaders(request, options), request, m)
	return m, err
}
