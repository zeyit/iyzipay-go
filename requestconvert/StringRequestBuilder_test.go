package requestconvert

import (
	"testing"
)

func TestEmptyBuilder(t *testing.T) {
	builder := NewStrRequestBuilder()
	if "[]" != builder.GetRequestString() {
		t.Error("Builder err")
	}
}

func TestBuilderConvert(t *testing.T) {
	builder := NewStrRequestBuilder()
	builder.Append("conversationId", "123456")
	builder.Append("locale", "tr")
	builder.AppendPrice("price", "1.0")

	str := "[conversationId=123456,locale=tr,price=1.0]"
	if str != builder.GetRequestString() {
		t.Error("Convert err")
	}
}
