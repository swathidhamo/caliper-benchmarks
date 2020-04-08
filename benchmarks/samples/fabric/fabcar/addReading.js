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

module.exports.info = 'Adding readings.';

let txIndex = 0;
let models = ['pH','Humidity','Temperature','Visual'];
let timestamps = ['12/02/2020 15:05', '01/02/2020 15:05','12/12/2018 11:05','12/02/2019 15:05','07/02/2018 15:05','09/09/2020 15:05','11/02/2019 15:05'];
let values = ['0.5','6.8','2,7','1.22','2.34','5.6','0.12','0.112'];
let bc, contx;

module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};

module.exports.run = function() {
    txIndex++;
    let sensorID = 'Client' + contx.clientIdx + '_NODE' + txIndex.toString();
    let sensorModel = models[Math.floor(Math.random() * models.length)];
    let sensorTimestamp= timestamps[Math.floor(Math.random() * timestamps.length)];
    let sensorValue = values[Math.floor(Math.random() * values.length)];
    let args = {
        chaincodeFunction: "addReading",
        chaincodeArguments: [sensorID, sensorModel, sensorTimestamp, sensorValue]
    };

    return bc.invokeSmartContract(contx, 'fabcar', 'v1', args, 30);
};

module.exports.end = function() {
    return Promise.resolve();
};
