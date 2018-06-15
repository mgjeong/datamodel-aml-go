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

// aml package which provides simplified APIs for aml data parsing.
package aml

// #include "camlrepresentation.h"
import "C"
import "unsafe"

// Structure represents Representation.
type Representation struct {
	representation C.representation_t
}

// Contructs an instance of Representation for the given aml file.
func CreateRepresentation(filePath string) (*Representation, AMLErrorCode) {
	instance := &Representation{}
	result := C.CreateRepresentation(C.CString(filePath), &(instance.representation))
	return instance, AMLErrorCode(int(result))
}

// Destroy an instance of Representation.
func (repInstance *Representation) DestroyRepresentation() AMLErrorCode {
	return AMLErrorCode(int(C.DestroyRepresentation(repInstance.representation)))
}

// Get AutomationML SystemUnitClassLib's unique ID.
func (repInstance *Representation) GetRepresentationId() (string, AMLErrorCode) {
	var id *C.char
	result := C.Representation_GetRepId(repInstance.representation, &id)
	return C.GoString(id), AMLErrorCode(int(result))
}

// Get AMLObject that contains configuration data which is present in RoleClassLib.
func (repInstance *Representation) GetConfigInfo() (*AMLObject, AMLErrorCode) {
	instance := &AMLObject{}
	result := C.Representation_GetConfigInfo(repInstance.representation, &(instance.amlObject))
	return instance, AMLErrorCode(int(result))
}

// Converts AMLObject to AML(XML) string to match the AML model information
// which is set on CreateRepresentation().
func (repInstance *Representation) DataToAml(amlObject *AMLObject) (string, AMLErrorCode) {
	if nil == amlObject {
		return "", AML_INVALID_PARAM
	}
	var value *C.char
	result := C.Representation_DataToAml(repInstance.representation, amlObject.getDataHandle(), &value)
	return C.GoString(value), AMLErrorCode(int(result))
}

// converts AML(XML) string to AMLObject to match the AML model information
// which is set on CreateRepresentation().
func (repInstance *Representation) AmlToData(amlStr string) (*AMLObject, AMLErrorCode) {
	instance := &AMLObject{}
	result := C.Representation_AmlToData(repInstance.representation, C.CString(amlStr), &(instance.amlObject))
	return instance, AMLErrorCode(int(result))
}

// AMLObject to Protobuf byte data to match the AML model information
// which is set on CreateRepresentation().
func (repInstance *Representation) DataToByte(amlObject *AMLObject) ([]byte, AMLErrorCode) {
	if nil == amlObject {
		return nil, AML_INVALID_PARAM
	}
	var data *C.uint8_t
	var size C.size_t
	result := C.Representation_DataToByte(repInstance.representation, amlObject.amlObject, &data, &size)
	byteSlice := (*[1 << 28]byte)(unsafe.Pointer(data))[:size:size]
	return byteSlice, AMLErrorCode(int(result))
}

// Converts Protobuf byte data to AMLObject to match the AML model information
// which is set on CreateRepresentation()
func (repInstance *Representation) ByteToData(data []byte) (*AMLObject, AMLErrorCode) {
	if nil == data {
		return nil, AML_INVALID_PARAM
	}
	instance := &AMLObject{}
	result := C.Representation_ByteToData(repInstance.representation, (*C.uint8_t)(&data[0]), C.size_t(len(data)), &(instance.amlObject))
	return instance, AMLErrorCode(int(result))
}
