package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/thanakritlee/dockerised-fabric-app/web/fabric"
	"github.com/thanakritlee/dockerised-fabric-app/web/router"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	channelConfigPath := filepath.Dir(dir) + "/fabric-network/config/channel.tx"
	chaincodePath := filepath.Dir(dir) + "/chaincode/"

	f := fabric.Fabric{
		OrdererID:       "orderer0.example.com",
		CaID:            "ca.org1.example.com",
		ChannelID:       "channel",
		ChannelConfig:   channelConfigPath,
		ChainCodeID:     "uniblock",
		ChaincodeGoPath: chaincodePath,
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		OrgMSP:          "Org1MSP",
		OrdererOrgName:  "OrdererOrg",
		ConfigFile:      "config.yaml",
		UserName:        "Admin",
		Initialised:     true,
	}

	if len(os.Args) > 1 && os.Args[1] == "init" {
		f.Initialised = false
	}

	err = f.Initialise()
	if err != nil {
		log.Fatal("Unable to initialise the Fabric SDK: ", err)
	}

	router := router.GetRouter()

	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("http server started on :%s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
