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

// UpdateHelpOptionValidate
//
// Update command help entry
func UpdateHelpRun(updateArgs []string) {
	if len(updateArgs) < 2 {
		var printText = `error: unknown option
` + UpdateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	//remove self command
	var options []string = updateArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		UpdateHelpUsage()
	}
	var argOptionsMap, err = util.StringArrayToMap(options)
	if err != nil {
		var printText = `error: ` + fmt.Sprint(err) + `
` + UpdateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	if argOptionsMap == nil || len(argOptionsMap) <= 0 {
		var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
` + UpdateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	var option, errvalidate = UpdateHelpOptionValidate(argOptionsMap)
	if errvalidate != nil {
		util.ExitPrintln(fmt.Sprint(errvalidate))
	}
	cmd.UpdateCmdRun(option)
}

// UpdateHelpOptionValidate check terminal option arguments
//
// Required: "-a"
//
// Option:  "-p", "-u", "-x", "-k", "-t", "-g"
func UpdateHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.UPDATE_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + UpdateHelpUsagePrintText())
	}
	var optionMaps, err = entity.ParseArgOptions(argOptionsMap)
	if err != nil {
		return entity.Option{}, err
	}

	//-a
	if _, err = entity.GetOptionMapsByAddress(optionMaps[sys.OPTION_S_ADDRESS]); err != nil {
		return entity.Option{}, err
	}
	//-p
	var mPort, okPort = optionMaps[sys.OPTION_S_PORT]
	if okPort {
		if _, err = entity.GetOptionMapsByPort(mPort); err != nil {
			return entity.Option{}, err
		}
	}
	//-u
	var mUsername, okUsername = optionMaps[sys.OPTION_S_USERNAME]
	if okUsername {
		if _, err = entity.GetOptionMapsByUsername(mUsername); err != nil {
			return entity.Option{}, err
		}
	}
	//[-x | -k]
	var mPassword, okPassword = optionMaps[sys.OPTION_S_PASSWORD]
	var mKey, okKey = optionMaps[sys.OPTION_S_KEY]
	if okPassword {
		if _, err = entity.GetOptionMapsByPassword(mPassword); err != nil {
			return entity.Option{}, err
		}
	}
	if okKey {
		if _, err = entity.GetOptionMapsByKey(mKey); err != nil {
			return entity.Option{}, err
		}
	}
	//-t
	var mTitle, okTitle = optionMaps[sys.OPTION_S_TITLE]
	if okTitle {
		if _, err = entity.GetOptionMapsByTitle(mTitle); err != nil {
			return entity.Option{}, err
		}
	}
	//-g
	var mGroup, okGroup = optionMaps[sys.OPTION_S_GROUP]
	if okGroup {
		if _, err = entity.GetOptionMapsByGroup(mGroup); err != nil {
			return entity.Option{}, err
		}
	}

	return UpdateHelpOptionMapsToOption(optionMaps)
}

// UpdateHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: "-a", "-p", "-u", "-x", "-k", "-t", "-g"
func UpdateHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
	var option = entity.Option{}
	option.SetAddress(optionMaps[sys.OPTION_S_ADDRESS])
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
	var mPassword, okPassword = optionMaps[sys.OPTION_S_PASSWORD]
	var mKey, okKey = optionMaps[sys.OPTION_S_KEY]
	if okPassword {
		var password, _ = entity.GetOptionMapsByPassword(mPassword)
		option.SetPassword(password)
	}
	if okKey {
		var key, _ = entity.GetOptionMapsByKey(mKey)
		option.SetKey(key)
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
	var _, okQuiet = optionMaps[sys.OPTION_S_QUIET]
	if okQuiet {
		option.SetQuiet(okQuiet)
	}
	return option, nil
}

// UpdateHelpUsagePrintText general help text
func UpdateHelpUsagePrintText() string {
	return `Run 'kush update --help' for usage.`
}

// UpdateHelpUsage print update command help usage
func UpdateHelpUsage() {
	var printText = `kush update -a 192.168.1.1 -x 123456 
kush update -a 192.168.1.1 -x 123456 -q
kush update -a 192.168.1.1 -k ~/.ssh/id_rsa_github.pub
` + UpdateHelpUsagePrintText()
	util.ExitPrintln(printText)
}
