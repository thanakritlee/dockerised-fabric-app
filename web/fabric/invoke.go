package fabric

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/thanakritlee/fabric-sdk-go/pkg/client/channel"
	msp "github.com/thanakritlee/fabric-sdk-go/pkg/client/msp"
	"github.com/thanakritlee/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/thanakritlee/fabric-sdk-go/pkg/core/config"
	"github.com/thanakritlee/fabric-sdk-go/pkg/fabsdk"
)

//InvokeChaincode will invoke a function inside the chaincode.
func InvokeChaincode(req []byte, function string) (channel.Response, error) {

	fabric := Fabric{
		OrdererID:       "orderer0.example.com",
		CaID:            "ca.org1.example.com",
		ChannelID:       "channel",
		ChannelConfig:   os.Getenv("GOPATH") + "/src/github.com/thanakritlee/dockerised-fabric-app/fabric-network/config/channel.tx",
		ChainCodeID:     "uniblock",
		ChaincodeGoPath: os.Getenv("GOPATH") + "/src/github.com/thanakritlee/dockerised-fabric-app/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		OrgMSP:          "Org1MSP",
		OrdererOrgName:  "OrdererOrg",
		ConfigFile:      "config.yaml",
		UserName:        "Admin",
	}

	fmt.Println("Invoking chaincode...")

	// Create a new Fabric SDK instance using config from ConfigFile.
	sdk, err := fabsdk.New(config.FromFile(fabric.ConfigFile))
	if err != nil {
		return channel.Response{}, errors.WithMessage(err, "failed to create SDK")
	}

	fabric.sdk = sdk
	// Close the SDK and release resources when done (returned).
	defer fabric.sdk.Close()

	// Register and Enroll user.
	mspClient, err := msp.New(fabric.sdk.Context())
	if err != nil {
		return channel.Response{}, err
	}

	username := "user0"

	signingIdentity, err := mspClient.GetSigningIdentity(username)
	if err != nil {
		return channel.Response{}, err
	}

	// Create a Fabric channel context.
	clientChannelContext := fabric.sdk.ChannelContext(
		fabric.ChannelID,
		fabsdk.WithUser(username),
		fabsdk.WithOrg(fabric.OrgName),
		fabsdk.WithIdentity(signingIdentity),
	)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return channel.Response{}, err
	}

	fabric.client = client

	//Prepare and execute the chaincode transaction.
	resp, err := fabric.client.Execute(channel.Request{
		ChaincodeID: fabric.ChainCodeID,
		Fcn:         function,
		Args:        [][]byte{req},
	}, channel.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return resp, err
	}

	fmt.Println("Invoked chaincode")

	return resp, nil

}
