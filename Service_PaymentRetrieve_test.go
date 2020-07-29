package iyzipay

import (
	"testing"

	"github.com/stretchr/testify/assert"
	request "github.com/zeyit/iyzipay-go/request"
)

func TestPaymentRetrieve(t *testing.T) {
	req := request.NewRetrievePaymentRequest()
	req.PaymentId = "12268808"
	result, _ := NewService().PaymentRetrieve(req, options)

	PrintResponse(result)

	assert := assert.New(t)

	assert.Equal(StatusSUCCESS, result.Status)

}
