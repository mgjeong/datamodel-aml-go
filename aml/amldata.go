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

// Structure represents AMLData.
type AMLData struct {
	amlData C.amlDataHandle_t
}

// Create an instance of AMLData.
func CreateAMLData() (*AMLData, AMLErrorCode) {
	instance := &AMLData{}
	result := C.CreateAMLData(&(instance.amlData))
	return instance, AMLErrorCode(int(result))
}

// Destroy an instance of AMLData.
func (dataInstance *AMLData) DestroyAMLData() AMLErrorCode {
	return AMLErrorCode(int(C.DestroyAMLData(dataInstance.amlData)))
}

// Set key/value as a string value to AMLData.
func (dataInstance *AMLData) SetValueStr(key string, value string) AMLErrorCode {
	return AMLErrorCode(int(C.AMLData_SetValueStr(dataInstance.amlData, C.CString(key), C.CString(value))))
}

// Set key/value as a string array value to AMLData.
func (dataInstance *AMLData) SetValueStrArr(key string, values []string) AMLErrorCode {
	nValues := make([](*_Ctype_char), 0)
	for i, _ := range values {
		char := C.CString(values[i])
		defer C.free(unsafe.Pointer(char))
		strptr := (*_Ctype_char)(unsafe.Pointer(char))
		nValues = append(nValues, strptr)
	}
	return AMLErrorCode(int(C.AMLData_SetValueStrArr(dataInstance.amlData, C.CString(key),
		(**_Ctype_char)(unsafe.Pointer(&nValues[0])), C.size_t(len(values)))))
}

// Set key/value as a AMLData value to AMLData.
func (dataInstance *AMLData) SetValueAMLData(key string, value *AMLData) AMLErrorCode {
	if nil == value {
		return AML_INVALID_DATA
	}
	return AMLErrorCode(int(C.AMLData_SetValueAMLData(dataInstance.amlData, C.CString(key), value.getDataHandle())))
}

// Get string value which matchs a key in AMLData.
func (dataInstance *AMLData) GetValueStr(key string) (string, AMLErrorCode) {
	var value *C.char
	result := C.AMLData_GetValueStr(dataInstance.amlData, C.CString(key), &value)
	return C.GoString(value), AMLErrorCode(int(result))
}

// Get string array value which matchs a key in AMLData.
func (dataInstance *AMLData) GetValueStrArr(key string) ([]string, AMLErrorCode) {
	var values **C.char
	var size C.size_t
	result := C.AMLData_GetValueStrArr(dataInstance.amlData, C.CString(key), &values, &size)
	if result != AML_OK {
		return nil, AMLErrorCode(int(result))
	}
	valueslice := (*[1 << 28]*C.char)(unsafe.Pointer(values))[:size:size]
	gostrings := make([]string, size)
	for i, s := range valueslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings, AMLErrorCode(int(result))
}

// Get AMLData value which matchs a key in AMLData.
func (dataInstance *AMLData) GetValueAMLData(key string) (*AMLData, AMLErrorCode) {
	instance := &AMLData{}
	result := C.AMLData_GetValueAMLData(dataInstance.amlData, C.CString(key), &(instance.amlData))
	return instance, AMLErrorCode(int(result))
}

// Get a list of key that AMLData has.
func (dataInstance *AMLData) GetKeys() ([]string, AMLErrorCode) {
	var keys **C.char
	var size C.size_t
	result := C.AMLData_GetKeys(dataInstance.amlData, &keys, &size)
	if result != AML_OK {
		return nil, AMLErrorCode(int(result))
	}
	keyslice := (*[1 << 28]*C.char)(unsafe.Pointer(keys))[:size:size]
	goKeys := make([]string, size)
	for i, key := range keyslice {
		goKeys[i] = C.GoString(key)
	}
	return goKeys, AMLErrorCode(int(result))
}

// Get AML datatype of value for the given key.
func (dataInstance *AMLData) GetValueType(key string) (AMLValueType, AMLErrorCode) {
	var valueType C.CAMLValueType
	result := C.AMLData_GetValueType(dataInstance.amlData, C.CString(key), &valueType)
	return AMLValueType(int(valueType)), AMLErrorCode(int(result))
}

func (dataInstance *AMLData) getDataHandle() C.amlDataHandle_t {
	return dataInstance.amlData
}
