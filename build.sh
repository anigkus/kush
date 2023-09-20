#!/usr/bin/env bash
# Copyright 2023 The https://github.com/anigkus Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this V_FILE except in compliance with the License.
# You may obtain a copy of the License at#

#     http://www.apache.org/licenses/LICENSE-2.0#

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#  
#   build package script
#  

#go args
V_GOOS=$1
V_GOARCH=$2
#V_GOOS=${V_GOOS:-darwin}
#V_GOARCH=${V_GOARCH:-amd64}

# build args
V_ARCH="amd64"
V_FILE='kush'

if [ ! -x "$(command -v go)" ] ||  [ ! -x "$(command -v upx)" ] ||  [ ! -x "$(command -v tar)" ]; then
  echo 'Error: [ go | upx | tar ] is not installed.' >&2
  exit 1
fi

if [ ! -z "$V_GOOS" ] || [ ! -z "$V_GOARCH" ]; then 
    if [ "$V_GOOS" == "windows" ]; then
        V_FILE='kush.exe'
    fi
    if [[ "$V_GOARCH" =~ ^[32|386]$ ]]; then
        V_ARCH='32'
    fi
    V_GOARCH=${V_GOARCH:-${V_ARCH}}
    CGO_ENABLED=0 GOOS=$V_GOOS GOARCH=$V_GOARCH go build -ldflags="-s -w" -o build/$V_GOOS/$V_GOARCH/$V_FILE  && cd build/$V_GOOS/$V_GOARCH && tar -czf kush_"$V_GOOS"_"$V_ARCH".tar.gz $V_FILE && rm -f $V_FILE && cd -
else
    OS="`uname`"
    case $OS in
    'Linux')
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/linux/amd64/kush && upx --best --lzma build/linux/amd64/kush && cd build/linux/amd64 && tar -czf kush_linux_amd64.tar.gz kush && rm -f kush && cd -
        ;;
    'FreeBSD')
        echo "Unsupported OS"
        exit 1
        ;;
    'WindowsNT')
        echo "Unsupported OS"
        exit 1
        ;;
    'Darwin') 
        #CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/darwin/amd64/kush && upx --best --lzma build/darwin/amd64/kush && cd build/darwin/amd64 && tar -czf kush_darwin_amd64.tar.gz kush && rm -f kush && cd -
        ;;
    'SunOS')
        echo "Unsupported OS"
        exit 1
        ;;
    'AIX') 
        echo "Unsupported OS"
        exit 1
        ;;
    *) 
        echo "Unsupported OS"
        exit 1
        ;;
    esac
fi


