package test

import (
	"github.com/omarattia3143/paytabs-backend-challenge/src/database"
	"github.com/omarattia3143/paytabs-backend-challenge/src/models"
	"github.com/omarattia3143/paytabs-backend-challenge/src/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	database.SetupAndConnectMemDb()
	database.SeedDb()
}

func TestStartTransferProcess(t *testing.T) {

	fromAccount := models.Account{Id: "fcc0de1d-0a37-407f-a6ac-854fd885c69b", Name: "Dynabox", Balance: "2260.45"}
	toAccount := models.Account{Id: "23477b82-84a1-41fe-b259-c4117918245a", Name: "Rooxo", Balance: "4254.33"}
	transferRequest := models.TransferRequest{FromAccount: fromAccount.Id, ToAccount: toAccount.Id, Amount: "50"}

	// call the code we are testing
	err := services.StartTransferProcess(transferRequest)

	// Check if fromAccount and toAccount are not nil
	if &fromAccount == nil || &toAccount == nil {
		t.Fatal("Accounts should not be nil")
	}

	// assert that the expectations were met
	assert.Nil(t, err)
}
