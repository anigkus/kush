// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// ExportCmdRun export command cmd entry
func ExportCmdRun(option entity.Option) {
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
		//full
		option.Columns = sys.OPTION_COLUMNS
	}
	if len(option.Output) <= 0 {
		option.Output = sys.DEFAULT_OUTPUT
	}

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
	searchhostsmapscolumns, err := util.HostMapsReturnColumns(searchhostmaps, option.Columns...)
	if err != nil {
		util.ExitPrintln(fmt.Sprintf("error: %v", err))
	}
	err = util.ExportDataPathFile(option.Output, searchhostsmapscolumns...)
	if err != nil {
		util.ExitPrintln(fmt.Sprint(err))
	}
}
