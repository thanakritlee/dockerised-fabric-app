package utils

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Response is the structure type template for all response.
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// UtilPOST is a utility function that when called will create an asset on the blockchain state.
func UtilPOST(stub shim.ChaincodeStubInterface, assetNameString string, assetIDString string, assetID string, asset interface{}, errorMessage error) error {
	// Create a composite key for the asset.
	compKey, err := stub.CreateCompositeKey(assetNameString, []string{assetIDString, assetID})
	if err != nil {
		return err
	}

	// Check if the composite key already exists.
	// Basically checking if the ID has already been taken.
	assetAsBytes, err := stub.GetState(compKey)
	if err != nil {
		return err
	} else if assetAsBytes != nil {
		// Display a custom error message if the ID has already
		// been taken.
		return errorMessage
	}

	// Package the asset GoLang object into bytes.
	assetAsBytes, err = json.Marshal(asset)
	if err != nil {
		return err
	}

	// Save the asset to the blockchain state.
	err = stub.PutState(compKey, assetAsBytes)
	if err != nil {
		return err
	}

	return nil
}

// UtilPUT is a utility function that when called will update an asset on the blockchain state.
func UtilPUT(stub shim.ChaincodeStubInterface, assetNameString string, assetIDString string, assetID string, asset interface{}, errorMessage error) error {
	// Create a composite key for the asset.
	compKey, err := stub.CreateCompositeKey(assetNameString, []string{assetIDString, assetID})
	if err != nil {
		return err
	}

	// Check if the composite key already exists.
	// Basically checking if the ID has already been taken.
	assetAsBytes, err := stub.GetState(compKey)
	if err != nil {
		return err
	} else if assetAsBytes == nil {
		// Display a custom error message when asset does not exist.
		return errorMessage
	}

	// Package the asset GoLang object into bytes.
	assetAsBytes, err = json.Marshal(asset)
	if err != nil {
		return err
	}

	// Update the asset on the blockchain state.
	err = stub.PutState(compKey, assetAsBytes)
	if err != nil {
		return err
	}

	return nil
}

// UtilGET is a utility functiion that when called will call a getter function of a given asset, and get the asset from the blockchain state.
func UtilGET(stub shim.ChaincodeStubInterface, requestStruct interface{}, function func(shim.ChaincodeStubInterface, []string) (Response, error)) (Response, error) {
	getRequestAsBytes, err := json.Marshal(requestStruct)
	if err != nil {
		return Response{}, err
	}

	// Get the asset using the asset id.
	res, err := function(stub, []string{string(getRequestAsBytes)})
	if err != nil {
		return Response{}, err
	}
	return res, nil
}
