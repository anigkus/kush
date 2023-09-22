// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"os"

	"github.com/anigkus/kush/cli"
	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/help"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// entry
func Main(args []string) {
	var size int = len(args)
	if size <= 1 {
		help.HelpUsage(args)
	}
	//remove self filename
	var outArg = args[1]
	switch outArg {
	case "":
		help.HelpUsage(args)
	case sys.OPTION_S_VERSION, sys.OPTION_L_VERSION:
		help.Version()
	case sys.OPTION_S_HELP, sys.OPTION_L_HELP:
		help.HelpUsage(args)
	case sys.COMMANDS_CREATE:
		//remove self filename
		help.CreateHelpRun(args[1:])
	case sys.COMMANDS_REMOVE:
		//remove self filename
		help.RemoveHelpRun(args[1:])
	case sys.COMMANDS_UPDATE:
		//remove self filename
		help.UpdateHelpRun(args[1:])
	case sys.COMMANDS_SEARCH:
		//remove self filename
		help.SearchHelpRun(args[1:])
	case sys.COMMANDS_EXPORT:
		//remove self filename
		help.ExportHelpRun(args[1:])
	case sys.COMMANDS_IMPORT:
		//remove self filename
		help.ImportHelpRun(args[1:])
	default:
		var argOptionsMap, err = util.DirectArrayToMap(args[1:])
		if err == nil {
			var option = DirectHandler(argOptionsMap)
			var terminal = new(cli.Terminal)
			if len(option.Password) > 0 {
				terminal = cli.New(option.Address, option.Username, option.Password, "", option.Port)
			} else if len(option.Key) > 0 {
				terminal = cli.New(option.Address, option.Username, "", option.Key, option.Port)
			}
			err := terminal.RunTerminal(os.Stdout, os.Stdin)
			if err != nil {
				util.ExitPrintln(fmt.Sprintf("error: %v", err))
			}
		} else {
			var printText = `error: unknown command "` + outArg + `" for "kush"
` + help.HelpUsagePrintText()
			if !help.ValidataOption(args) {
				printText = `error: unknown option "` + outArg + `" for "kush"
` + help.HelpUsagePrintText()
			}
			util.ExitPrintln(printText)
		}
	}
}

// DirectHandler
func DirectHandler(argOptionsMap map[string]string) entity.Option {
	var option = entity.Option{}
	var mAdddress = argOptionsMap[sys.OPTION_S_ADDRESS]
	var address, err = entity.GetOptionMapsByAddress(mAdddress)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	} else {
		option.SetAddress(address)
	}
	var mPort, okPort = argOptionsMap[sys.OPTION_S_PORT]
	if okPort {
		var port, err = entity.GetOptionMapsByPort(mPort)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
		option.SetPort(port)
	} else {
		option.SetPort(sys.DEFAULT_PORT)
	}
	var mUsername, okUsername = argOptionsMap[sys.OPTION_S_USERNAME]
	if okUsername {
		var username, err = entity.GetOptionMapsByUsername(mUsername)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
		option.SetUsername(username)
	}
	var mPassword, okPassword = argOptionsMap[sys.OPTION_S_PASSWORD]
	var mKey, okKey = argOptionsMap[sys.OPTION_S_KEY]
	if !okPassword && !okKey {
		util.ExitPrintln("error: required choose one( -x | -k )")
	}
	if okPassword && okKey {
		util.ExitPrintln("error: can only choose one( -x | -k )")
	}
	if okPassword {
		var password, err = entity.GetOptionMapsByPassword(mPassword)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
		option.SetPassword(password)
	}
	if okKey {
		var key, err = entity.GetOptionMapsByKey(mKey)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
		option.SetKey(key)
	}
	return option
}
