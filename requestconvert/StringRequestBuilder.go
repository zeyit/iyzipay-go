package requestconvert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type StrRequestBuilder struct {
	requestItems []string
}

func NewStrRequestBuilder() IStrRequestBuilder {
	return &StrRequestBuilder{}
}

func (b *StrRequestBuilder) Append(key string, value interface{}) IStrRequestBuilder {
	if value != nil {
		switch value.(type) {
		case RequestStringConvertible:
			b.appendKeyValue(key, (value.(RequestStringConvertible)).ToPKIRequestString())
		case *int:
			if v, ok := value.(*int); ok && v != nil {
				b.appendKeyValue(key, strconv.FormatInt(int64(*v), 10))
			}
		case int:
			b.appendKeyValue(key, strconv.FormatInt(int64(value.(int)), 10))

		case int64:
			b.appendKeyValue(key, strconv.FormatInt(value.(int64), 10))
		case string:
			b.appendKeyValue(key, value.(string))
		case bool:
			b.appendKeyValue(key, strconv.FormatBool(value.(bool)))

		default:
			fmt.Println("Value ", value)
			panic("Tip desteklenmiyor")

		}
	}
	return b
}

func (b *StrRequestBuilder) AppendPrice(key string, value interface{}) IStrRequestBuilder {
	if value != nil {
		if v, ok := value.(string); ok && v != "" {
			p, _ := strconv.ParseFloat(value.(string), 32)
			price := fmt.Sprintf("%.2f", p)
			if strings.HasSuffix(price, ".00") {
				price = strings.Replace(price, ".00", ".0", 1)
			}
			b.appendKeyValue(key, price)
		}
	}
	return b
}

func (b *StrRequestBuilder) AppendSuper(baseRequestString string) IStrRequestBuilder {

	if len(baseRequestString) > 2 {
		b.requestItems = append(b.requestItems, baseRequestString[1:len(baseRequestString)-1])
	}
	return b
}

func (b *StrRequestBuilder) AppendList(key string, value interface{}) IStrRequestBuilder {
	if value != nil {
		s := reflect.ValueOf(value)
		if s.Kind() != reflect.Slice {

			fmt.Println("InterfaceSlice() given a non-slice type")
			return b
		}

		ret := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}
		appendedValue := []string{}
		for _, v := range ret {

			if val, ok := v.(RequestStringConvertible); ok {
				appendedValue = append(appendedValue, val.ToPKIRequestString())
			}
			if val, ok := v.(int); ok {
				appendedValue = append(appendedValue, strconv.FormatInt(int64(val), 10))
			}
		}
		b.appendKeyValue(key, "["+strings.Join(appendedValue, ",")+"]")
	}
	return b
}

func (b *StrRequestBuilder) appendKeyValue(key, value string) {
	if value != "" {
		b.requestItems = append(b.requestItems, key+"="+value)
	}
}

func (b *StrRequestBuilder) GetRequestString() string {
	return "[" + strings.Join(b.requestItems, ",") + "]"
}
