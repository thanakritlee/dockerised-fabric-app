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

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&req)
	u.CheckError(err, w)

	reqByte, err := json.Marshal(req)
	u.CheckError(err, w)

	// Invoke the chaincode function `CreateStudent`
	resp, err := fabric.InvokeChaincode(reqByte, "CreateStudent")
	u.CheckError(err, w)

	// Unmarshal the payload peer response from chaincode.
	response := Response{}
	err = json.Unmarshal(resp.Payload, &response)
	u.CheckError(err, w)

	// Return HTTP response.
	res := u.Message("success")
	res["transactionid"] = resp.TransactionID
	res["data"] = response
	u.Response(w, res)
}
