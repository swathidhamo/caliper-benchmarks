/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the Sensor structure, with 3 properties.  Structure tags are used by encoding/json library 
type Sensor struct {
	Model  string `json:"model"`
	TimeStamp string `json:"timestamp"`
	Value string `json:"value"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "getReadingForID" {
		return s.getReadingForID(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "addReading" {
		return s.addReading(APIstub, args)
	} else if function == "getReading" {
		return s.getReading(APIstub)
	} else if function == "updateReading" {
		return s.updateReading(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) getReadingForID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sensorAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(sensorAsBytes)
}
//here
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	sensors := []Sensor{
		Sensor{Model: "pH", TimeStamp:"12/02/2020 15:05", Value: "23.2"},
		Sensor{Model: "pH", TimeStamp:"15/02/2020 15:05", Value: "130.2"},
		Sensor{Model: "pH", TimeStamp:"14/02/2020 15:05", Value: "34.2"},
		Sensor{Model: "pH", TimeStamp:"13/02/2020 15:05", Value: "1.2"},
		Sensor{Model: "pH", TimeStamp:"16/02/2020 15:05", Value: "0.2"},
		Sensor{Model: "pH", TimeStamp:"17/02/2020 15:05", Value: "67.2"},
		Sensor{Model: "pH", TimeStamp:"18/02/2020 15:05", Value: "134.2"},
		Sensor{Model: "pH", TimeStamp:"20/02/2109 15:05", Value: "72.2"},
		Sensor{Model: "pH", TimeStamp:"21/02/2020 15:05", Value: "90.2"},
	}

	i := 0
	for i < len(sensors) {
		fmt.Println("i is ", i)
		sensorAsBytes, _ := json.Marshal(sensors[i])
		APIstub.PutState("NODE"+strconv.Itoa(i), sensorAsBytes)
		fmt.Println("Added", sensors[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) addReading(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
//here
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var sensor = Sensor{TimeStamp: args[1], Model: args[2], Value: args[3]}

	sensorAsBytes, _ := json.Marshal(sensor)
	APIstub.PutState(args[0], sensorAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getReading(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "NODE0"
	endKey := "NODE999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record values\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) updateReading(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}


	sensorAsBytes, _ := APIstub.GetState(args[0])
	sensor := Sensor{}

	json.Unmarshal(sensorAsBytes, &sensor)
	sensor.Value = args[1]
   
	sensorAsBytes, _ = json.Marshal(sensor)
	APIstub.PutState(args[0], sensorAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
