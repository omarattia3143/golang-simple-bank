package database

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
	"os"
)

var DB *memdb.MemDB

// SetupAndConnectMemDb SetupMemDb setup memory database, I am using go-memDb
func SetupAndConnectMemDb() {
	// add database schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"account": {
				Name: "account",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "Id"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"balance": {
						Name:   "balance",
						Unique: false,
						// as go-memDb doesn't support decimal types, so I will be using string instead
						Indexer: &memdb.StringFieldIndex{Field: "Balance"},
					},
				},
			},
		},
	}

	// connect to db
	var err error
	DB, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

func SeedDb() {
	// read accounts from the provided json file in the challenge PDF
	jsonFile, err := os.ReadFile("mocks/accounts-mock.json")
	if err != nil {
		panic(err)
	}

	// deserialize the json and map it to account model
	var accounts []*models.Account
	err = json.Unmarshal(jsonFile, &accounts)
	if err != nil {
		panic(err)
	}

	// make a write transaction to insert accounts into db
	txn := DB.Txn(true)

	// loop over accounts and add them to db
	for _, s := range accounts {
		if err := txn.Insert("account", s); err != nil {
			panic(err)
		}
	}

	// commit the transaction
	txn.Commit()
	fmt.Println("Accounts are added to the memory database")
}
