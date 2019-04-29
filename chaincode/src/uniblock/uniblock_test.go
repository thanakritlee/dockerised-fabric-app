package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
	"github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/controllers"
)

// invoke calls the chaincode function with the given arguments.
func invoke(t *testing.T, stub *shim.MockStub, args [][]byte) interface{} {

	res := stub.MockInvoke("1", args)
	require.Equal(t, int32(shim.OK), res.Status)

	result := struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}{}
	err := json.Unmarshal(res.Payload, &result)
	require.Nil(t, err)

	return result.Data

}

// TestInit tests that the chaincode init function works.
func TestInit(t *testing.T) {

	chaincode := new(controllers.UniBlock)
	stub := shim.NewMockStub("uniblock", chaincode)

	res := stub.MockInit("1", [][]byte{[]byte("init")})
	require.Equal(t, int32(shim.OK), res.Status)

}

// TestCreateStudnet tests that a student can be created.
func TestCreateStudent(t *testing.T) {

	chaincode := new(controllers.UniBlock)
	stub := shim.NewMockStub("uniblock", chaincode)

	student := controllers.Student{
		StudentID:  "0001",
		FirstName:  "FIRSTNAME",
		MiddleName: "MIDDLENAME",
		LastName:   "LASTNAME",
		Age:        69,
	}

	studentAsByte, err := json.Marshal(student)
	require.Nil(t, err)

	resp := invoke(t, stub, [][]byte{[]byte("CreateStudent"), studentAsByte})

	respStudent := controllers.Student{}
	err = mapstructure.Decode(resp, &respStudent)
	require.Nil(t, err)

	student.DocType = "Student"

	require.Equal(t, student, respStudent)

}
