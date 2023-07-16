package smartcontract

import "github.com/pkg/errors"

type ContractInterface interface {
	Send(ctx TransactionContextInterface, senderID string, receiverID string, amount uint64) error
	GetBalance(ctx TransactionContextInterface, accountID string) (uint64, error)
}

type TransactionContextInterface interface {
	GetStub() StubInterface
	GetTransactionID() string
	GetChannelID() string
	GetContract() ContractInterface
}

type StubInterface interface {
	GetState(key string) ([]byte, error)
	PutState(key string, value []byte) error
}

type MockTransactionContext struct {
	Stub          *MockStub
	TransactionID string
	ChannelID     string
	Contract      ContractInterface
}

func (ctx *MockTransactionContext) GetStub() StubInterface {
	return ctx.Stub
}

func (ctx *MockTransactionContext) GetTransactionID() string {
	return ctx.TransactionID
}

func (ctx *MockTransactionContext) GetChannelID() string {
	return ctx.ChannelID
}

func (ctx *MockTransactionContext) GetContract() ContractInterface {
	return ctx.Contract
}

type MockStub struct {
	StateMap map[string][]byte
}

func (stub *MockStub) GetState(key string) ([]byte, error) {
	if val, ok := stub.StateMap[key]; ok {
		return val, nil
	}
	return nil, errors.New("")
}

func (stub *MockStub) PutState(key string, value []byte) error {
	stub.StateMap[key] = value
	return nil
}
