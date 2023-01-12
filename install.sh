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
#  bash install script
#  

V_VERSION="$1"
V_GOARCH=$2


if [ -z "$V_VERSION" ]; then
    V_VERSION=`sudo curl https://github.com/anigkus/kush/releases | grep -oiE "<a([^>]+) class=\"Link--primary\">([^<]+)</a>" | awk -F '>' '{print $2}' | awk -F '<' '{print $1}'`
else
    V_VERSION="v$V_VERSION"
fi

OS="`uname`"
case $OS in
'Linux')
    V_FILE='kush_linux_amd64.tar.gz'
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
    V_FILE='kush_darwin_amd64.tar.gz'
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

sudo curl -sLX GET https://github.com/anigkus/kush/releases/download/$V_VERSION/$V_FILE | sudo tar -xz -C /usr/local/bin/ 2>/dev/null