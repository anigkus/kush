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

// CreateCmdRun create command cmd entry
func UpdateCmdRun(option entity.Option) {
	strport := ""
	if option.Port > 0 {
		strport = strconv.Itoa(int(option.Port))
	}
	if err := util.ManyIsNotEmpty(option.Username, option.Key, option.Password, strport, option.Title, option.Group); err != nil {
		util.ExitPrintln(fmt.Errorf("%s%s", "error: ", "at least one field needs to be updated"))
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
	if len(diskhosts) <= 0 {
		util.ExitPrintln(fmt.Errorf("%s%s", "error: ", "there is no match to the rocord that needs to be updated"))
	}
	matchnum := 0
	for _, diskhost := range diskhosts {
		if diskhost.Address == option.Address {
			matchnum++
		}
	}
	if matchnum <= 0 {
		util.ExitPrintln(fmt.Errorf("%s%s", "error: ", "there is no match to the recore that needs to be updated"))
	}
	if matchnum != 1 {
		util.ExitPrintln(fmt.Errorf("%s%s", "error: ", "only one record can be updated at a time"))
	}

	confirm := true
	if !option.GetQuiet() {
		confirminput := ""
		fmt.Printf("%s\n\n", UpdateConfirmInformation(option))
		fmt.Print("Please confirm the above information? Y/N: ")
		fmt.Scanln(&confirminput)
		if len(confirminput) <= 0 || strings.ToUpper(confirminput) != "Y" {
			confirm = false
		}
	}
	if confirm {
		var updatehosts []cli.Host = []cli.Host{}
		for _, diskhost := range diskhosts {
			if diskhost.Address == option.Address {
				if len(option.Username) > 0 {
					diskhost.Username = option.Username
				}
				if len(option.Password) > 0 {
					diskhost.Password = option.Password
					diskhost.Authtype = sys.AUTHTYPE_X
				}
				if len(option.Key) > 0 {
					diskhost.Key = option.Key
					diskhost.Authtype = sys.AUTHTYPE_K
				}
				if option.Port > 0 {
					diskhost.Port = option.Port
				}
				if len(option.Group) > 0 {
					diskhost.Group = option.Group
				}
				if len(option.Title) > 0 {
					diskhost.Title = option.Title
				}
				updatehosts = append(updatehosts, diskhost)
				break
			}
			updatehosts = append(updatehosts, diskhost)
		}
		err := util.UpdateDataPathJsonFile(updatehosts)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
}

// UpdateConfirmInformation console confirm information
func UpdateConfirmInformation(option entity.Option) string {
	var informations []string = []string{}

	if len(option.Username) > 0 {
		informations = append(informations, fmt.Sprintf(
			"Username: %s",
			option.Username))
	}
	if len(option.Password) > 0 {
		informations = append(informations, fmt.Sprintf(
			"Password: %s",
			option.Password))
	}
	if len(option.Key) > 0 {
		informations = append(informations, fmt.Sprintf(
			"Key: %s",
			option.Key))
	}
	if option.Port > 0 {
		informations = append(informations, fmt.Sprintf(
			"Port: %s",
			strconv.Itoa(int(option.Port))))
	}
	if len(option.Group) > 0 {
		informations = append(informations, fmt.Sprintf(
			"Group: %s",
			option.Group))
	}
	if len(option.Title) > 0 {
		informations = append(informations, fmt.Sprintf(
			"Title: %s",
			option.Title))
	}
	return strings.Join(informations, "\n")
}
