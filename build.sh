#!/usr/bin/env bash
# Copyright 2023 The https://github.com/anigkus Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
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
goos=$1
goarch=$2

# build args
arch="64"
file='kush'

if [ ! -z "$goos" ] && [ ! -z "$goarch" ]; then
    if [ "$goos" == "windows" ]; then
        file='kush.exe'
    fi
    if [[ "$goarch" =~ ^[32|386]$ ]]; then
        arch='32'
    fi
    CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -ldflags="-s -w" -o build/$goos/$goarch/$file  && upx --best --lzma build/$goos/$goarch/$file && cd build/$goos/$goarch && tar -czf kush_$goos_x$arch.tar.gz $file && rm -f $file && cd -
else
    OS="`uname`"
    case $OS in
    'Linux')
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/linux/amd64/kush && upx --best --lzma build/linux/amd64/kush && cd build/linux/amd64 && tar -czf kush_linux_x64.tar.gz kush && rm -f kush && cd -
        ;;
    'FreeBSD')
        echo "Unsupported OS"
        ;;
    'WindowsNT')
        echo "Unsupported OS"
        ;;
    'Darwin') 
        CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/darwin/amd64/kush && upx --best --lzma build/darwin/amd64/kush && cd build/darwin/amd64 && tar -czf kush_darwin_x64.tar.gz kush && rm -f kush && cd -
        ;;
    'SunOS')
      echo "Unsupported OS"
        ;;
    'AIX') 
        echo "Unsupported OS"
        ;;
    *) 
        echo "Unsupported OS"
        ;;
    esac
fi


