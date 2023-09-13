package models

type TransferRequest struct {
	FromAccount string `json:"fromAccount,omitempty"`
	ToAccount   string `json:"toAccount,omitempty"`
	Amount      string `json:"amount,omitempty"`
}
