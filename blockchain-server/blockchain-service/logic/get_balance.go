package logic

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
)

// GetBalance ...
func GetBalance(channelClient *channel.Client, clientContext contextAPI.ClientProvider, accountID string, channelID string) {
	// Prepare the request
	request := channel.Request{
		ChaincodeID: channelID,
		Fcn:         "getbalance",
		Args:        [][]byte{[]byte(accountID)},
	}

	// Invoke the "send" function in the smart contract
	response, err := channelClient.Execute(request, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Printf("Failed to invoke 'send' function in the smart contract: %s\n", err)
		return
	}

	// Process the response from the smart contract
	fmt.Printf("Response from smart contract: %s\n", string(response.Payload))
}
