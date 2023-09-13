package services

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/omarattia3143/paytabs-backend-challenge/src/database"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
)

type AccountService interface {
	GetAllAccounts() *[]models.Account
	GetAccount(id string) *models.Account
	UpdateAccountsBalance(fromAccount *models.Account, toAccount *models.Account) error
}

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

func UpdateAccountsBalance(fromAccount *models.Account, toAccount *models.Account) error {
	// commit transaction once to avoid race cases
	txn := database.DB.Txn(true)

	var err error
	err = updateAccount(fromAccount, txn)
	if err != nil {
		return err
	}

	err = updateAccount(toAccount, txn)
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func updateAccount(accountToUpdate *models.Account, txn *memdb.Txn) error {
	raw, err := txn.First("account", "id", accountToUpdate.Id)
	if err != nil {
		return err
	}

	if raw != nil {
		dbAccount := raw.(*models.Account)
		// balance mapping
		dbAccount.Balance = accountToUpdate.Balance

		if err := txn.Insert("account", dbAccount); err != nil {
			return err
		}
	}
	return nil
}
