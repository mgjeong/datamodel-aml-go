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

package aml

// #include "camlinterface.h"
import "C"
import "unsafe"

// Structure represents AMLObject.
type AMLObject struct {
	amlObject C.amlObjectHandle_t
}

// Create an instance of AMLObject.
func CreateAMLObject(deviceId string, timeStamp string) (*AMLObject, AMLErrorCode) {
	instance := &AMLObject{}
	result := C.CreateAMLObject(C.CString(deviceId), C.CString(timeStamp), &(instance.amlObject))
	return instance, AMLErrorCode(int(result))
}

// Create an instance of AMLObject with id.
func CreateAMLObjectWithID(deviceId string, timeStamp string, id string) (*AMLObject, AMLErrorCode) {
	instance := &AMLObject{}
	result := C.CreateAMLObjectWithID(C.CString(deviceId), C.CString(timeStamp), C.CString(id), &(instance.amlObject))
	return instance, AMLErrorCode(int(result))
}

// Destroy an instance of AMLObject.
func (amlInstance *AMLObject) DestroyAMLObject() AMLErrorCode {
	return AMLErrorCode(int(C.DestroyAMLObject(amlInstance.amlObject)))
}

// Add AMLData to AMLObject using AMLData key that to match AMLData value.
func (amlInstance *AMLObject) AddData(name string, data *AMLData) AMLErrorCode {
	if nil == data {
		return AML_INVALID_DATA
	}
	return AMLErrorCode(int(C.AMLObject_AddData(amlInstance.amlObject, C.CString(name), data.getDataHandle())))
}

// Get AMLData which matched input name string with AMLObject's amlDatas key.
func (amlInstance *AMLObject) GetData(name string) (*AMLData, AMLErrorCode) {
	instance := &AMLData{}
	result := C.AMLObject_GetData(amlInstance.amlObject, C.CString(name), &(instance.amlData))
	return instance, AMLErrorCode(int(result))
}

// Get a list of AMLData names that AMLObject has.
func (amlInstance *AMLObject) GetDataNames() ([]string, AMLErrorCode) {
	var names **C.char
	var size C.size_t
	result := C.AMLObject_GetDataNames(amlInstance.amlObject, &names, &size)
	if result != AML_OK {
		return nil, AMLErrorCode(int(result))
	}
	nameslice := (*[1 << 28]*C.char)(unsafe.Pointer(names))[:size:size]
	gostrings := make([]string, size)
	for i, s := range nameslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings, AMLErrorCode(int(result))
}

// Get deviceId of AMLObject.
func (amlInstance *AMLObject) GetDeviceId() (string, AMLErrorCode) {
	var id *C.char
	result := C.AMLObject_GetDeviceId(amlInstance.amlObject, &id)
	return C.GoString(id), AMLErrorCode(int(result))
}

// Get timeStamp of AMLObject.
func (amlInstance *AMLObject) GetTimeStamp() (string, AMLErrorCode) {
	var timeStamp *C.char
	result := C.AMLObject_GetTimeStamp(amlInstance.amlObject, &timeStamp)
	return C.GoString(timeStamp), AMLErrorCode(int(result))
}

// Get id of AMLObject.
func (amlInstance *AMLObject) GetId() (string, AMLErrorCode) {
	var id *C.char
	result := C.AMLObject_GetId(amlInstance.amlObject, &id)
	return C.GoString(id), AMLErrorCode(int(result))
}

func (amlInstance *AMLObject) getDataHandle() C.amlObjectHandle_t {
	return amlInstance.amlObject
}
