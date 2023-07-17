package handlers

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func InitAccounts(contract *client.Contract) {
	fmt.Printf("\n--> Submit Transaction: InitAccounts, function creates the initial set of accounts on the ledger \n")

	_, err := contract.SubmitTransaction("InitAccounts")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}
