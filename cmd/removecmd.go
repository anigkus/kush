// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anigkus/kush/cli"
	"github.com/anigkus/kush/entity"
	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

// RemoveCmdRun remove command cmd entry
func RemoveCmdRun(option entity.Option) {
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
	if len(searchhostmaps) <= 0 {
		util.ExitPrintln(fmt.Errorf("%s%s", "error: ", "there is no match to the record that needs to be deleted"))
	}
	confirm := true
	if !option.GetQuiet() {
		confirminput := ""
		util.RenderTable(searchhostmaps, sys.COLUMN_ADDRESS, sys.COLUMN_USERNAME, sys.COLUMN_PORT, sys.COLUMN_AUTHTYPE, sys.COLUMN_AUTH)
		fmt.Print("Please confirm the above information? Y/N: ")
		fmt.Scanln(&confirminput)
		if len(confirminput) <= 0 || strings.ToUpper(confirminput) != "Y" {
			confirm = false
		}
	}
	var delhosts []cli.Host = []cli.Host{}
	for _, searchhostmap := range searchhostmaps {
		var delhost = cli.Host{}
		delhost.Address = searchhostmap[sys.HOST_ADDRESS]
		delhost.Username = searchhostmap[sys.HOST_USERNAME]
		var port, _ = strconv.Atoi(searchhostmap[sys.HOST_PORT])
		delhost.Port = int32(port)
		delhost.Authtype = searchhostmap[sys.HOST_AUTHTYPE]
		delhost.Key = searchhostmap[sys.HOST_KEY]
		delhost.Password = searchhostmap[sys.HOST_PASSWORD]
		delhosts = append(delhosts, delhost)
	}
	if confirm {
		err := util.RemoveDataPathJsonFile(delhosts...)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
}

// RemoveConfirmInformation console confirm information
func RemoveConfirmInformation(searchhostmaps []map[string]string) string {
	var informations []string = []string{}
	for _, searchhostmap := range searchhostmaps {
		var securetext = ""
		switch searchhostmap[sys.HOST_AUTHTYPE] {
		case sys.AUTHTYPE_K:
			securetext = "Key: " + searchhostmap[sys.HOST_KEY]
		case sys.AUTHTYPE_X:
			securetext = "Password: " + searchhostmap[sys.HOST_PASSWORD]
		}
		var outtext = fmt.Sprintf(
			"%s , %s , %s , %s",
			"Address: "+searchhostmap[sys.HOST_ADDRESS],
			"Username: "+searchhostmap[sys.HOST_USERNAME],
			"Port: "+searchhostmap[sys.HOST_PORT],
			securetext)
		informations = append(informations, outtext)
	}
	return strings.Join(informations, "\n")
}
