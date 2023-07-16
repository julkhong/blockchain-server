package smartcontract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/julkhong/blockchain/blockchain-server/blockchain-service/logic"
	log "github.com/sirupsen/logrus"

)

func main() {
	amountTransferChaincode, err := contractapi.NewChaincode(&logic.AmountTransferContract{})
	if err != nil {
		log.Panicf("Error creating amount transfer chaincode: %v", err)
	}
	if startErr := amountTransferChaincode.Start(); startErr != nil{
		log.Panicf("Error starting amount transfer chaincode: %v", startErr)
	}
	Update addresses if they have been overwritten by flags
}
