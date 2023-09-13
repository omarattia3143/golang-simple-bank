package models

import "github.com/google/uuid"

type MyUUID uuid.UUID

type Account struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Balance string `json:"balance,omitempty"`
}

func (u *MyUUID) UnmarshalJSON(data []byte) error {
	// Remove the quotes around the UUID string
	data = data[1 : len(data)-1]

	// Parse the UUID
	id, err := uuid.Parse(string(data))
	if err != nil {
		return err
	}

	*u = MyUUID(id)
	return nil
}

//type MyDecimal decimal.Decimal
//type Account struct {
//	Id      MyUUID    `json:"id,omitempty"`
//	Name    string    `json:"name,omitempty"`
//	Balance MyDecimal `json:"balance,omitempty"`
//}
//
//
//func (d *MyDecimal) UnmarshalJSON(data []byte) error {
//	// Remove the quotes around the decimal string
//	data = data[1 : len(data)-1]
//	value, err := decimal.NewFromString(string(data))
//	if err != nil {
//		return err
//	}
//
//	*d = MyDecimal(value)
//	return nil
//}
