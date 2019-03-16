package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/utils"
)

// UniBlock chaincode implementation
type UniBlock struct {
}

// Main entry function for the chaincode controller.
func Main() (c *UniBlock) {
	err := shim.Start(new(UniBlock))
	if err != nil {
		fmt.Printf("Error starting UniBlock chaincode: %s", err)
		return nil
	}
	return new(UniBlock)
}

// Init chaincode init API.
func (c *UniBlock) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke chaincode functions.
func (c *UniBlock) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var res utils.Response
	var err error

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is running: " + function)

	switch function {
	case "CreateStudent":
		res, err = c.CreateStudent(stub, args)
	default:
		fmt.Println("Invoke did not find the function: " + function)
		return shim.Error("Recieved unknown function invocation")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	resAsBytes, err := json.Marshal(res)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(resAsBytes)
}
