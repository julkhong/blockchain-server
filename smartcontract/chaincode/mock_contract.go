package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AmountTransferContractM struct {
	contractapi.Contract
}

// InitM initializes an account
func (a *AmountTransferContract) InitM(ctx TransactionContextInterface, publicKey string, amount uint64) (string, error) {
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

func (a *AmountTransferContract) InitAccountsM(ctx TransactionContextInterface) error {
	accounts := []Account{
		{ID: "acc1", Balance: 100, PublicKey: "key1"},
		{ID: "acc2", Balance: 100, PublicKey: "key2"},
		{ID: "acc3", Balance: 100, PublicKey: "key3"},
		{ID: "acc4", Balance: 100, PublicKey: "key4"},
		{ID: "acc5", Balance: 100, PublicKey: "key5"},
		{ID: "acc6", Balance: 100, PublicKey: "key6"},
	}

	for _, asset := range accounts {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return fmt.Errorf("failed to serialize new account: %w", err)
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put new account state: %w", err)
		}
	}

	return nil
}

// SendM transfers amount from sender to a receiver
func (a *AmountTransferContract) SendM(ctx TransactionContextInterface, senderID, receiverID string, amount uint64) error {
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

func (a *AmountTransferContract) GetBalanceM(ctx TransactionContextInterface, accountID string) (uint64, error) {
	accountJSON, err := ctx.GetStub().GetState(accountID)
	if err != nil {
		return 0, fmt.Errorf("failed to read account from the world state: %w", err)
	}
	if accountJSON == nil {
		return 0, fmt.Errorf("account with ID %s does not exist", accountID)
	}

	account := new(Account)
	err = json.Unmarshal(accountJSON, account)
	if err != nil {
		return 0, fmt.Errorf("failed to deserialize account JSON: %w", err)
	}

	return account.Balance, nil
}
