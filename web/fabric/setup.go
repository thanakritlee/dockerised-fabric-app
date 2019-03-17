package fabric

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
	"github.com/thanakritlee/fabric-sdk-go/pkg/core/config"
)

// Fabric implementaion
type Fabric struct {
	ConfigFile      string
	OrgID           string
	OrdererID       string
	ChannelID       string
	ChainCodeID     string
	initialised     bool
	ChannelConfig   string
	ChaincodeGoPath string
	OrgAdmin        string
	OrgName         string
	OrgMSP          string
	OrdererOrgName  string
	UserName        string
	client          *channel.Client
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	event           *event.Client
}

// Initialise sets up the Fabric network and connect the SDK to it.
func (fabric *Fabric) Initialise() error {

	if fabric.initialised {
		return errors.New("sdk is already intialised")
	}

	sdk, err := fabsdk.New(config.FromFile(fabric.ConfigFile))

	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}

	fabric.sdk = sdk
	fmt.Println("SDK created")

	err = createChannel(fabric)
	if err != nil {
		return errors.Cause(err)
	}

	err = joinChannel(fabric)
	if err != nil {
		return errors.Cause(err)
	}

	err = installChaicode(fabric)
	if err != nil {
		return errors.Cause(err)
	}

	err = instantiateChaincode(fabric)
	if err != nil {
		return errors.Cause(err)
	}

	return nil

}

func (fabric *Fabric) closeSDK() error {
	if !fabric.initialised {
		return errors.New("sdk is not initialised")
	}

	fabric.sdk.Close()

	return nil
}

func createChannel(fabric *Fabric) error {
	fmt.Println("Creating channel...")

	//clientContext allows creation of transactions using the supplied identity as the credential.
	clientContext := fabric.sdk.Context(fabsdk.WithUser(fabric.OrgAdmin), fabsdk.WithOrg(fabric.OrdererOrgName))

	resMgmtClient, err := resmgmt.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client")
	}

	mspClient, err := mspclient.New(fabric.sdk.Context(), mspclient.WithOrg(fabric.OrgName))
	if err != nil {
		return errors.Cause(err)
	}

	adminIdentity, err := mspClient.GetSigningIdentity(fabric.OrgAdmin)
	if err != nil {
		return errors.Cause(err)
	}
	req := resmgmt.SaveChannelRequest{
		ChannelID:         fabric.ChannelID,
		ChannelConfigPath: fabric.ChannelConfig,
		SigningIdentities: []msp.SigningIdentity{adminIdentity},
	}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(fabric.OrdererID))
	if err != nil && txID.TransactionID == "" {
		return errors.WithMessage(err, "failed to save channel")
	}

	fmt.Println("Created channel")

	return nil

}

func joinChannel(fabric *Fabric) error {
	fmt.Println("Joining channel...")

	// Create Admin context for the Organisation.
	adminContext := fabric.sdk.Context(fabsdk.WithUser(fabric.OrgAdmin), fabsdk.WithOrg(fabric.OrgName))

	// Create a Fabric Client.
	orgResMgmt, err := resmgmt.New(adminContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new resource management client")
	}

	fabric.admin = orgResMgmt

	// Join all peers to the channel.
	if err = orgResMgmt.JoinChannel(fabric.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(fabric.OrdererID)); err != nil {
		return errors.WithMessage(err, "org peers failed to join channel")
	}

	fmt.Println("Joined channel")

	return nil

}

func installChaicode(fabric *Fabric) error {
	fmt.Println("Installing chaincode")

	ccPkg, err := packager.NewCCPackage(fabric.ChainCodeID, fabric.ChaincodeGoPath)
	if err != nil {
		return errors.Cause(err)
	}

	// Install the chaincode to the organisation peers.
	installCCReq := resmgmt.InstallCCRequest{
		Name:    fabric.ChainCodeID,
		Path:    fabric.ChainCodeID,
		Version: "0",
		Package: ccPkg,
	}
	_, err = fabric.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.Cause(err)
	}

	fmt.Println("Installed chaincode")

	return nil

}

func instantiateChaincode(fabric *Fabric) error {
	fmt.Println("Instantiating chaincode")

	// Set up the chaincode policy.
	ccPolicy := cauthdsl.SignedByAnyMember([]string{fabric.OrgMSP})

	// Instantiate chaincode on the channel using the organisation resource manager.
	resp, err := fabric.admin.InstantiateCC(
		fabric.ChannelID,
		resmgmt.InstantiateCCRequest{
			Name:    fabric.ChainCodeID,
			Path:    fabric.ChainCodeID,
			Version: "0",
			Args:    [][]byte{},
			Policy:  ccPolicy,
		},
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
	)
	if err != nil && resp.TransactionID == "" {
		return errors.Cause(err)
	}

	fmt.Println("Instantiated chaincode")

	return nil

}
