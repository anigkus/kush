// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// SearchCmdRun search command cmd entry
func SearchCmdRun(option entity.Option) {
	if len(option.Filter) <= 0 {
		option.Filter = map[string]string{}
	}
	if len(option.Address) > 0 {
		option.Filter[sys.OPTION_FILTERS[0]] = option.Address
	}
	if option.Port > 0 {
		option.Filter[sys.OPTION_FILTERS[1]] = strconv.Itoa(int(option.Port))
	}
	if len(option.Username) > 0 {
		option.Filter[sys.OPTION_FILTERS[2]] = option.Username
	}
	if len(option.Title) > 0 {
		option.Filter[sys.OPTION_FILTERS[3]] = option.Title
	}
	if len(option.Group) > 0 {
		option.Filter[sys.OPTION_FILTERS[4]] = option.Group
	}
	if len(option.Columns) <= 0 {
		option.Columns = append(option.Columns, sys.COLUMN_ADDRESS)
	}
	if len(option.Sort) <= 0 {
		option.Sort = sys.DEFAULT_SORT
	}
	//
	kushjson, err := util.GetKushJson()
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
	bytes, err := util.ReadFileToBytes(kushjson)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
	diskhosts, err := util.ByteToHosts(bytes)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
	searchhostmaps, err := util.HostMapsFilter(option.Filter, diskhosts)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
	SearchTableInteractive(searchhostmaps, option.Columns, option.Sort)
}

// SearchTableInteractive
func SearchTableInteractive(searchhostmaps []map[string]string, columns []string, sort string) {
	//sort
	searchhostmaps, err := util.HostMapsSort(searchhostmaps, sort)
	if err != nil {
		fmt.Printf("%s.\n", fmt.Sprintf("error: %v", err))
	}
	util.RenderTable(searchhostmaps, columns...)
	// user
	var filterinput string = ""
	fmt.Print("> ")
	fmt.Scanln(&filterinput)

	switch filterinput {
	case "exit":
		os.Exit(0)
	case "help":
		util.Clear()
		SearchTableInteractive(searchhostmaps, columns, sort)
	case "clear":
		util.Clear()
		SearchTableInteractive(searchhostmaps, columns, sort)
	default:
		if len(filterinput) <= 0 {
			SearchTableInteractive(searchhostmaps, columns, sort)
			return
		}
		var matchhostmaps []map[string]string = []map[string]string{}
		for _, searchhostmap := range searchhostmaps {
			var linestring = strings.ToLower(util.MapValueToString(searchhostmap, ","))
			if strings.Contains(linestring, strings.ToLower(filterinput)) {
				matchhostmaps = append(matchhostmaps, searchhostmap)
			}
		}
		matchlen := len(matchhostmaps)
		switch matchlen {
		case 0:
			// no match
			fmt.Printf("No match.\n")
			SearchTableInteractive(searchhostmaps, columns, sort)
		case 1:
			// ssh
			var terminal, err = util.SearchHostToTerminal(matchhostmaps[0])
			if err != nil {
				util.ExitPrintln(fmt.Sprint(err))
			}
			err = terminal.RunTerminal(os.Stdout, os.Stdin)
			if err != nil {
				fmt.Printf("%s.\n", fmt.Sprintf("error: %v", err))
				SearchTableInteractive(searchhostmaps, columns, sort)
			}
		default:
			//match much
			SearchTableInteractive(matchhostmaps, columns, sort)
		}
	}
}
