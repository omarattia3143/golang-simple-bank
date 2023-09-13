package models

type Account struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Balance string `json:"balance,omitempty"`
}
