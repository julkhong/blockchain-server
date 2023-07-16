package logic

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
)

// Send ...
func Send(channelClient *channel.Client, clientContext contextAPI.ClientProvider, senderID, receiverID string, amount uint64, channelID string) {
	// Prepare the request
	request := channel.Request{
		ChaincodeID: channelID,
		Fcn:         "send",
		Args:        [][]byte{[]byte(senderID), []byte(receiverID), []byte(fmt.Sprintf("%d", amount))},
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
