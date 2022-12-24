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

// ImportHelpRun
//
// import command help entry
func ImportHelpRun(importArgs []string) {
	//remove self command
	var options []string = importArgs[1:]
	if len(options) > 0 && (options[0] == sys.OPTION_S_HELP || options[0] == sys.OPTION_L_HELP) {
		ImportHelpUsage()
	}
	var option = entity.Option{}
	if len(importArgs) > 1 {
		var argOptionsMap, err = util.StringArrayToMap(options)
		if err != nil {
			var printText = `error: ` + fmt.Sprint(err) + `
	` + ImportHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		if argOptionsMap == nil || len(argOptionsMap) <= 0 {
			var printText = `error: unknown option "` + util.StringArrayToString(options) + `"
	` + ImportHelpUsagePrintText()
			util.ExitPrintln(printText)
		}
		fmt.Printf("argOptionsMap: %v", argOptionsMap)
		if option, err = ImportHelpOptionValidate(argOptionsMap); err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
	cmd.ImportCmdRun(option)
}

// ImportHelpOptionValidate check terminal option arguments
//
// Option: "-O"
func ImportHelpOptionValidate(argOptionsMap map[string]string) (entity.Option, error) {
	var keyError []string
	for key := range argOptionsMap {
		if !util.InArray(key, sys.IMPORT_OPTION) {
			keyError = append(keyError, key)
		}
	}
	if len(keyError) > 0 {
		return entity.Option{}, errors.New(`error: unknown option "` + util.StringArrayToString(keyError) + `"
` + ImportHelpUsagePrintText())
	}
	var optionMaps, err = entity.ParseArgOptions(argOptionsMap)
	if err != nil {
		return entity.Option{}, err
	}
	// -O
	var mOutput, okOutput = optionMaps[sys.OPTION_S_OUTPUT]
	if okOutput {
		var _, err = entity.GetOptionMapsByOutput(mOutput)
		if err != nil {
			return entity.Option{}, err
		}
	}
	return ImportHelpOptionMapsToOption(optionMaps)
}

// ImportHelpOptionMapsToOption convert terminal option arguments to option struct
//
// Assembly parameters: "-O"
func ImportHelpOptionMapsToOption(optionMaps map[string]string) (entity.Option, error) {
	var option = entity.Option{}
	var mOutput, okOutput = optionMaps[sys.OPTION_S_OUTPUT]
	if okOutput {
		var output, _ = entity.GetOptionMapsByOutput(mOutput)
		option.SetOutput(output)
	}
	return option, nil
}

// ImportHelpUsagePrintText general help text
func ImportHelpUsagePrintText() string {
	return `Run 'kush import --help' for usage.`
}

// ImportHelpUsage print Import command help usage
func ImportHelpUsage() {
	var printText = `kush import -o json > /tmp/kush.json
kush import -o json -C 'ADDRESS,USERNAME' > ~/.kush/kush.json
` + ImportHelpUsagePrintText()
	util.ExitPrintln(printText)
}
