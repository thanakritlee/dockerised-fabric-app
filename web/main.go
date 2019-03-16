package main

import (
	"fmt"
	"os"

	"github.com/thanakritlee/dockerised-fabric-app/web/fabric"
)

func main() {
	f := fabric.Fabric{
		OrdererID:       "orderer0.example.com",
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

	err := f.Initialise()
	if err != nil {
		fmt.Printf("Unable to initialise the Fabric SDK: %v\n", err)
		return
	}

	resp, err := f.InvokeChaincode()
	if err != nil {
		fmt.Printf("Unable to invoke chaincode: %v\n", err)
	}
	fmt.Println(resp)
}
