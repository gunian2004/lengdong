package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// BlockchainClient represents a client for interacting with the blockchain
type BlockchainClient struct {
	sdk           *fabsdk.FabricSDK
	channelClient *channel.Client
	channelID     string
	chaincodeName string
	orgAdmin      string
	orgName       string
}

// NewBlockchainClient creates a new blockchain client
func NewBlockchainClient(configPath, channelID, chaincodeName, orgAdmin, orgName string) (*BlockchainClient, error) {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create SDK: %v", err)
	}

	clientChannelContext := sdk.ChannelContext(channelID,
		fabsdk.WithUser(orgAdmin),
		fabsdk.WithOrg(orgName))

	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel client: %v", err)
	}

	return &BlockchainClient{
		sdk:           sdk,
		channelClient: channelClient,
		channelID:     channelID,
		chaincodeName: chaincodeName,
		orgAdmin:      orgAdmin,
		orgName:       orgName,
	}, nil
}

// AddProduct adds a product to the blockchain
func (bc *BlockchainClient) AddProduct(sku string, productData interface{}) (string, error) {
	productJSON, err := json.Marshal(productData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product data: %v", err)
	}

	req := channel.Request{
		ChaincodeID: bc.chaincodeName,
		Fcn:         "addProduct",
		Args:        [][]byte{[]byte(sku), productJSON},
	}

	resp, err := bc.channelClient.Execute(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute chaincode: %v", err)
	}

	return string(resp.TransactionID), nil
}

// UpdateLogistics updates logistics information on the blockchain
func (bc *BlockchainClient) UpdateLogistics(sku string, logisticsData interface{}) (string, error) {
	logisticsJSON, err := json.Marshal(logisticsData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal logistics data: %v", err)
	}

	req := channel.Request{
		ChaincodeID: bc.chaincodeName,
		Fcn:         "updateLogistics",
		Args:        [][]byte{[]byte(sku), logisticsJSON},
	}

	resp, err := bc.channelClient.Execute(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute chaincode: %v", err)
	}

	return string(resp.TransactionID), nil
}

// TransferOwnership transfers ownership of a product
func (bc *BlockchainClient) TransferOwnership(sku string, from, to string) (string, error) {
	req := channel.Request{
		ChaincodeID: bc.chaincodeName,
		Fcn:         "transferOwnership",
		Args:        [][]byte{[]byte(sku), []byte(from), []byte(to)},
	}

	resp, err := bc.channelClient.Execute(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute chaincode: %v", err)
	}

	return string(resp.TransactionID), nil
}

// QueryProduct queries a product from the blockchain
func (bc *BlockchainClient) QueryProduct(sku string) ([]byte, error) {
	req := channel.Request{
		ChaincodeID: bc.chaincodeName,
		Fcn:         "queryProduct",
		Args:        [][]byte{[]byte(sku)},
	}

	resp, err := bc.channelClient.Query(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query chaincode: %v", err)
	}

	return resp.Payload, nil
}

// Close closes the blockchain client
func (bc *BlockchainClient) Close() {
	bc.sdk.Close()
}
