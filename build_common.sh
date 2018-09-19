###############################################################################
# Copyright 2018 Samsung Electronics All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
###############################################################################

#!/bin/bash
set +e
#Colors
RED="\033[0;31m"
GREEN="\033[0;32m"
BLUE="\033[0;34m"
NO_COLOUR="\033[0m"

PROJECT_ROOT=$(pwd)
export GOPATH=$PWD
DEP_ROOT=$(pwd)/dependencies
AML_TARGET_ARCH="$(uname -m)"
AML_INSTALL_PREREQUISITES=false
AML_BUILD_MODE="release"
AML_LOGGING="off"
AML_DISABLE_PROTOBUF=false

install_dependencies() {
    if [ "debug" = ${AML_BUILD_MODE} ]; then
        AML_LOGGING="on"
    fi

    # clone datamodel-aml-c library
    if [ -d "./dependencies/datamodel-aml-c" ] ; then
        echo "dependencies/datamodel-aml-c folder exist"
    else
        mkdir dependencies
        cd ./dependencies
        git clone git@github.sec.samsung.net:RS7-EdgeComputing/datamodel-aml-c.git
    fi

    # Build datamodel-aml-c library
    cd $PROJECT_ROOT/dependencies/datamodel-aml-c
	 git fetch origin
	 git checkout v1.0_rel
    echo -e "${GREEN}Building datamodel-aml-c library and its dependencies${NO_COLOUR}"
    ./build_common.sh --target_arch=${AML_TARGET_ARCH} --install_prerequisites=${AML_INSTALL_PREREQUISITES} --build_mode=${AML_BUILD_MODE} --logging=${AML_LOGGING} --disable_protobuf=${AML_DISABLE_PROTOBUF}
    echo -e "${GREEN}Install dependencies done${NO_COLOUR}"
}

build_x86_and_64() {
    cd $PROJECT_ROOT/src/go/
    #build aml SDK
    cd ./aml
    go build
    go install
    #build samples
    cd ../samples
    go build -a sample.go
}

build_arm64() {
    cd $PROJECT_ROOT/src/go/
    #build aml SDK
    cd ./aml
    CGO_ENABLED=1 CC=/usr/bin/aarch64-linux-gnu-gcc-4.8 CXX=/usr/bin/aarch64-linux-gnu-g++-4.8 GOOS=linux GOARCH=arm64 go build
    CGO_ENABLED=1 CC=/usr/bin/aarch64-linux-gnu-gcc-4.8 CXX=/usr/bin/aarch64-linux-gnu-g++-4.8 GOOS=linux GOARCH=arm64 go install
    #build samples
    cd ../samples
    CGO_ENABLED=1 CC=/usr/bin/aarch64-linux-gnu-gcc-4.8 CXX=/usr/bin/aarch64-linux-gnu-g++-4.8 GOOS=linux GOARCH=arm64 go build -a sample.go
}

build_armhf() {
    cd $PROJECT_ROOT/src/go/
    #build aml SDK
    cd ./aml
    CGO_LDFLAGS+='-Bstatic -Bdynamic -lstdc++ -lm' GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc-4.8 CXX=arm-linux-gnueabihf-g++-4.8 go build
    CGO_LDFLAGS+='-Bstatic -Bdynamic -lstdc++ -lm' GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc-4.8 CXX=arm-linux-gnueabihf-g++-4.8 go install

    #build samples
    cd ../samples
    CGO_LDFLAGS+='-Bstatic -Bdynamic -lstdc++ -lm' GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc-4.8 CXX=arm-linux-gnueabihf-g++-4.8 go build -a sample.go
}

build_armhf_native() {
    cd $PROJECT_ROOT/src/go/
    #build aml SDK
    cd ./aml
    CGO_ENABLED=1 GOOS=linux GOARCH=arm go build
    CGO_ENABLED=1 GOOS=linux GOARCH=arm go install
    #build samples
    cd ../samples
    CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -a sample.go
}

clean_aml() {
    echo -e "Cleaning ${BLUE}${PROJECT_ROOT}${NO_COLOUR}"
    echo -e "Deleting  ${RED}${PROJECT_ROOT}/src/${NO_COLOUR}"
    rm -rf ./src
    echo -e "Deleting  ${RED}${PROJECT_ROOT}/pkg/${NO_COLOUR}"
    rm -rf ./dependencies/datamodel-aml-c
    echo -e "Finished Cleaning ${BLUE}${AML}${NO_COLOUR}"
}

usage() {
    echo -e "${BLUE}Usage:${NO_COLOUR} ./build_common.sh <option>"
    echo -e "${GREEN}Options:${NO_COLOUR}"
    echo "  --target_arch=[x86|x86_64|armhf|arm64|armhf-native]         :  Choose Target Architecture"
    echo "  --build_mode=[release|debug](default: release)              :  Build in release or debug mode"
    echo "  --disable_protobuf=[true|false](default: false)             :  Disable protobuf feature"
    echo "  --install_prerequisites=[true|false](default: false)        :  Install the prerequisite S/W to build aml [Protocol-buffer]"
    echo "  -c                                                          :  Clean aml Repository and its dependencies"
    echo "  -h / --help                                                 :  Display help and exit [Be careful it will also remove GOPATH:src, pkg and bin]"
    echo -e "${GREEN}Examples: ${NO_COLOUR}"
    echo -e "${BLUE}  build:-${NO_COLOUR}"
    echo "  $ ./build_common.sh --target_arch=x86_64"
    echo "  $ ./build_common.sh --install_prerequisites=true --target_arch=x86_64 "
    echo -e "${BLUE}  debug mode build:-${NO_COLOUR}"
    echo "  $ ./build_common.sh --target_arch=x86_64 --build_mode=debug"
    echo -e "${BLUE}  clean:-${NO_COLOUR}"
    echo "  $ ./build_common.sh -c"
    echo -e "${BLUE}  help:-${NO_COLOUR}"
    echo "  $ ./build_common.sh -h"
}

build_aml() {
    cd $PROJECT_ROOT
    if [ ! -d "./src/go" ] ; then
        mkdir src
        cd src
        mkdir go
    fi
    cd $PROJECT_ROOT
    #copy aml SDK files
    cp -r aml ./src/go
    #copy aml samples
    cp -r samples ./src/go
    # Copy unit test cases
    cp -r unittests ./src/go
    #copy the datamodel-aml-c/cpp libraries
    if [ ! -d "./src/go/extlibs" ] ; then
        cd src/go
        mkdir extlibs
    fi
	
    TARGET_ARCH=${AML_TARGET_ARCH}
    if [ "armhf-native" = ${TARGET_ARCH} ]; then
         TARGET_ARCH="armhf";
    fi
	
    cd $PROJECT_ROOT
    cp -r dependencies/datamodel-aml-c/dependencies/datamodel-aml-cpp/out/linux/${TARGET_ARCH}/${AML_BUILD_MODE}/lib* ./src/go/extlibs
    cp -r dependencies/datamodel-aml-c/out/linux/${TARGET_ARCH}/${AML_BUILD_MODE}/lib* ./src/go/extlibs

    export CGO_CFLAGS=-I$PWD/dependencies/datamodel-aml-c/include
    export CGO_LDFLAGS=-L$PWD/src/go/extlibs
    export CGO_LDFLAGS=$CGO_LDFLAGS" -lcaml -laml"

    if [ "x86" = ${AML_TARGET_ARCH} ]; then
         build_x86_and_64;
    elif [ "x86_64" = ${AML_TARGET_ARCH} ]; then
         build_x86_and_64;
    elif [ "arm64" = ${AML_TARGET_ARCH} ]; then
         build_arm64;
    elif [ "armhf" = ${AML_TARGET_ARCH} ]; then
         build_armhf;
    elif [ "armhf-native" = ${AML_TARGET_ARCH} ]; then
         build_armhf_native;
    else
         echo -e "${RED}Not a supported architecture${NO_COLOUR}"
         usage; exit 1;
    fi
}

process_cmd_args() {
    if [ "$#" -eq 0  ]; then
        echo -e "No argument.."
        usage; exit 1
    fi

    while [ "$#" -gt 0  ]; do
        case "$1" in
            --install_prerequisites=*)
                AML_INSTALL_PREREQUISITES="${1#*=}";
                if [ ${AML_INSTALL_PREREQUISITES} != true ] && [ ${AML_INSTALL_PREREQUISITES} != false ]; then
                    echo -e "${RED}Unknown option for --install_prerequisites${NO_COLOUR}"
                    shift 1; exit 0
                fi
                echo -e "${GREEN}Install the prerequisites before build: ${AML_INSTALL_PREREQUISITES}${NO_COLOUR}"
                shift 1;
                ;;
            --target_arch=*)
                AML_TARGET_ARCH="${1#*=}";
                echo -e "${GREEN}Target Arch is: $AML_TARGET_ARCH${NO_COLOUR}"
                shift 1
                ;;
            --build_mode=*)
                AML_BUILD_MODE="${1#*=}";
                echo -e "${GREEN}Build mode is: $AML_BUILD_MODE${NO_COLOUR}"
                shift 1;
                ;;
            --disable_protobuf=*)
                AML_DISABLE_PROTOBUF="${1#*=}";
                if [ ${AML_DISABLE_PROTOBUF} != true ] && [ ${AML_DISABLE_PROTOBUF} != false ]; then
                    echo -e "${RED}Unknown option for --disable_protobuf${NO_COLOUR}"
                    shift 1; exit 0
                fi
                echo -e "${GREEN}is Protobuf disabled : $AML_DISABLE_PROTOBUF${NO_COLOUR}"
                shift 1;
                ;;
            -c)
                clean_aml
                shift 1; exit 0
                ;;
            -h)
                usage; exit 0
                ;;
            --help)
                usage; exit 0
                ;;
            -*)
                echo -e "${RED}"
                echo "unknown option: $1" >&2;
                echo -e "${NO_COLOUR}"
                usage; exit 1
                ;;
             *)
                echo -e "${RED}"
                echo "unknown option: $1" >&2;
                echo -e "${NO_COLOUR}"
                usage; exit 1
                ;;
        esac
    done
}

process_cmd_args "$@"
install_dependencies
build_aml
echo -e "${GREEN}Build done${NO_COLOUR}"

