/*******************************************************************************
 * Copyright 2018 Samsung Electronics All Rights Reserved.
 *
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
 *
 *******************************************************************************/

package utils

import (
	aml "go/aml"

	"time"
)

var FilePath = "sample_data_model.aml"
var TestKey = "testKey"
var TestValue = "testValue"
var TestDeviceID = "testDevice"
var TestTimeStamp = "2018-04-06:02:12:34"
var TestID = "testID"
var ErrorCode aml.AMLErrorCode
var ValueType aml.AMLValueType

func GetAMLObject() *aml.AMLObject {

	// create "Model" data
	model, _ := aml.CreateAMLData()
	model.SetValueStr("a", "Model_107.113.97.248")
	model.SetValueStr("b", "SR-P7-970")

	// create "Sample" data
	axis, _ := aml.CreateAMLData()
	axis.SetValueStr("x", "20")
	axis.SetValueStr("y", "110")
	axis.SetValueStr("z", "80")

	info, _ := aml.CreateAMLData()
	info.SetValueStr("id", "f437da3b")
	info.SetValueAMLData("axis", axis)

	sample, _ := aml.CreateAMLData()
	sample.SetValueAMLData("info", info)
	appendix := [3]string{"935", "52303", "1442"}
	sample.SetValueStrArr("appendix", appendix[:])

	// set data to object
	amlObj, _ := aml.CreateAMLObject("SAMPLE001", time.Now().Format("20060102150405"))
	amlObj.AddData("Model", model)
	amlObj.AddData("Sample", sample)

	return amlObj
}
