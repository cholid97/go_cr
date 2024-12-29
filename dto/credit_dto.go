package dto

type UserContractResponse struct {
	FullName            string `json:"fullname,omitempty"`
	Asset               string `json:"asset,omitempty"`
	InstallmentAmount   int32  `json:"installment_amount,omitempty"`
	InstallmentInterest int32  `json:"installment_interest,omitempty"`
}
