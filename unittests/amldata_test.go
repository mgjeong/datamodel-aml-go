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

package unittests

import (
	aml "go/aml"

	utils "go/unittests/utils"

	"testing"
)

func TestCreateAMLData(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	if nil == amlData {
		t.Errorf("Error amlObject is NULL")
	}
}

func TestDestroyAMLData(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	utils.ErrorCode = amlData.DestroyAMLData()
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestSetValueStr(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	utils.ErrorCode = amlData.SetValueStr(utils.TestKey, utils.TestValue)
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestSetValueStrArr(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	stringArray := [3]string{"935", "52303", "1442"}
	utils.ErrorCode = amlData.SetValueStrArr(utils.TestKey, stringArray[:])
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestSetValueAMLData(t *testing.T) {
	amlDataObject, _ := aml.CreateAMLData()
	amlData, _ := aml.CreateAMLData()
	amlData.SetValueStr("x", "20")
	utils.ErrorCode = amlDataObject.SetValueAMLData(utils.TestKey, amlData)
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestSetValueNullAMLData(t *testing.T) {
	amlDataObject, _ := aml.CreateAMLData()
	utils.ErrorCode = amlDataObject.SetValueAMLData(utils.TestKey, nil)
	if 2 != utils.ErrorCode {
		t.Errorf("Error wrong errorcode")
	}
}

func TestGetValueStr(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	amlData.SetValueStr(utils.TestKey, utils.TestValue)
	value, _ := amlData.GetValueStr(utils.TestKey)
	if value != utils.TestValue {
		t.Errorf("Error value mismatch")
	}
}

func TestGetValueStrArr(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	stringArray := [2]string{"935", "52303"}
	utils.ErrorCode = amlData.SetValueStrArr(utils.TestKey, stringArray[:])
	value, _ := amlData.GetValueStrArr(utils.TestKey)
	if value[0] != stringArray[0] || value[1] != stringArray[1] {
		t.Errorf("Error value mismatch")
	}
}

func TestGetValueAMLData(t *testing.T) {
	amlDataObject, _ := aml.CreateAMLData()
	amlData, _ := aml.CreateAMLData()
	amlData.SetValueStr("x", "20")
	utils.ErrorCode = amlDataObject.SetValueAMLData(utils.TestKey, amlData)
	data, _ := amlDataObject.GetValueAMLData(utils.TestKey)
	value, _ := data.GetValueStr("x")
	if "20" != value {
		t.Errorf("Error value mismatch")
	}
}

func TestGetKeys(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	amlData.SetValueStr(utils.TestKey, utils.TestValue)
	keys, _ := amlData.GetKeys()
	if keys[0] != utils.TestKey {
		t.Errorf("Error key mismatch")
	}
}

func TestGetValueType(t *testing.T) {
	amlData, _ := aml.CreateAMLData()
	amlData.SetValueStr(utils.TestKey, utils.TestValue)
	keys, _ := amlData.GetKeys()
	utils.ValueType, _ = amlData.GetValueType(keys[0])
	if 0 != utils.ValueType {
		t.Errorf("Error value type mismatch")
	}
}
