package handlers

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// GetBalance ...
func GetBalance(contract *client.Contract, accountID string) string {
	fmt.Printf("\n--> Evaluate Transaction: GetBalance, function returns balance\n")

	evaluateResult, err := contract.EvaluateTransaction("GetBalance", accountID)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
	return result
}
