package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/controllers"
)

var c controllers.UniBlock

func main() {
	uniBlock := controllers.Main()
	c = *uniBlock
}

// Init chaincode init API.
func Init(stub shim.ChaincodeStubInterface) pb.Response {
	return c.Init(stub)
}

// Invoke chaincode functions.
func Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return c.Invoke(stub)
}
