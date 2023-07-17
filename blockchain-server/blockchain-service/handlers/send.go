package handlers

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Send ...
func Send(contract *client.Contract, senderID, receiverID string, amount string) {
	fmt.Printf("\n--> Async Submit Transaction: Send, sends amount")

	_, commit, err := contract.SubmitAsync("Send", client.WithArguments(senderID, receiverID, amount))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}

	fmt.Printf("\n*** Successfully submitted transaction to transfer amount")
	fmt.Println("*** Waiting for transaction commit.")

	if commitStatus, err := commit.Status(); err != nil {
		panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !commitStatus.Successful {
		panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}
