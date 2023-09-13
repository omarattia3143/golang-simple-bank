package test

import (
	"github.com/omarattia3143/paytabs-backend-challenge/src/database"
	"github.com/omarattia3143/paytabs-backend-challenge/src/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	database.SetupAndConnectMemDb()
	database.SeedDb()
}
func TestGetAllAccounts_ShouldReturnFullList(t *testing.T) {
	accounts := services.GetAllAccounts()
	assert.NotNil(t, accounts)
}

func TestGetAccount_ShouldReturnSingleAccount(t *testing.T) {
	account := services.GetAccount("fcc0de1d-0a37-407f-a6ac-854fd885c69b")
	assert.NotNil(t, account)
}
