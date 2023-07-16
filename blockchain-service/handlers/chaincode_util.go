package handlers

//func initializeSDK() *fabsdk.FabricSDK {
//	// Load configuration from file
//	configProvider := config.FromFile("config.yaml")
//
//	// Initialize the SDK
//	sdk, err := fabsdk.New(configProvider)
//	if err != nil {
//		log.Fatal("Failed to create SDK: ", err)
//	}
//
//	return sdk
//}
//
//func createClient(sdk *fabsdk.FabricSDK, channelID, chaincodeID, username string) (*channel.Client, contextAPI.ClientProvider) {
//	// Create channel client
//	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(username))
//
//	channelClient, err := channel.New(clientChannelContext)
//	if err != nil {
//		log.Fatal("Failed to create channel client: ", err)
//	}
//
//	// Get a user context for performing the transaction
//	clientContext := sdk.Context(fabsdk.WithUser(username))
//
//	return channelClient, clientContext
//}
//
//func invokeSmartContract(client *channel.Client, ctx contextAPI.ClientProvider, chaincodeID, functionName string, args [][]byte) ([]byte, error) {
//	request := channel.Request{
//		ChaincodeID: chaincodeID,
//		Fcn:         functionName,
//		Args:        args,
//	}
//
//	response, err := client.Execute(request, channel.WithRetry(retry.DefaultChannelOpts))
//	if err != nil {
//		return nil, fmt.Errorf("failed to execute smart contract function: %v", err)
//	}
//
//	return response.Payload, nil
//}
//
//func querySmartContract(client *channel.Client, ctx contextAPI.ClientProvider, chaincodeID, functionName string, args [][]byte) ([]byte, error) {
//	request := channel.Request{
//		ChaincodeID: chaincodeID,
//		Fcn:         functionName,
//		Args:        args,
//	}
//
//	response, err := client.Query(request)
//	if err != nil {
//		return nil, fmt.Errorf("failed to query smart contract function: %v", err)
//	}
//
//	return response.Payload, nil
//}
