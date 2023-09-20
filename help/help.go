// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package help

import (
	"strings"

	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// HelpUsage print help usage
func HelpUsage(args []string) {
	var printText string = `Usage:  kush COMMAND [OPTIONS]

COMMAND:
  create    Create the kush for host
  remove    Remove the kush for host
  update    Update the kush for host
  search    Search the kush for host
  export    Export the kush to  file
  import    Import the kush for file
	
OPTIONS:
  -a, --address         [required] Address 
  -p, --port            [optional] Port, DEFAULT:22
  -u, --username        [optional] Username, DEFAULT:root
  -x, --password        [required] Password, Can only choose one( -x | -k )
  -k, --key             [required] Key, Can only choose one( -x | -k )
  -t, --title           [optional] Title
  -f, --filter          [optional] Filter( ADDRESS , USER , PORT , AUTHTYPE , AUTH , GROUP , TITLE )
  -q, --quiet           [optional] Quiet
  -o, --output          [optional] Output, DEFAULT:json
  -c, --columns         [optional] columns( ADDRESS , USER , PORT ,AUTHTYPE , AUTH , GROUP , TITLE )
  -w, --wide            [optional] Wide
  -g, --group           [optional] Group, DEFAULT:default
  -s, --sort            [optional] Sort[ASC]( ADDRESS , USER , PORT , AUTHTYPE , GROUP ), DEFAULT:ADDRESS
  -v, --version         version
  -h, --help            help

Run 'kush COMMAND --help' for more information on a command.

Direct Connection:

kush user@host -x 123456
kush user@host -k ~/.ssh/id_rsa_github.pub

To get more help with kush, check out our guides at https://github.com/anigkus/kush`
	util.ExitPrintln(printText)
}

// Version
func Version() {
	var printText string = "Version:    " + sys.VERSION
	util.ExitPrintln(printText)
}

// ValidataCmd check terminal input cmd parameters
func ValidataCmd(args []string) bool {
	if len(args) > 1 {
		var command = args[1]
		var index = strings.Index(command, "-")
		if index == -1 {
			if !util.InArray(command, sys.COMMANDS) {
				return false
			}
		}
	}
	return true
}

// ValidataOption check terminal input option parameters
func ValidataOption(args []string) bool {
	if len(args) > 1 {
		var option = args[1]
		var hyphen = strings.Index(option, "-")
		if hyphen != -1 {
			if !util.InArray(option, sys.OPTIONS) {
				return false
			}
		}
	}
	return true
}

// HelpUsagePrintText general help text
func HelpUsagePrintText() string {
	return `Run 'kush --help' for usage.`
}
