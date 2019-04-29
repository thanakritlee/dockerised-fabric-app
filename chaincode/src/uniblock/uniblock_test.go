package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/controllers"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {

	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed: ", string(res.Message))
		t.FailNow()
	}

}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {

	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke failed", string(res.Message))
		t.FailNow()
	}

	result := struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}{}
	err := json.Unmarshal(res.Payload, &result)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	if result.Status != 200 {
		fmt.Println("Error: ", result.Data)
		t.FailNow()
	}

}

func TestInit(t *testing.T) {

	chaincode := new(controllers.UniBlock)
	stub := shim.NewMockStub("uniblock", chaincode)

	checkInit(t, stub, [][]byte{[]byte("init")})

}
