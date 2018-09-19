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

package main

import (
	aml "go/aml"

	"fmt"
	"time"
)

func printAMLData(amlData *aml.AMLData, depth int) {
	var indent string
	for i := 0; i < depth; i++ {
		indent = indent + "   "
	}
	fmt.Printf("%s{\n", indent)

	keys, _ := amlData.GetKeys()
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s    \"%s\" : ", indent, keys[i])
		valType, _ := amlData.GetValueType(keys[i])
		if aml.AMLVALTYPE_STRING == valType {
			value, _ := amlData.GetValueStr(keys[i])
			fmt.Printf("%s", value)
		} else if aml.AMLVALTYPE_STRINGARRAY == valType {
			values, _ := amlData.GetValueStrArr(keys[i])
			fmt.Printf("[")
			for j := 0; j < len(values); j++ {
				fmt.Printf("%s", values[j])
				if j != len(values)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf("]")
		} else if aml.AMLVALTYPE_AMLDATA == valType {
			amlData, _ := amlData.GetValueAMLData(keys[i])
			fmt.Printf("\n")
			printAMLData(amlData, depth+1)
		}
		if i != (len(keys) - 1) {
			fmt.Printf(",")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%s}", indent)
}

func printAMLObject(amlObject *aml.AMLObject) {
	fmt.Printf("{\n")
	deviceId, _ := amlObject.GetDeviceId()
	fmt.Printf("   \"device\" : %s,\n", deviceId)
	timeStamp, _ := amlObject.GetTimeStamp()
	fmt.Printf("   \"timeStamp\" : %s,\n", timeStamp)
	id, _ := amlObject.GetId()
	fmt.Printf("   \"id\" : %s,\n\n", id)

	dataNames, _ := amlObject.GetDataNames()
	for i := 0; i < len(dataNames); i++ {
		data, _ := amlObject.GetData(dataNames[i])
		fmt.Printf("    \"%s\" : \n", dataNames[i])
		printAMLData(data, 1)

		if i != (len(dataNames))-1 {
			fmt.Printf(",\n")
		}
	}
}

func main() {
	repObject, _ := aml.CreateRepresentation("sample_data_model.aml")
	repId, _ := repObject.GetRepresentationId()
	fmt.Printf("\nRepresentation Id : %s\n\n", repId)
	amlObject, _ := repObject.GetConfigInfo()
	printAMLObject(amlObject)
	fmt.Printf("\n}")
	fmt.Printf("\n\nDestroy AMLObject [Result]: %d\n", amlObject.DestroyAMLObject())
	fmt.Printf("\n-------------------------------------------------------------\n")

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
	amlObj, _ := aml.CreateAMLObject("Robot0001", time.Now().Format("20060102150405"))
	amlObj.AddData("Model", model)
	amlObj.AddData("Sample", sample)

	// print object
	printAMLObject(amlObj)
	fmt.Printf("\n}")
	fmt.Printf("\n\n-------------------------------------------------------------\n\n")

	//convert data to aml string
	fmt.Println("DataToAml:")
	amlStr, _ := repObject.DataToAml(amlObj)
	fmt.Printf("%s", amlStr)
	fmt.Printf("\n\n-------------------------------------------------------------\n\n")

	//convert to aml string
	fmt.Println("AmlToData:")
	objectFromStr, _ := repObject.AmlToData(amlStr)
	printAMLObject(objectFromStr)
	fmt.Printf("\n}")
	fmt.Printf("\n\n-------------------------------------------------------------\n\n")

	//convert data to byte
	byteArray, _ := repObject.DataToByte(amlObj)
	fmt.Println("DataToByte done")
	fmt.Printf("\n-------------------------------------------------------------\n\n")

	fmt.Println("ByteToData:")
	byteAMLObj, _ := repObject.ByteToData(byteArray)
	printAMLObject(byteAMLObj)
	fmt.Printf("\n}")
	fmt.Printf("\n\n-------------------------------------------------------------\n\n")
}
