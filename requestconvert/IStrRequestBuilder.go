package requestconvert

type IStrRequestBuilder interface {
	Append(key string, value interface{}) IStrRequestBuilder
	AppendPrice(key string, value interface{}) IStrRequestBuilder
	AppendSuper(baseRequestString string) IStrRequestBuilder
	AppendList(key string, value interface{}) IStrRequestBuilder
	GetRequestString() string
}
