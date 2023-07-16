package logic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testPublicKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt/LtqNLiYkmf4jjO3XXG\nD6fI5gg30yMLc+uD1MKuKzrCIBg7K2K9DWWwozqI4hfU8bV8wJt8fXvF7wF4U+eO\nsUzYyxNqAlYZW4r9DQ77RyQtskU+x+MT7E8aPjVswm1Yv9Fe7UsjDqTm2ugjDX3Z\nRds7Qbgs5TwHkBNlrrDGDLm/BH2N3BpC9Ry0ToI8Z9T7nhiAA3Fz1HqQuIi6A1IK\n1G0S5KT4a/yFlKklkBYol/y+DZx2fWx13dWkKofu30AVAGUuewU9HEt9zLr/U55w\n9WwRr9htB79j0TmqCkncSnwqHQxvR/7zL7COV3v2G5rJ2anb7wYh0OHNJ1jMaXz7\nEwIDAQAB\n-----END PUBLIC KEY-----\n"

func BenchmarkSend(b *testing.B) {
	// Create a new instance of the smart contract
	contract := &AmountTransferContract{}
	ctx := &MockTransactionContext{
		Stub:          &MockStub{StateMap: make(map[string][]byte)},
		TransactionID: "test-transaction",
		ChannelID:     "test-channel",
		Contract:      contract,
	}

	// Initialize the contract
	accountA, err := contract.Init(ctx, testPublicKey, 10000000000000)
	require.NoError(b, err)
	accountB, err := contract.Init(ctx, testPublicKey, 10000000000000)
	require.NoError(b, err)

	// Perform the benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = contract.Send(ctx, accountA, accountB, 1)
		require.NoError(b, err)
	}
}

func BenchmarkGetBalance(b *testing.B) {
	// Create a new instance of the smart contract
	contract := &AmountTransferContract{}
	ctx := &MockTransactionContext{
		Stub:          &MockStub{StateMap: make(map[string][]byte)},
		TransactionID: "test-transaction",
		ChannelID:     "test-channel",
		Contract:      contract,
	}
	// Initialize the contract
	accountA, err := contract.Init(ctx, testPublicKey, 100)
	require.NoError(b, err)
	expectedRes := &Account{
		ID:        accountA,
		Balance:   100,
		PublicKey: testPublicKey,
	}
	// Perform the benchmarking
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, getBalanceErr := contract.GetBalance(ctx, accountA)
		require.NoError(b, getBalanceErr)
		require.Equal(b, res, expectedRes.Balance)
	}
}
