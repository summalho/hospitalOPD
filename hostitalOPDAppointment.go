package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type SimpleChaincode struct {
}

type Patient struct {
	PloicyId         string `json:"ploicyId"`
	City             string `json:"city"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Contact_Number   int64  `json:"contact_Number"`
	Hospital         string `json:"hospital"`
	AppointmentTime  string `json:"appointmentTime"`
	UnavailedBalance int64  `json:"unavailedBalance"`
	ClaimedAmount    int64  `json:"claimedAmount"`
}

type Hospital struct {
	PloicyId         string `json:"ploicyId"`
	Hospital         string `json:"hospital"`
	AppointmentTime  string `json:"appointmentTime"`
	UnavailedBalance int64  `json:"unavailedBalance"`
	ClaimedAmount    int64  `json:"claimedAmount"`
}

func (this *Patient) convert(row *shim.Row) {
	fmt.Println("Inside convert")

	this.PloicyId = row.Columns[0].GetString_()
	this.City = row.Columns[1].GetString_()
	this.FirstName = row.Columns[2].GetString_()
	this.LastName = row.Columns[3].GetString_()
	this.Contact_Number = row.Columns[4].GetInt64()
	this.Hospital = row.Columns[5].GetString_()
	this.AppointmentTime = row.Columns[6].GetString_()

	fmt.Println(this.PloicyId, this.AppointmentTime)
}

func (this *Hospital) convertHospitalEntries(row *shim.Row) {
	fmt.Println("Inside convert")

	this.PloicyId = row.Columns[0].GetString_()
	this.Hospital = row.Columns[1].GetString_()
	this.AppointmentTime = row.Columns[2].GetString_()
	this.UnavailedBalance = row.Columns[3].GetInt64()
	this.ClaimedAmount = row.Columns[4].GetInt64()

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

	fmt.Println("Inside Init")

	err = stub.CreateTable("Patient_Details", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"PloicyId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"City", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"FirstName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"LastName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"PhoneNumber", shim.ColumnDefinition_INT64, false},
		&shim.ColumnDefinition{"Hospital", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"AppointmentTime", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Unavailed_Balance", shim.ColumnDefinition_INT64, false},
		&shim.ColumnDefinition{"Claimed_Amount", shim.ColumnDefinition_INT64, false}})

	err = stub.CreateTable("Hospital_Details", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"PloicyId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"Hospital", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"AppointmentTime", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Unavailed_Balance", shim.ColumnDefinition_INT64, false},
		&shim.ColumnDefinition{"Claimed_Amount", shim.ColumnDefinition_INT64, false}})

	if err != nil {
		return nil, errors.New("Table cannot be created")
	}

	var str string
	str = "table created"
	bytesreturn := []byte(str)

	fmt.Println("exiting Init")

	return bytesreturn, nil

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	var recordAdded []byte

	fmt.Println("Inside invoke")

	if function == "addPatientInTable" {
		recordAdded, err = addPatientInTable(stub, args)
	}
	if function == "UpdatePatientBalance" {
		recordAdded, err = UpdatePatientBalance(stub, args)
	}
	if err != nil {
		return nil, errors.New("Record cannot be added")
	}
	return recordAdded, nil

}

func UpdatePatientBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	var newColumns []*shim.Column
	var newColumnsForHospitalTable []*shim.Column

	Oldcol1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	columns = append(columns, Oldcol1)

	row, err := stub.GetRow("Patient_Details", columns)

	if err != nil {
		return nil, fmt.Errorf("getRow operation failed. %s", err)
	}

	unAvailedBalanceInt := row.GetColumns()[7].GetInt64()
	fmt.Println("Unavailed Balane ", unAvailedBalanceInt)

	claimedAmount := args[2]
	claimedAmountInt, _ := strconv.ParseInt(claimedAmount, 10, 64)

	if unAvailedBalanceInt >= claimedAmountInt {
		unAvailedBalanceInt = unAvailedBalanceInt - claimedAmountInt
	}
	if unAvailedBalanceInt < claimedAmountInt {
		unAvailedBalanceInt = 0
	}

	i := int(unAvailedBalanceInt)
	fmt.Println("Unavailed Balance after transaction  ", strconv.Itoa(i))
	fmt.Println("claimed Amount  ", claimedAmount)

	col0 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[0].GetString_()}}
	col1 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[1].GetString_()}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[2].GetString_()}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[3].GetString_()}}
	col4 := shim.Column{Value: &shim.Column_Int64{Int64: row.GetColumns()[4].GetInt64()}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[5].GetString_()}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: row.GetColumns()[6].GetString_()}}
	col7 := shim.Column{Value: &shim.Column_Int64{Int64: unAvailedBalanceInt}}
	col8 := shim.Column{Value: &shim.Column_Int64{Int64: claimedAmountInt}}

	newColumns = append(newColumns, &col0)
	newColumns = append(newColumns, &col1)
	newColumns = append(newColumns, &col2)
	newColumns = append(newColumns, &col3)
	newColumns = append(newColumns, &col4)
	newColumns = append(newColumns, &col5)
	newColumns = append(newColumns, &col6)
	newColumns = append(newColumns, &col7)
	newColumns = append(newColumns, &col8)

	newColumnsForHospitalTable = append(newColumnsForHospitalTable, &col0)
	newColumnsForHospitalTable = append(newColumnsForHospitalTable, &col5)
	newColumnsForHospitalTable = append(newColumnsForHospitalTable, &col6)
	newColumnsForHospitalTable = append(newColumnsForHospitalTable, &col7)
	newColumnsForHospitalTable = append(newColumnsForHospitalTable, &col8)

	rowReplaced := shim.Row{Columns: newColumns}
	ok, err := stub.ReplaceRow("Patient_Details", rowReplaced)

	if err != nil {
		return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
	}
	if !ok {
		return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
	}

	hospitalRowReplaced := shim.Row{Columns: newColumnsForHospitalTable}
	ok, err = stub.ReplaceRow("Patient_Details", hospitalRowReplaced)

	if err != nil {
		return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
	}
	if !ok {
		return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
	}

	return []byte("Updated Balance in Network"), nil
}

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
	var hospitalTableCols []*shim.Column

	fmt.Println("inside add patientdetails in table")

	col4Int, _ := strconv.ParseInt(args[4], 10, 64)
	col7Int, _ := strconv.ParseInt(args[7], 10, 64)
	col8Int, _ := strconv.ParseInt(args[8], 10, 64)

	col0 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
	col1 := shim.Column{Value: &shim.Column_String_{String_: args[1]}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: args[2]}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: args[3]}}
	col4 := shim.Column{Value: &shim.Column_Int64{Int64: col4Int}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: args[5]}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: args[6]}}
	col7 := shim.Column{Value: &shim.Column_Int64{Int64: col7Int}}
	col8 := shim.Column{Value: &shim.Column_Int64{Int64: col8Int}}

	columns = append(columns, &col0)
	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)
	columns = append(columns, &col7)
	columns = append(columns, &col8)

	hospitalTableCols = append(hospitalTableCols, &col0)
	hospitalTableCols = append(hospitalTableCols, &col5)
	hospitalTableCols = append(hospitalTableCols, &col6)
	hospitalTableCols = append(hospitalTableCols, &col7)
	hospitalTableCols = append(hospitalTableCols, &col8)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow("Patient_Details", row)

	rowString1 := fmt.Sprintf("%s", row)

	fmt.Println("rowString1 ", rowString1)

	fmt.Println(ok)

	if err != nil {
		return nil, fmt.Errorf("insertRow operation failed. %s", err)
	}

	if !ok {
		return nil, errors.New("insertRow operation failed. Row with given key already exists")
	}

	hospitalRow := shim.Row{Columns: hospitalTableCols}
	ok, err = stub.InsertRow("Hospital_Details", hospitalRow)

	rowString := fmt.Sprintf("%s", hospitalRow)
	fmt.Println("rowString for hospital Table", rowString)

	if err != nil {
		return nil, fmt.Errorf("insertRow operation failed. %s", err)
	}

	if !ok {
		return nil, errors.New("insertRow operation failed. Row with given key already exists")
	}

	fmt.Println("Exiting patient details after adding in Patient Table and Hospital table ")
	return []byte(rowString), nil

}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside query")

	var err error
	var patientDetails []byte
	if function == "getAllPatientsForHospital" {
		patientDetails, err = getAllPatientsForHospital(stub, args)
	}
	/*if function == "getPatientsForPolicyId" {
		patientDetails, err = getPatientsForPolicyId(stub, args)
	}
	if function == "viewAppointments" {
		patientDetails, err = viewAppointments(stub, args)
	}*/

	if err != nil {
		return nil, errors.New("Record cannot be added")
	}
	return patientDetails, nil

}

func getAllPatientsForHospital(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	fmt.Println("Inside get patient details")

	col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}

	columns = append(columns, col1)

	hospitalObjectList := []*Hospital{}

	rowChannel, err := stub.GetRows("Hospital_Details", columns)

	fmt.Println("rowChannel : ", rowChannel)

	//var rows []shim.Row

	for {
		select {
		case row, ok := <-rowChannel:
			fmt.Println("OK Status : ", ok)
			if !ok {
				rowChannel = nil
			} else {
				hospitalObject := new(Hospital)
				hospitalObject.convertHospitalEntries(&row)
				hospitalObjectList = append(hospitalObjectList, hospitalObject)
				//rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	jsonRows, _ := json.Marshal(hospitalObjectList)

	//jsonHospitalObjectList, err := json.Marshal(hospitalObjectList)

	if err != nil {
		return nil, errors.New("Error marshalling Json")
	}

	fmt.Println(string(jsonRows))

	return jsonRows, nil

}
