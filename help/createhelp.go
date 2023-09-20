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
	"golang.org/x/term"
)

// CreateHelpOptionValidate
//
// Create command help entry
func CreateHelpRun(createArgs []string) {
	if len(createArgs) < 2 {
		var printText = `error: unknown option
` + CreateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	//remove self command
	var options []string = createArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		CreateHelpUsage()
	}
	var argOptionsMap, err = util.StringArrayToMap(options)
	if err != nil {
		var printText = `error: ` + fmt.Sprint(err) + `
` + CreateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	if argOptionsMap == nil || len(argOptionsMap) <= 0 {
		var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
` + CreateHelpUsagePrintText()
		util.ExitPrintln(printText)
	}
	var option, errvalidate = CreateHelpOptionValidate(argOptionsMap)
	if errvalidate != nil {
		util.ExitPrintln(fmt.Sprint(errvalidate))
	}
	cmd.CreateCmdRun(option)
}

// CreateHelpOptionValidate check terminal option arguments
//
// Required: "-a" , "[-x|-k]
//
// Option: "-p", "-u",  "-t", "-g", "-q"
func CreateHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.CREATE_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + CreateHelpUsagePrintText())
	}
	var optionMaps, err = entity.ParseArgOptions(argOptionsMap)
	if err != nil {
		return entity.Option{}, err
	}

	//-a
	if _, err = entity.GetOptionMapsByAddress(optionMaps[sys.OPTION_S_ADDRESS]); err != nil {
		return entity.Option{}, err
	}

	//[-x | -k]
	var mPassword, okPassword = optionMaps[sys.OPTION_S_PASSWORD]
	var mKey, okKey = optionMaps[sys.OPTION_S_KEY]
	if okPassword && okKey {
		return entity.Option{}, errors.New("error: can only choose one( -x | -k )")
	}
	if !okPassword && !okKey {
		return entity.Option{}, errors.New("error: required choose one( -x | -k )")
	}
	if okPassword && mPassword != "" {
		util.PasswordWarning()
	}
	if okPassword && mPassword == "" && !okKey {
		fmt.Print("🔑: ")
		password, err := term.ReadPassword(0)
		if err != nil {
			return entity.Option{}, errors.New("error: password exception")
		}
		optionMaps[sys.OPTION_S_PASSWORD] = string(password)
		mPassword, okPassword = optionMaps[sys.OPTION_S_PASSWORD]
		fmt.Println()
	}

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
	} else {
		if _, err = util.CurrentUser(); err != nil {
			return entity.Option{}, errors.New("error: get current user exception")
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

	return CreateHelpOptionMapsToOption(optionMaps)
}

// CreateHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: -a", "-p", "-u", "-x", "-k", "-t", "-g", "-q"
func CreateHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
	var option = entity.Option{}
	option.SetAddress(optionMaps[sys.OPTION_S_ADDRESS])
	var mPort, okPort = optionMaps[sys.OPTION_S_PORT]
	if okPort {
		var port, _ = entity.GetOptionMapsByPort(mPort)
		option.SetPort(port)
	} else {
		option.SetPort(sys.DEFAULT_PORT)
	}
	var mUsername, okUsername = optionMaps[sys.OPTION_S_USERNAME]
	if okUsername {
		var username, _ = entity.GetOptionMapsByUsername(mUsername)
		option.SetUsername(username)
	} else {
		var currentuser, _ = util.CurrentUser()
		option.SetUsername(currentuser)
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
	} else {
		option.SetGroup(sys.DEFAULT_GROUP)
	}
	var _, okQuiet = optionMaps[sys.OPTION_S_QUIET]
	if okQuiet {
		option.SetQuiet(okQuiet)
	}
	return option, nil
}

// CreateHelpUsagePrintText general help text
func CreateHelpUsagePrintText() string {
	return `Run 'kush create --help' for usage.`
}

// CreateHelpUsage print create command help usage
func CreateHelpUsage() {
	var printText = `kush create -a 192.168.1.1 -x 123456 
kush create -a 192.168.1.1 -x 123456 -q
kush create -a 192.168.1.1 -k ~/.ssh/id_rsa_github.pub
` + CreateHelpUsagePrintText()
	util.ExitPrintln(printText)
}
