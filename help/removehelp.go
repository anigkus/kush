// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package help

import (
	"errors"
	"fmt"

	"github.com/anigkus/kush/cmd"
	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// RemoveHelpRun
//
// remove command help entry
func RemoveHelpRun(removeArgs []string) {
	//remove self command
	var options []string = removeArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		RemoveHelpUsage()
	}
	var option = entity.Option{}
	if len(removeArgs) > 1 {
		var argOptionsMap, err = util.StringArrayToMap(options)
		if err != nil {
			var printText = `error: ` + fmt.Sprint(err) + `
	` + RemoveHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if argOptionsMap == nil || len(argOptionsMap) <= 0 {
			var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
	` + RemoveHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if option, err = RemoveHelpOptionValidate(argOptionsMap); err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
	cmd.RemoveCmdRun(option)
}

// RemoveHelpOptionValidate check terminal option arguments
//
// Option: "-A", "-P", "-U", "-T", "-G", "-F"
func RemoveHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.REMOVE_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + RemoveHelpUsagePrintText())
	}
	var optionMaps, err = entity.ParseArgOptions(argOptionsMap)
	if err != nil {
		return entity.Option{}, err
	}

	//-A
	var mAdderess, okAdderess = optionMaps[sys.OPTION_S_ADDRESS]
	if okAdderess {
		if _, err = entity.GetOptionMapsByAddress(mAdderess); err != nil {
			return entity.Option{}, err
		}
	}
	//-P
	var mPort, okPort = optionMaps[sys.OPTION_S_PORT]
	if okPort {
		if _, err = entity.GetOptionMapsByPort(mPort); err != nil {
			return entity.Option{}, err
		}
	}
	//-U
	var mUsername, okUsername = optionMaps[sys.OPTION_S_USERNAME]
	if okUsername {
		if _, err = entity.GetOptionMapsByUsername(mUsername); err != nil {
			return entity.Option{}, err
		}
	}
	//-T
	var mTitle, okTitle = optionMaps[sys.OPTION_S_TITLE]
	if okTitle {
		if _, err = entity.GetOptionMapsByTitle(mTitle); err != nil {
			return entity.Option{}, err
		}
	}
	//-G
	var mGroup, okGroup = optionMaps[sys.OPTION_S_GROUP]
	if okGroup {
		if _, err = entity.GetOptionMapsByGroup(mGroup); err != nil {
			return entity.Option{}, err
		}
	}
	//-F
	var mFilter, okFilter = optionMaps[sys.OPTION_S_FILTER]
	if okFilter {
		if _, err = entity.GetOptionMapsByFilter(mFilter); err != nil {
			return entity.Option{}, err
		}
	}
	return RemoveHelpOptionMapsToOption(optionMaps)
}

// RemoveHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: "-A", "-P", "-U", "-T", "-G", "-F", "-Q"
func RemoveHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
	var option = entity.Option{}

	var mAdderess, okAdderess = optionMaps[sys.OPTION_S_ADDRESS]
	if okAdderess {
		var address, _ = entity.GetOptionMapsByAddress(mAdderess)
		option.SetAddress(address)
	}
	var mPort, okPort = optionMaps[sys.OPTION_S_PORT]
	if okPort {
		var port, _ = entity.GetOptionMapsByPort(mPort)
		option.SetPort(port)
	}
	var mUsername, okUsername = optionMaps[sys.OPTION_S_USERNAME]
	if okUsername {
		var username, _ = entity.GetOptionMapsByUsername(mUsername)
		option.SetUsername(username)
	}
	var mTitle, okTitle = optionMaps[sys.OPTION_S_TITLE]
	if okTitle {
		var title, _ = entity.GetOptionMapsByTitle(mTitle)
		option.SetTitle(title)
	}
	var mGroup, okGroup = optionMaps[sys.OPTION_S_GROUP]
	if okGroup {
		var group, _ = entity.GetOptionMapsByGroup(mGroup)
		option.SetGroup(group)
	}
	var mFilter, okFilter = optionMaps[sys.OPTION_S_FILTER]
	if okFilter {
		var filtermap, _ = entity.GetOptionMapsByFilter(mFilter)
		option.SetFilter(filtermap)
	}
	var _, okQuiet = optionMaps[sys.OPTION_S_QUIET]
	if okQuiet {
		option.SetQuiet(okQuiet)
	}
	return option, nil
}

// RemoveHelpUsagePrintText general help text
func RemoveHelpUsagePrintText() string {
	return `Run 'kush remove --help' for usage.`
}

// RemoveHelpUsage print Remove command help usage
func RemoveHelpUsage() {
	var printText = `kush remove -A 192.168.1.1 -U root -P 22 
kush remove -A 192.168.1.1 -U root -P 22 -Q
kush remove -A 192.168.1.1 -F 'GROUP=group1' -Q
` + RemoveHelpUsagePrintText()
	util.ExitPrintln(printText)
}
