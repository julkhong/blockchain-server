package smartcontract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	log "github.com/sirupsen/logrus"
)

func main() {
	amountTransferChaincode, err := contractapi.NewChaincode(&AmountTransferContract{})
	if err != nil {
		log.Panicf("Error creating amount transfer chaincode: %v", err)
	}
	if startErr := amountTransferChaincode.Start(); startErr != nil {
		log.Panicf("Error starting amount transfer chaincode: %v", startErr)
	}
}
