package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AmountTransferContract is a smart contract to manage amount transfer between account
type AmountTransferContract struct {
	contractapi.Contract
}

// Account represents an account
type Account struct {
	ID        string `json:"id"`
	Balance   uint64 `json:"balance"`
	PublicKey string `json:"publicKey"`
}

// Init initializes an account
func (a *AmountTransferContract) Init(ctx TransactionContextInterface, publicKey string, amount uint64) (string, error) {
	uniqueID := uuid.New()
	accountID := uniqueID.String()
	newAccount := Account{
		ID:        accountID,
		Balance:   amount,
		PublicKey: publicKey,
	}
	accountJSON, err := json.Marshal(newAccount)
	if err != nil {
		return "", fmt.Errorf("failed to serialize new account: %w", err)
	}
	err = ctx.GetStub().PutState(accountID, accountJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put new account state: %w", err)
	}
	return accountID, nil
}

// Send transfers amount from sender to a receiver
func (a *AmountTransferContract) Send(ctx TransactionContextInterface, senderID, receiverID string, amount uint64) error {
	senderJSON, err := ctx.GetStub().GetState(senderID)
	if err != nil {
		return fmt.Errorf("failed to read sender account from the world state: %w", err)
	}
	if senderJSON == nil {
		return fmt.Errorf("account with ID %s does not exist", senderID)
	}

	receiverJSON, err := ctx.GetStub().GetState(receiverID)
	if err != nil {
		return fmt.Errorf("failed to read receiver account from the world state: %w", err)
	}
	if receiverJSON == nil {
		return fmt.Errorf("account with ID %s does not exist", receiverID)
	}

	account := new(Account)
	err = json.Unmarshal(senderJSON, account)
	if err != nil {
		return fmt.Errorf("failed to deserialize sender account JSON: %w", err)
	}
	balance := int(account.Balance) - int(amount)
	if balance < 0 {
		return fmt.Errorf("failed to transfer as sender has insufficient balance")
	}
	account.Balance = uint64(balance)
	senderJSON, err = json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to serialize sender account: %w", err)
	}
	err = ctx.GetStub().PutState(senderID, senderJSON)
	if err != nil {
		return fmt.Errorf("failed to update sender account in the world state: %w", err)
	}

	account = new(Account)
	err = json.Unmarshal(receiverJSON, account)
	if err != nil {
		return fmt.Errorf("failed to deserialize receiver account JSON: %w", err)
	}
	balance = int(account.Balance) + int(amount)
	account.Balance = uint64(balance)
	receiverJSON, err = json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to serialize receiver account: %w", err)
	}
	err = ctx.GetStub().PutState(receiverID, receiverJSON)
	if err != nil {
		return fmt.Errorf("failed to update receiver account in the world state: %w", err)
	}
	return nil
}

func (a *AmountTransferContract) GetBalance(ctx TransactionContextInterface, accountID string) (uint64, error) {
	accountJSON, err := ctx.GetStub().GetState(accountID)
	if err != nil {
		return 0, fmt.Errorf("failed to read account from the world state: %w", err)
	}
	if accountJSON == nil {
		return 0, fmt.Errorf("asset with ID %s does not exist", accountID)
	}

	account := new(Account)
	err = json.Unmarshal(accountJSON, account)
	if err != nil {
		return 0, fmt.Errorf("failed to deserialize account JSON: %w", err)
	}

	return account.Balance, nil
}
