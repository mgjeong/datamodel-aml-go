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

func TestCreateRepresentation(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	if nil == repObject {
		t.Errorf("Error repObject is NULL")
	}
}

func TestCreateRepresentationN(t *testing.T) {
	_, errorCode := aml.CreateRepresentation("")
	if aml.AML_OK == errorCode {
		t.Errorf("Error worng errorcode")
	}
}

func TestDestroyRepresentation(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	utils.ErrorCode = repObject.DestroyRepresentation()
	if 0 != utils.ErrorCode {
		t.Errorf("Error errorcode is not AML_OK")
	}
}

func TestGetRepresentationId(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	id, _ := repObject.GetRepresentationId()
	if len(id) <= 0 {
		t.Errorf("Error id is empty")
	}
}

func TestGetConfigInfo(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObject, _ := repObject.GetConfigInfo()
	if nil == amlObject {
		t.Errorf("Error amlObject is NULL")
	}
}

func TestDataToAml(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObject := utils.GetAMLObject()
	amlString, _ := repObject.DataToAml(amlObject)
	if "" == amlString {
		t.Errorf("Error String is empty")
	}
}

func TestAmlToData(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObject := utils.GetAMLObject()
	amlStr, _ := repObject.DataToAml(amlObject)
	objectFromStr, _ := repObject.AmlToData(amlStr)
	if nil == objectFromStr {
		t.Errorf("Error objectFromStr is NULL")
	}
}

func TestDataToByte(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObject := utils.GetAMLObject()
	byteData, errorCode := repObject.DataToByte(amlObject)
	if nil == byteData || errorCode != aml.AML_OK {
		t.Errorf("Error byteData is NULL")
	}
}

func TestByteToData(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObject := utils.GetAMLObject()
	byteData, _ := repObject.DataToByte(amlObject)
	amlObjectFromByte, errorCode := repObject.ByteToData(byteData)
	if nil == amlObjectFromByte || errorCode != aml.AML_OK {
		t.Errorf("Error amlObjectFromByte is NULL")
	}
}

func TestNegativeDataToByte(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	byteData, errorCode := repObject.DataToByte(nil)
	if nil != byteData || errorCode != aml.AML_INVALID_PARAM {
		t.Errorf("Failed")
	}
}

func TestNegativeByteToData(t *testing.T) {
	repObject, _ := aml.CreateRepresentation(utils.FilePath)
	amlObjectFromByte, errorCode := repObject.ByteToData(nil)
	if nil != amlObjectFromByte || errorCode != aml.AML_INVALID_PARAM {
		t.Errorf("Failed")
	}
}
