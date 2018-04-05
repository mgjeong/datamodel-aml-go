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

type AMLErrorCode int

// Constants represents AML error codes.
const (
	AML_OK                 = 0
	AML_INVALID_PARAM      = 1
	AML_INVALID_DATA       = 2
	AML_INVALID_FILE_PATH  = 3
	AML_INVALID_AML_SCHEMA = 4
	AML_INVALID_XML_STRING = 5
	AML_NO_MEMORY          = 6
	AML_KEY_NOT_EXIST      = 7
	AML_KEY_ALREADY_EXIST  = 8
	AML_INVALID_DATA_TYPE  = 9
)
