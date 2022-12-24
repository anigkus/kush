// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"reflect"
	"strings"
)

type Host struct {
	Address  string `json:"Address"`
	Port     int32  `json:"Port"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Key      string `json:"Key"`
	Group    string `json:"Group"`
	Title    string `json:"Title"`
	Authtype string `json:"Authtype"`
}

func (host Host) String() string {
	valueof := reflect.ValueOf(host)
	var elements []string
	for i := 0; i < valueof.NumField(); i++ {
		var element = fmt.Sprint(valueof.Field(i).Interface())
		elements = append(elements, element)
	}
	return strings.Join(elements, ",")
}
