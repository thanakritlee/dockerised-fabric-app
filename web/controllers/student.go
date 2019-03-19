package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thanakritlee/dockerised-fabric-app/web/fabric"
	u "github.com/thanakritlee/dockerised-fabric-app/web/utils"
)

// Response is the structure type template for all response.
type Response struct {
	Status int     `json:"status"`
	Data   Student `json:"data"`
}

// Student information object.
type Student struct {
	StudentID  string `json:"studentid"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	LastName   string `json:"lastname"`
	Age        int64  `json:"age"`
	DocType    string `json:"doctype"`
}

// CreateStudent is a controller to create and store a student in the blockchain.
func CreateStudent(w http.ResponseWriter, r *http.Request) {

	type request struct {
		StudentID  string `json:"studentid"`
		FirstName  string `json:"firstname"`
		MiddleName string `json:"middlename"`
		LastName   string `json:"lastname"`
		Age        int64  `json:"age"`
	}

	req := request{}

	defer r.Body.Close()

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&req)
	if isError := u.CheckError(err, w); isError {
		return
	}

	reqByte, err := json.Marshal(req)
	if isError := u.CheckError(err, w); isError {
		return
	}

	// Invoke the chaincode function `CreateStudent`
	resp, err := fabric.InvokeChaincode(reqByte, "CreateStudent")
	if isError := u.CheckError(err, w); isError {
		return
	}

	// Unmarshal the payload peer response from chaincode.
	response := Response{}
	err = json.Unmarshal(resp.Payload, &response)
	if isError := u.CheckError(err, w); isError {
		return
	}

	// Return HTTP response.
	res := u.Message("success")
	res["transactionid"] = resp.TransactionID
	res["data"] = response
	u.Response(w, res)
}
