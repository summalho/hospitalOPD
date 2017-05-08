package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestGetPatientDetails(t *testing.T) {

	fmt.Println("Inside test method")
	//var tableName = "Patient_Details"
	//var id = "1234"

	stub := shim.NewMockStub("mockstub", new(SimpleChaincode))

	if stub == nil {
		t.Fatalf("mockStub creation failed")
	}

	stub.MockTransactionStart("t2")

	_, err := GetPatientDetails(stub, []string{})

	if err != nil {

		t.Fatalf(err.Error())

	}

	stub.MockTransactionEnd("t2")

}
