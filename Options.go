package iyzipay

type Options struct {
	APIKey    string
	SecretKey string
	BaseURL   string
}

func NewOptions() *Options {
	return &Options{}
}
