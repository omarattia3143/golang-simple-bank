package services

import (
	"errors"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
	"github.com/shopspring/decimal"
)

func StartTransferProcess(transferRequest models.TransferRequest) error {
	// get both accounts
	fromAccount := GetAccount(transferRequest.FromAccount)
	toAccount := GetAccount(transferRequest.ToAccount)

	// validate that fromAccount can transfer the amount
	balance, _ := decimal.NewFromString(fromAccount.Balance)
	amount, _ := decimal.NewFromString(transferRequest.Amount)
	toAccountBalance, _ := decimal.NewFromString(toAccount.Balance)

	result := balance.Sub(amount)
	if result.LessThan(decimal.NewFromInt(0)) {
		return errors.New("transfer rejected, balance is not enough for the transaction")
	}

	// do the transactions on from and to accounts
	//from
	fromAccount.Balance = result.String()

	//to
	toAccount.Balance = toAccountBalance.Add(amount).String()

	// update accounts in database
	err := UpdateAccountsBalance(fromAccount, toAccount)
	if err != nil {
		return err
	}

	return nil
}
