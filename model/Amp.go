package model

type Amp struct {
	*ApmResource
}

func NewAmp() *Amp {
	return &Amp{ApmResource: NewApmResource()}
}
