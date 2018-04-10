# DataModel AML library (go)

datamodel-aml-go is a go binding written over [datamodel-aml-c](https://github.sec.samsung.net/RS7-EdgeComputing/datamodel-aml-c), that provides the way to present raw data(key/value based) to AutomationML(AML) standard format.
 - Transform raw data to AML data(XML).
 - Serialization / Deserialization AML data using protobuf.
 
 ## Prerequisites ##
 
- Go compiler
  - Version : 1.9
  - [How to install](https://golang.org/doc/install)
  
- datamodel-aml-c
  - Since datamodel-aml-c will be downloaded and built when datamodel-aml-go is built, check the prerequisites of it.
  - [Prerequisites](https://github.sec.samsung.net/RS7-EdgeComputing/datamodel-aml-c)

## How to build ##
1. Goto: ~/datamodel-aml-go/
2. Run the script:

   ```
   ./build.sh         : Native build for x86_64
   ./build_32.sh      : Native build for x86
   ./build_arm.sh     : Cross compile build for armhf
   ./build_arm64.sh   : Cross compile build for arm64
   ./unittests.sh     : Native unit tests build for x86_64 [It will first build aml library after that unittests]
   ```
**Note:** </br>
For getting help about script option: **$ ./build_common.sh --help** </br>

## How to run ##

### Prerequisites ###
 Built datamodel-aml-go package

### sample ###
1. Goto: ~/${GOPATH}/src/go/samples/
2. export LD_LIBRARY_PATH=../extlibs/

3. Run the sample:
    ```
   ./sample
    ```
## Unit test and code coverage report [x86_64 arch]

### Pre-requisite
Built datamodel-aml-go package.

### Run the unit test cases and generate coverage report 
1. Goto: ~/datamodel-aml-go/

2. Run the script </br>
   ` $ ./build.sh`

3. To open coverage report in web browser: </br>
     (a) Goto: ~/datamodel-aml-go/src/go/unittests

     (b) `$ go tool cover -html=coverage.out`   

**Note:** </br>
For running unit test for other architecture please refer **~/datamodel-aml-go/src/go/unittests/build.sh** </br>
     
## Usage guide for aml library (For micro-services) ##
1. The microservice which wants to use aml GO library has to import aml package:
    `import go/aml`
2. Reference aml library APIs : [doc/godoc/aml.html](doc/godoc/aml.html)
