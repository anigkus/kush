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

// ExportHelpRun
//
// export command help entry
func ExportHelpRun(exportArgs []string) {
	//remove self command
	var options []string = exportArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		ExportHelpUsage()
	}
	var option = entity.Option{}
	if len(exportArgs) > 1 {
		var argOptionsMap, err = util.StringArrayToMap(options)
		if err != nil {
			var printText = `error: ` + fmt.Sprint(err) + `
	` + ExportHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if argOptionsMap == nil || len(argOptionsMap) <= 0 {
			var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
	` + ExportHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if option, err = ExportHelpOptionValidate(argOptionsMap); err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
	cmd.ExportCmdRun(option)
}

// ExportHelpOptionValidate check terminal option arguments
//
// Option: "-A", "-P", "-U", "-T", "-G", "-F", "-C", "-O"
func ExportHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.EXPORT_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + ExportHelpUsagePrintText())
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
	//-C
	var mColumns, okColumns = optionMaps[sys.OPTION_S_COLUMNS]
	if okColumns {
		var _, err = entity.GetOptionMapsByColumns(mColumns)
		if err != nil {
			return entity.Option{}, err
		}
	}
	// -O
	var mOutput, okOutput = optionMaps[sys.OPTION_S_OUTPUT]
	if okOutput {
		var _, err = entity.GetOptionMapsByOutput(mOutput)
		if err != nil {
			return entity.Option{}, err
		}
	}
	return ExportHelpOptionMapsToOption(optionMaps)
}

// ExportHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: "-A", "-P", "-U", "-T", "-G", "-F", "-C", "-O"
func ExportHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
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
	var mColumns, okColumns = optionMaps[sys.OPTION_S_COLUMNS]
	if okColumns {
		var columns, _ = entity.GetOptionMapsByColumns(mColumns)
		option.SetColumns(columns)
	}
	var mOutput, okOutput = optionMaps[sys.OPTION_S_OUTPUT]
	if okOutput {
		var output, _ = entity.GetOptionMapsByOutput(mOutput)
		option.SetOutput(output)
	}
	return option, nil
}

// ExportHelpUsagePrintText general help text
func ExportHelpUsagePrintText() string {
	return `Run 'kush export --help' for usage.`
}

// ExportHelpUsage print Export command help usage
func ExportHelpUsage() {
	var printText = `kush export -o json > /tmp/kush.json
kush export -o json -C 'ADDRESS,USERNAME' > ~/.kush/kush.json
` + ExportHelpUsagePrintText()
	util.ExitPrintln(printText)
}
