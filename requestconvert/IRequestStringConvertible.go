package requestconvert

type RequestStringConvertible interface {
	ToPKIRequestString() string
}
