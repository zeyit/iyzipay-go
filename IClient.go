package iyzipay

type IClient interface {
	Get(url string, headers map[string]string, result interface{}) error
	Post(url string, headers map[string]string, request interface{}, result interface{}) error
	Put(url string, headers map[string]string, request interface{}, result interface{}) error
	Delete(url string, headers map[string]string, request interface{}, result interface{}) error
	Init() IClient
}
