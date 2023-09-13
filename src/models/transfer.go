package models

type Transfer struct {
	FromAccount MyUUID `json:"fromAccount,omitempty"`
	ToAccount   MyUUID `json:"toAccount,omitempty"`
	Amount      string `json:"amount,omitempty"`
}
