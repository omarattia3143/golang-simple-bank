package services

import (
	"fmt"
	"github.com/omarattia3143/paytabs-backend-challenge/src/database"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
)

func GetAllAccounts() *[]models.Account {

	// Create read-only transaction
	txn := database.DB.Txn(false)
	defer txn.Abort()

	// List all the accounts
	it, err := txn.Get("account", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("get all accounts from db")
	var accounts []models.Account
	for obj := it.Next(); obj != nil; obj = it.Next() {
		account := obj.(*models.Account)
		accounts = append(accounts, *account)
	}
	return &accounts
}

func GetAccount(id string) *models.Account {

	log := fmt.Sprintf("getting account id %s", id)
	fmt.Println(log)

	// Create read-only transaction
	txn := database.DB.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("account", "id", id)
	if err != nil {
		panic(err)
	}
	if raw == nil {
		return nil
	}
	return raw.(*models.Account)
}
