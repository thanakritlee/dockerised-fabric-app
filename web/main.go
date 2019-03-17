package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thanakritlee/dockerised-fabric-app/web/fabric"
	"github.com/thanakritlee/dockerised-fabric-app/web/router"
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

	router := router.GetRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("http server started on :%s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
