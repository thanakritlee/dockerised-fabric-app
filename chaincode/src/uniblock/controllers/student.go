package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Student information object.
type Student struct {
	StudentID  string `json:"studentid"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	LastName   string `json:"lastname"`
	Age        int64  `json:"age"`
	DocType    string `json:"doctype"`
}

// CreateStudent create and store a student object into the blockchain.
func (c *UniBlock) CreateStudent(stub shim.ChaincodeStubInterface, args []string) (utils.Response, error) {

	type request struct {
		StudentID  string `json:"studentid"`
		FirstName  string `json:"firstname"`
		MiddleName string `json:"middlename"`
		LastName   string `json:"lastname"`
		Age        int64  `json:"age"`
	}

	req := request{}

	err := json.Unmarshal([]byte(args[0]), &req)
	if err != nil {
		return utils.Response{}, err
	}

	student := Student{
		StudentID:  req.StudentID,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Age:        req.Age,
		DocType:    "Student",
	}

	fmt.Println("Creating student...")

	err = utils.UtilPOST(stub, "student", "studentid", student.StudentID, student, fmt.Errorf("Student %v already exist", student.StudentID))
	if err != nil {
		return utils.Response{}, err
	}

	fmt.Println("Created student")

	return utils.Response{
		Status: 200,
		Data:   student,
	}, nil

}
