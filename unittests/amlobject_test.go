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

func TestCreateAMLObject(t *testing.T) {
	amlObject, _ := aml.CreateAMLObject(utils.TestDeviceID, utils.TestTimeStamp)
	if nil == amlObject {
		t.Errorf("Error amlObject is NULL")
	}
}

func TestCreateAMLObjectWithID(t *testing.T) {
	amlObject, _ := aml.CreateAMLObjectWithID(utils.TestDeviceID, utils.TestTimeStamp, utils.TestID)
	if nil == amlObject {
		t.Errorf("Error amlObject is NULL")
	}
}

func TestDestroyAMLObject(t *testing.T) {
	amlObject, _ := aml.CreateAMLObject(utils.TestDeviceID, utils.TestTimeStamp)
	utils.ErrorCode = amlObject.DestroyAMLObject()
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestAddData(t *testing.T) {
	amlObject, _ := aml.CreateAMLObject(utils.TestDeviceID, utils.TestTimeStamp)
	model, _ := aml.CreateAMLData()
	model.SetValueStr(utils.TestKey, utils.TestValue)
	utils.ErrorCode = amlObject.AddData("Sample", model)
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestAddNullData(t *testing.T) {
	amlObject, _ := aml.CreateAMLObject(utils.TestDeviceID, utils.TestTimeStamp)
	utils.ErrorCode = amlObject.AddData("Sample", nil)
	if 2 != utils.ErrorCode {
		t.Errorf("Error wrong errorcode")
	}
}

func TestGetData(t *testing.T) {
	amlObject, _ := aml.CreateAMLObject(utils.TestDeviceID, utils.TestTimeStamp)
	model, _ := aml.CreateAMLData()
	model.SetValueStr(utils.TestKey, utils.TestValue)
	amlObject.AddData("Sample", model)
	data, _ := amlObject.GetData("Sample")
	if nil == data {
		t.Errorf("Error data is null")
	}
}

func TestGetDataNames(t *testing.T) {
	amlObject, _ := aml.CreateAMLObjectWithID(utils.TestDeviceID, utils.TestTimeStamp, utils.TestID)
	model, _ := aml.CreateAMLData()
	model.SetValueStr(utils.TestKey, utils.TestValue)
	amlObject.AddData("Sample", model)
	names, _ := amlObject.GetDataNames()
	if "Sample" != names[0] {
		t.Errorf("Error name mismatch")
	}
}

func TestGetDeviceId(t *testing.T) {
	amlObject, _ := aml.CreateAMLObjectWithID(utils.TestDeviceID, utils.TestTimeStamp, utils.TestID)
	value, _ := amlObject.GetDeviceId()
	if utils.TestDeviceID != value {
		t.Errorf("Error device id mismatch")
	}
}

func TestGetTimeStamp(t *testing.T) {
	amlObject, _ := aml.CreateAMLObjectWithID(utils.TestDeviceID, utils.TestTimeStamp, utils.TestID)
	value, _ := amlObject.GetTimeStamp()
	if utils.TestTimeStamp != value {
		t.Errorf("Error timeStamp mismatch")
	}
}

func TestGetId(t *testing.T) {
	amlObject, _ := aml.CreateAMLObjectWithID(utils.TestDeviceID, utils.TestTimeStamp, utils.TestID)
	value, _ := amlObject.GetId()
	if "" == value {
		t.Errorf("Error ID mismatch")
	}
}
