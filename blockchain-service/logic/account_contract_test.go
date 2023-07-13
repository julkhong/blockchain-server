package logic

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContract_SendAndGetBalance(t *testing.T) {
	// Create a sample contract instance
	contract := NewContract()

	// Generate sample private and public keys
	privateKey, publicKey := generateKeyPair()

	// Create a sample account #1
	accountID := "account1"
	account := &Account{
		ID:        accountID,
		Balance:   big.NewInt(150),
		PublicKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(publicKey)),
	}
	contract.Accounts[accountID] = account

	// Create a sample account
	accountID = "account2"
	account = &Account{
		ID:        accountID,
		Balance:   big.NewInt(100),
		PublicKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(publicKey)),
	}
	contract.Accounts[accountID] = account

	// Simulate a send request
	fromID := "account1"
	toID := "account2"
	amount := 20.5
	request := createRequest(1, fromID, toID, amount, privateKey)
	response := contract.Send(fromID, toID, amount, request)

	// Check the send response
	if response.ReturnCode != "ok" {
		t.Errorf("Send failed: %s", response.Message)
	}

	// Check the account balances after the send
	expectedBalance := 129.5
	assert.Equal(t, expectedBalance, contract.Accounts[fromID].Balance, "Invalid receiver balance after send")

	expectedBalance = 120.5
	assert.Equal(t, expectedBalance, contract.Accounts[toID].Balance, "Invalid receiver balance after send")

	// Simulate a getBalance request
	request = createRequest(2, accountID, "", 0, privateKey)
	response = contract.GetBalance(accountID, request)

	// Check the getBalance response
	if response.ReturnCode != "ok" {
		t.Errorf("GetBalance failed: %s", response.Message)
	}

	// Check the account balance
	expectedBalance = 120.5
	assert.Equal(t, expectedBalance, contract.Accounts[toID].Balance, "Invalid receiver balance after send")

}

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey, &privateKey.PublicKey
}

func createRequest(id int, fromID, toID string, amount float64, privateKey *rsa.PrivateKey) Request {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	params := []float64{amount}
	key := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	signature := generateSignature(id, timestamp, params, key, privateKey)

	return Request{
		ID:        id,
		Timestamp: timestamp,
		Params:    params,
		Key:       key,
		Signature: signature,
	}
}

func generateSignature(id int, timestamp int64, params []float64, key string, privateKey *rsa.PrivateKey) string {
	str := fmt.Sprintf("%d%d", id, timestamp)
	for _, param := range params {
		str += fmt.Sprintf("%f", param)
	}
	str += key

	hash := sha256.Sum256([]byte(str))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	return base64.StdEncoding.EncodeToString(signature)
}
