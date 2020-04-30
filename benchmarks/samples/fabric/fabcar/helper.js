/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';



let txIndex = 0;
let pHValues = ['3','4','5','6','7','9','9','10'];
let coordinateValues = ['1.0','2.6','3.4','4.5','5.7','6.9','7.1'];
let tempValues = ['10','20','30','40'];
let timestampValues = ['12/02/2020 15:05', '01/02/2020 15:05','12/12/2018 11:05','12/02/2019 15:05','07/02/2018 15:05','09/09/2020 15:05','11/02/2019 15:05'];
let bc, contx;

module.exports.enrichLedger = async function (bc, contx, args, color, make, model, owner) {

    while (txIndex < args.assets) {
        txIndex++;
      //  let sensorID = 'Client' + contx.clientIdx + '_NODE' + txIndex.toString();
     let sensorID = txIndex.toString();
     let sensorpHValue = pHValues[Math.floor(Math.random() * pHValues.length)];
     let sensorLatValue = coordinateValues[Math.floor(Math.random() * coordinateValues.length)];
     let sensorLngValue = coordinateValues[Math.floor(Math.random() * coordinateValues.length)];
     let sensortempValue = tempValues[Math.floor(Math.random() * tempValues.length)];
     let sensortimestampValue = timestampValues[Math.floor(Math.random() * timestampValues.length)];
  
     let args = {
        chaincodeFunction: "addReading",
        chaincodeArguments: [sensorID, sensorLatValue,sensorLngValue,sensorpHValue,sensortempValue,sensortimestampValue]
      };
   
        await bc.invokeSmartContract(contx, 'fabcar', 'v1', args, 30);
    }

};
