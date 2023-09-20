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

// SearchHelpRun
//
// search command help entry
func SearchHelpRun(searchArgs []string) {
	//remove self command
	var options []string = searchArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		SearchHelpUsage()
	}
	var option = entity.Option{}
	if len(searchArgs) > 1 {
		var argOptionsMap, err = util.StringArrayToMap(options)
		if err != nil {
			var printText = `error: ` + fmt.Sprint(err) + `
	` + SearchHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if argOptionsMap == nil || len(argOptionsMap) <= 0 {
			var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
	` + SearchHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if option, err = SearchHelpOptionValidate(argOptionsMap); err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
	cmd.SearchCmdRun(option)
}

// SearchHelpOptionValidate check terminal option arguments
//
// Option: "-a", "-p", "-u", "-t", "-g", "-f", "-c", "-s"
func SearchHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.SEARCH_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + SearchHelpUsagePrintText())
	}
	var optionMaps, err = entity.ParseArgOptions(argOptionsMap)
	if err != nil {
		return entity.Option{}, err
	}

	//-a
	var mAdderess, okAdderess = optionMaps[sys.OPTION_S_ADDRESS]
	if okAdderess {
		if _, err = entity.GetOptionMapsByAddress(mAdderess); err != nil {
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
	//-f
	var mFilter, okFilter = optionMaps[sys.OPTION_S_FILTER]
	if okFilter {
		if _, err = entity.GetOptionMapsByFilter(mFilter); err != nil {
			return entity.Option{}, err
		}
	}
	//-c
	var mColumns, okColumns = optionMaps[sys.OPTION_S_COLUMNS]
	if okColumns {
		var _, err = entity.GetOptionMapsByColumns(mColumns)
		if err != nil {
			return entity.Option{}, err
		}
	}
	//-s
	var mSort, okSort = optionMaps[sys.OPTION_S_SORT]
	if okSort {
		var _, err = entity.GetOptionMapsBySort(mSort)
		if err != nil {
			return entity.Option{}, err
		}
	}

	return SearchHelpOptionMapsToOption(optionMaps)
}

// SearchHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: "-a", "-p", "-u", "-t", "-g", "-f", "-c", "-s"
func SearchHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
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
	var mSort, okSort = optionMaps[sys.OPTION_S_SORT]
	if okSort {
		var sort, _ = entity.GetOptionMapsBySort(mSort)
		option.SetSort(sort)
	}
	return option, nil
}

// SearchHelpUsagePrintText general help text
func SearchHelpUsagePrintText() string {
	return `Run 'kush search --help' for usage.`
}

// SearchHelpUsage print Search command help usage
func SearchHelpUsage() {
	var printText = `kush search -a 192.168.1.1 
kush search -a test.host1.com -f 'GROUP=group1' -ss ADDRESS
` + SearchHelpUsagePrintText()
	util.ExitPrintln(printText)
}
