package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/julkhong/blockchain/smartcontract/chaincode"
	log "github.com/sirupsen/logrus"
)

func main() {
	amountTransferChaincode, err := contractapi.NewChaincode(&chaincode.AmountTransferContract{})
	if err != nil {
		log.Panicf("Error creating amount transfer chaincode: %v", err)
	}
	if startErr := amountTransferChaincode.Start(); startErr != nil {
		log.Panicf("Error starting amount transfer chaincode: %v", startErr)
	}
}
