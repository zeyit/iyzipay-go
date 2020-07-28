package iyzipay

import (
	"testing"
)

func TestAPI(t *testing.T) {
	result, err := NewService().APITest(options)
	PrintResponse(result)
	if result.Status != StatusSUCCESS {
		t.Errorf("Status: %s", result.Status)
	}
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
