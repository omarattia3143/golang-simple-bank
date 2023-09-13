package models

import "github.com/google/uuid"

type MyUUID uuid.UUID

func (u *MyUUID) UnmarshalJSON(data []byte) error {
	// remove quotes
	data = data[1 : len(data)-1]

	id, err := uuid.Parse(string(data))
	if err != nil {
		return err
	}

	*u = MyUUID(id)
	return nil
}
