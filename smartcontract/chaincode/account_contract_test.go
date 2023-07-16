package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestAmountTransferContract_Send(t *testing.T) {
	// Create a new instance of the smart contract
	amountTransferContract := new(AmountTransferContract)

	// Create a new mock transaction context
	ctx := &MockTransactionContext{
		Stub:          &MockStub{StateMap: make(map[string][]byte)},
		TransactionID: "test-transaction",
		ChannelID:     "test-channel",
		Contract:      amountTransferContract,
	}
	senderAccount := &Account{
		ID:        "sender-id",
		Balance:   100,
		PublicKey: "",
	}
	receiverAccount := &Account{
		ID:        "receiver-id",
		Balance:   100,
		PublicKey: "",
	}

	tests := []struct {
		name                string
		senderID            string
		receiverID          string
		amount              uint64
		sender              *Account
		receiver            *Account
		ctxErr              bool
		expectedErr         error
		expectedSenderAmt   uint64
		expectedReceiverAmt uint64
	}{
		{
			name:                "happy path",
			senderID:            "sender-id",
			receiverID:          "receiver-id",
			amount:              20,
			sender:              senderAccount,
			receiver:            receiverAccount,
			ctxErr:              false,
			expectedErr:         nil,
			expectedSenderAmt:   80,
			expectedReceiverAmt: 120,
		},
		{
			name:                "happy path - zero amount",
			senderID:            "sender-id",
			receiverID:          "receiver-id",
			amount:              0,
			sender:              senderAccount,
			receiver:            receiverAccount,
			ctxErr:              false,
			expectedErr:         nil,
			expectedSenderAmt:   100,
			expectedReceiverAmt: 100,
		},
		{
			name:                "sad path - not enough amount",
			senderID:            "sender-id",
			receiverID:          "receiver-id",
			amount:              110,
			sender:              senderAccount,
			receiver:            receiverAccount,
			ctxErr:              false,
			expectedErr:         fmt.Errorf("failed to transfer as sender has insufficient balance"),
			expectedSenderAmt:   100,
			expectedReceiverAmt: 100,
		},
		{
			name:                "sad path - no account or generic get state error",
			senderID:            "sender-id",
			receiverID:          "receiver-id",
			amount:              110,
			sender:              senderAccount,
			receiver:            receiverAccount,
			ctxErr:              true,
			expectedErr:         errors.New("failed to read sender account from the world state"),
			expectedSenderAmt:   100,
			expectedReceiverAmt: 100,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//setup accounts
			senderJSON, err := json.Marshal(tt.sender)
			require.NoError(t, err)
			receiverJSON, err := json.Marshal(tt.receiver)
			require.NoError(t, err)
			ctx.Stub.StateMap[tt.senderID] = senderJSON
			ctx.Stub.StateMap[tt.receiverID] = receiverJSON
			//logs for better visibility
			log.WithFields(log.Fields{
				"beforeSenderAccount": tt.sender,
			}).Info("test name:", tt.name)
			log.WithFields(log.Fields{
				"beforeReceiverAccount": tt.receiver,
			}).Info("test name:", tt.name)

			if tt.ctxErr {
				ctx.Stub = &MockStub{StateMap: make(map[string][]byte)}
			}
			sendErr := amountTransferContract.Send(ctx, tt.senderID, tt.receiverID, tt.amount)
			switch {
			case tt.ctxErr:
				assert.Contains(t, sendErr.Error(), tt.expectedErr.Error(), "")
			default:
				assert.Equal(t, tt.expectedErr, sendErr)
				updatedSenderJSON := ctx.Stub.StateMap[tt.senderID]
				// Deserialize the updated asset
				updatedSenderAccount := new(Account)
				err = json.Unmarshal(updatedSenderJSON, updatedSenderAccount)
				require.NoError(t, err)
				// Verify the updated accounts
				assert.Equal(t, int64(tt.expectedSenderAmt), int64(updatedSenderAccount.Balance))
				updatedReceiverJSON := ctx.Stub.StateMap[tt.receiverID]
				// Deserialize the updated asset
				updatedReceiverAccount := new(Account)
				err = json.Unmarshal(updatedReceiverJSON, updatedReceiverAccount)
				require.NoError(t, err)
				assert.Equal(t, int64(tt.expectedReceiverAmt), int64(updatedReceiverAccount.Balance))
				log.WithFields(log.Fields{
					"afterSenderAccount": updatedSenderAccount,
				}).Info("test name:", tt.name)
				log.WithFields(log.Fields{
					"afterReceiverAccount": updatedReceiverAccount,
				}).Info("test name:", tt.name)
			}
		})
	}
}
func TestAmountTransferContract_GetBalance(t *testing.T) {
	// Create a new instance of the smart contract
	amountTransferContract := new(AmountTransferContract)

	// Create a new mock transaction context
	ctx := &MockTransactionContext{
		Stub:          &MockStub{StateMap: make(map[string][]byte)},
		TransactionID: "test-transaction",
		ChannelID:     "test-channel",
		Contract:      amountTransferContract,
	}
	dummyAccount := &Account{
		ID:        "account-id",
		Balance:   120,
		PublicKey: "",
	}

	tests := []struct {
		name               string
		accountID          string
		balance            uint64
		ctxErr             bool
		expectedErr        error
		expectedAccountAmt uint64
	}{
		{
			name:               "happy path",
			accountID:          "account-id",
			balance:            uint64(120),
			ctxErr:             false,
			expectedErr:        nil,
			expectedAccountAmt: 120,
		},
		{
			name:               "sad path - no account or generic get state error",
			accountID:          "sender-id",
			balance:            uint64(120),
			ctxErr:             true,
			expectedErr:        errors.New("failed to read account from the world state"),
			expectedAccountAmt: 120,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//setup accounts
			accountJSON, err := json.Marshal(dummyAccount)
			require.NoError(t, err)
			ctx.Stub.StateMap[tt.accountID] = accountJSON

			if tt.ctxErr {
				ctx.Stub = &MockStub{StateMap: make(map[string][]byte)}
			}
			getBalanceRes, getBalanceErr := amountTransferContract.GetBalance(ctx, tt.accountID)
			switch {
			case tt.ctxErr:
				assert.Contains(t, getBalanceErr.Error(), tt.expectedErr.Error(), "")
			default:
				assert.Equal(t, tt.expectedErr, getBalanceErr)
				assert.Equal(t, tt.balance, getBalanceRes)
			}
		})
	}
}
