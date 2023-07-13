package logic

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
)

type Account struct {
	ID        string   `json:"id"`
	Balance   *big.Int `json:"balance"`
	PublicKey string   `json:"publicKey"`
}

type Request struct {
	ID        int       `json:"id"`
	Timestamp int64     `json:"timestamp"`
	Params    []float64 `json:"params"`
	Key       string    `json:"key"`
	Signature string    `json:"signature"`
}

type Response struct {
	ReturnCode string   `json:"returnCode"`
	ID         int      `json:"id"`
	Message    string   `json:"message,omitempty"`
	Amount     *big.Int `json:"amount,omitempty"`
}

type AccountContract struct {
	Accounts      map[string]*Account
	LastRequestID int
	PrivateKey    *rsa.PrivateKey
}

func NewContract() *AccountContract {
	return &AccountContract{
		Accounts:      make(map[string]*Account),
		LastRequestID: 0,
		PrivateKey:    generatePrivateKey(),
	}
}

func generatePrivateKey() *rsa.PrivateKey {
	// Generate and return a new RSA private key
	// Make sure to securely store the private key in a production environment
	return nil
}

func (c *AccountContract) Send(fromID, toID string, amount float64, request Request) Response {
	// Check if the accounts exist
	fromAccount, fromExists := c.Accounts[fromID]
	toAccount, toExists := c.Accounts[toID]
	if !fromExists || !toExists {
		return Response{
			ReturnCode: "error",
			ID:         request.ID,
			Message:    "Invalid account IDs",
		}
	}

	// Validate the request signature
	validSignature := validateSignature(request, c.PrivateKey.Public().(*rsa.PublicKey))
	if !validSignature {
		return Response{
			ReturnCode: "error",
			ID:         request.ID,
			Message:    "Invalid request signature",
		}
	}

	// Check if the sender has sufficient balance
	if fromAccount.Balance.Cmp(big.NewInt(int64(amount))) < 0 {
		return Response{
			ReturnCode: "error",
			ID:         request.ID,
			Message:    "Insufficient balance",
		}
	}

	// Update the balances
	fromAccount.Balance.Sub(fromAccount.Balance, big.NewInt(int64(amount)))
	toAccount.Balance.Add(toAccount.Balance, big.NewInt(int64(amount)))

	return Response{
		ReturnCode: "ok",
		ID:         request.ID,
	}
}

func (c *AccountContract) GetBalance(accountID string, request Request) Response {
	// Check if the account exists
	account, exists := c.Accounts[accountID]
	if !exists {
		return Response{
			ReturnCode: "error",
			ID:         request.ID,
			Message:    "Invalid account ID",
		}
	}

	// Validate the request signature
	validSignature := validateSignature(request, c.PrivateKey.Public().(*rsa.PublicKey))
	if !validSignature {
		return Response{
			ReturnCode: "error",
			ID:         request.ID,
			Message:    "Invalid request signature",
		}
	}

	return Response{
		ReturnCode: "ok",
		ID:         request.ID,
		Amount:     account.Balance,
	}
}

func validateSignature(request Request, publicKey *rsa.PublicKey) bool {
	// Build the string representation of the request fields
	str := fmt.Sprintf("%d%d", request.ID, request.Timestamp)
	for i := 0; i < len(request.Params); i++ {
		str += fmt.Sprintf("%f", request.Params[i])
	}
	str += request.Key

	// Decode the request signature
	signature, err := base64.StdEncoding.DecodeString(request.Signature)
	if err != nil {
		return false
	}

	// Verify the request signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, []byte(str), signature)
	if err != nil {
		return false
	}

	return true
}
