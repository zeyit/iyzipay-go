package model

type InstallmentInfo struct {
	*IyzipayResource
	InstallmentDetails []*InstallmentDetail `json:"installmentDetails,omitempty"`
}

func NewInstallmentInfo() *InstallmentInfo {
	return &InstallmentInfo{IyzipayResource: NewIyzipayResource()}
}
