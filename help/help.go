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
  -A, --address         [required] Address 
  -P, --port            [optional] Port, DEFAULT:22
  -U, --username        [optional] Username, DEFAULT:root
  -X, --password        [required] Password, Can only choose one( -X | -K )
  -K, --key             [required] Key, Can only choose one( -X | -K )
  -T, --title           [optional] Title
  -F, --filter          [optional] Filter( ADDRESS , USER , PORT , AUTHTYPE , AUTH , GROUP , TITLE )
  -Q, --quiet           [optional] Quiet
  -O, --output          [optional] Output, DEFAULT:json
  -C, --columns         [optional] columns( ADDRESS , USER , PORT ,AUTHTYPE , AUTH , GROUP , TITLE )
  -W, --wide            [optional] Wide
  -G, --group           [optional] Group, DEFAULT:default
  -S, --sort            [optional] Sort[ASC]( ADDRESS , USER , PORT , AUTHTYPE , GROUP ), DEFAULT:ADDRESS
  -V, --version         version
  -H, --help            help

Run 'kush COMMAND --help' for more information on a command.

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
