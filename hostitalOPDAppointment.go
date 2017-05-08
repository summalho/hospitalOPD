package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"strconv"
)

type SimpleChaincode struct {
}

type Patient struct {
	PloicyId        string `json:"ploicyId"`
	City            string `json:"city"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Contact_Number  string `json:"contact_Number"`
	Hospital        string `json:"hospital"`
	AppointmentTime string `json:"appointmentTime"`
	CurrentBalance  string `json:"currentBalance"`
	BalanceUsed     string `json:"balanceUsed"`
}

func (this *Patient) convert(row *shim.Row) {
	fmt.Println("Inside cnvert")

	this.PloicyId = row.Columns[0].GetString_()
	this.City = row.Columns[1].GetString_()
	this.FirstName = row.Columns[2].GetString_()
	this.LastName = row.Columns[3].GetString_()
	this.Contact_Number = row.Columns[4].GetString_()
	this.Hospital = row.Columns[5].GetString_()
	this.AppointmentTime = row.Columns[6].GetString_()

	fmt.Println(this.PloicyId, this.AppointmentTime)
}

func main() {
	fmt.Println("Inside main")

	fmt.Println("Inside main method")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)

	}
}

//creates the patient table
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//first call to init will create the table to store patient details

	var err error

	fmt.Println("Inside init")

	err = stub.CreateTable("Patient_Details", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"PloicyId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"City", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"FirstName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"LastName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"PhoneNumber", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Hospital", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"AppointmentTime", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"CurrentBalance", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"BalanceUsed", shim.ColumnDefinition_STRING, false}})

	if err != nil {
		return nil, errors.New("Table cannot be created")
	}

	var str string
	str = "table created"
	bytesreturn := []byte(str)

	fmt.Println("exiting init")

	return bytesreturn, nil

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	var recordAdded []byte

	fmt.Println("Inside invoke")

	if function == "addPatientInTable" {
		recordAdded, err = addPatientInTable(stub, args)
	}
	/*if function == "UpdatePatientBalance" {
		recordAdded, err = UpdatePatientBalance(stub, args)
	}*/
	if err != nil {
		return nil, errors.New("Record cannot be added")
	}
	return recordAdded, nil

}

/*func UpdatePatientBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var currentBalanceInt int
	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}

	columns = append(columns, col1)

	row, err := stub.GetRow(args[0], columns)

	if err != nil {
		return nil, fmt.Errorf("getRows operation failed. %s", err)
	}

	var cols *shim.Column
	cols = row.GetColumns()[7]
	var currentBalance = cols.GetString_()

	currentBalanceInt = strconv.Atoi(currentBalance)

	currentBalanceInt = currentBalanceInt - 10

	currentBalance = strconv.Itoa(currentBalanceInt)

	row.Columns[7].Value = currentBalance

	stub.ReplaceRow(args[0], row)

}*/

/*func GetPatientBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column

	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}

	columns = append(columns, col1)

	row, err := stub.GetRow(args[0], columns)

	if err != nil {
		return nil, fmt.Errorf("getRows operation failed. %s", err)
	}

	var cols *shim.Column
	cols = row.GetColumns()[7]
	var currentBalance = cols.GetString_()

	return []byte(currentBalance), nil

}*/

//addPatient in table
func addPatientInTable(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []*shim.Column

	fmt.Println("inside ass patientdetails in table")

	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}

	col2 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}

	col3 := shim.Column{Value: &shim.Column_String_{String_: args[3]}}

	col4 := shim.Column{Value: &shim.Column_String_{String_: args[4]}}

	col5 := shim.Column{Value: &shim.Column_String_{String_: args[5]}}

	col6 := shim.Column{Value: &shim.Column_String_{String_: args[6]}}

	col7 := shim.Column{Value: &shim.Column_String_{String_: args[7]}}

	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)
	columns = append(columns, &col7)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow(args[0], row)

	fmt.Println(ok)

	if err != nil {
		return nil, fmt.Errorf("insertRow operation failed. %s", err)
	}

	if !ok {
		return nil, errors.New("insertRow operation failed. Row with given key already exists")
	}

	fmt.Println("Exiting patient details after adding ")
	return nil, nil

}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside query")

	var err error
	var patientDetails []byte
	if function == "GetPatientDetails" {
		patientDetails, err = GetPatientDetails(stub, args)
	}
	/*if function == "UpdatePatientBalance" {
		patientDetails, err = UpdatePatientBalance(stub, args)
	}*/
	if err != nil {
		return nil, errors.New("Record cannot be added")
	}
	return patientDetails, nil

}

func GetPatientDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column

	fmt.Println("Inside get patient details")

	if len(args) != 2 {
		return nil, errors.New("This method requires two arguments. Table name and primary key attriubute")
	}

	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}

	columns = append(columns, col1)

	row, err := stub.GetRow(args[0], columns)

	if err != nil {
		return nil, errors.New("Get row operation failed")
	}

	if len(row.Columns) == 0 {
		return nil, errors.New("no row returned")

	}
	fmt.Println("--------------------------------------\n")

	var patientObject *Patient
	var patientObjectList []*Patient

	patientObject = new(Patient)

	patientObject.convert(&row)

	patientObjectList = append(patientObjectList, patientObject)

	jsonPatientObjectList, err := json.Marshal(patientObjectList)

	if err != nil {
		return nil, errors.New("Error arshalling Json")
	}

	fmt.Println(string(jsonPatientObjectList))

	return jsonPatientObjectList, nil

}
