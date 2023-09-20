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
func CreateCmdRun(option entity.Option) {
	if len(option.Password) > 0 && len(option.Key) > 0 {
		var printText = fmt.Sprintf("error:%s", "can only choose one( -x | -k )")
		util.ExitPrintln(printText)
	}
	var addhost = cli.Host{}
	addhost.Address = option.Address
	addhost.Port = option.Port
	addhost.Username = option.Username
	addhost.Group = option.Group
	addhost.Title = option.Title
	if len(option.Password) > 0 {
		addhost.Authtype = sys.AUTHTYPE_X
		addhost.Password = option.Password
		addhost.Key = ""
	}
	if len(option.Key) > 0 {
		addhost.Authtype = sys.AUTHTYPE_K
		addhost.Key = option.Key
		addhost.Password = ""
	}
	confirm := true
	if !option.GetQuiet() {
		fmt.Printf("%s\n\n", ConfirmInformation(addhost))
		confirminput := ""
		fmt.Print("Please confirm the above information? Y/N: ")
		fmt.Scanln(&confirminput)
		if len(confirminput) <= 0 || strings.ToUpper(confirminput) != "Y" {
			confirm = false
		}
	}
	if confirm {
		err := util.CreateDataPathJsonFile(addhost)
		if err != nil {
			util.ExitPrintln(fmt.Sprint(err))
		}
	}
}

// ConfirmInformation console confirm information
func ConfirmInformation(host cli.Host) string {
	var securetext = ""
	switch host.Authtype {
	case sys.AUTHTYPE_K:
		securetext = "Key: " + host.Key
	case sys.AUTHTYPE_X:
		// bytePassword := []byte(host.Password)
		// resBytes, _ := bcrypt.GenerateFromPassword(bytePassword, 15) //string(resBytes)
		// err := bcrypt.CompareHashAndPassword(resBytes, bytePassword)
		securetext = "Password: ******"
	}
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		"Address: "+host.Address,
		"Username: "+host.Username,
		"Port: "+strconv.Itoa(int(host.Port)),
		securetext,
		"Group: "+host.Group,
		"Title: "+host.Title)

}
