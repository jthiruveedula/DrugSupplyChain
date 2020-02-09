/*
Drug supply chain for tracking cancer medicine
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the letter of credit
type CancerDrug struct {
	DrugId			string		`json:"drugId"`
	DrugName		string		`json:"drugName"`
	Price    string   `json:"price"`
	Manufacturer		string		`json:"manufacturer"`
	MFGDate		string		`json:"mfgdate"`
	ExpDate			string		`json:"expDate"`
	BatchCode			string		`json:"batchcode"`
	Status   string     `json:"status"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "DrugCreation" {
		return s.DrugCreation(APIstub, args)
	} else if function == "IssueAuth" {
		return s.IssuedAuth(APIstub, args)
	} else if function == "Order_Delear" {
			return s.Order_Delear(APIstub, args)
	} else if function == "Order_Hosp" {
			return s.Order_Hosp(APIstub, args)
	}else if function == "DrugHist" {
		return s.DrugHist(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}





// This function is initiate by manf 
func (s *SmartContract) DrugCreation(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	DR := CancerDrug{}

	err  := json.Unmarshal([]byte(args[0]),&DR)
if err != nil {
		return shim.Error("Not able to parse args into DR")
	}
	DR.Status = "Drug Created by MNF"
	DRBytes, err := json.Marshal(DR)
    APIstub.PutState(DR.DrugID,DRBytes)
	fmt.Println("Drug created -> ", DR)

	

	return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) IssueAuth(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	DrugId := struct {
		DrugId  string `json:"DrugId"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&DrugId)
	if err != nil {
		return shim.Error("Not able to parse args into DrugId")
	}
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	DRAsBytes, _ := APIstub.GetState(DrugId.DrugId)

	 IA := CancerDrug{}

	err = json.Unmarshal(DrugIdAsBytes, &IA)

	if err != nil {
		return shim.Error("Issue with DrugId json unmarshaling")
	}


	IA.Status= "FDACertified Drug"
	DRBytes, err := json.Marshal(IA)

	if err != nil {
		return shim.Error("Issue with DR json marshaling")
	}

    APIstub.PutState(IA.DrugId,DRBytes)
	fmt.Println("DR Issued -> ", IA)


	return shim.Success(nil)
}

func (s *SmartContract) Order_Delear(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	DrugId := struct {
		DrugId  string `json:"DrugId"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&DrugId)
	if err != nil {
		return shim.Error("Not able to parse args into DrugId")
	}

	DRAsBytes, _ := APIstub.GetState(DrugId.DrugId)

	 DR :=CancerDrug{}

	err = json.Unmarshal(DRBytes, &DR)

	if err != nil {
		return shim.Error("Issue with DR json unmarshaling")
	}


	DR.Status= "Order Success and sending ordered items to delar"
	DRAsBytes, err := json.Marshal(DR)

	if err != nil {
		return shim.Error("Issue with DR json marshaling")
	}

    APIstub.PutState(DR.DrugId,DRAsBytes)
	fmt.Println("DR Accepted -> ", DR)


	

	return shim.Success(nil)
}
func (s *SmartContract) Order_Hosp(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	DrugId := struct {
		DrugId  string `json:"DrugId"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&DrugId)
	if err != nil {
		return shim.Error("Not able to parse args into DrugId")
	}

	DRAsBytes, _ := APIstub.GetState(DrugId.DrugId)

	DR :=CancerDrug{}

	err = json.Unmarshal(DRBytes, &DR)

	if err != nil {
		return shim.Error("Issue with DR json unmarshaling")
	}


	DR.Status: "Drug reached to Hospital"
	DRAsBytes, err := json.Marshal(DR)

	if err != nil {
		return shim.Error("Issue with DR json marshaling")
	}

    APIstub.PutState(DR.DrugId,DRAsBytes)
	fmt.Println("DR Accepted -> ", DR)


	

	return shim.Success(nil)
}

func (s *SmartContract) DrugHist(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	DrugId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(DrugId)
	if err != nil {
		return shim.Error("Error retrieving Drug history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving DrugId history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- DrugHist returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
