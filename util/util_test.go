// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/anigkus/kush/cli"
)

func TestArgStringToMap(t *testing.T) {
	argString := "-H x -j 12 -F GROUP=group1 -X 1"
	fmt.Println(argString)
	var mapResult, err = StringToMap(argString)
	if err != nil {
		t.Errorf("call ArgStringToMap: %s", err)
	} else {
		fmt.Println(mapResult)
	}

	argString = "-H x -j 12 -F GROUP=group1 -U"
	fmt.Println(argString)
	mapResult, err = StringToMap(argString)
	if err != nil {
		t.Errorf("call ArgStringToMap: %s", err)
	} else {
		fmt.Println(len(mapResult))
		fmt.Println(mapResult)
	}
}

func TestArrayToMap(t *testing.T) {
	// argString := "-H x -j 12 -F GROUP=group1 -X"
	// argString := "-H 192.168.1.1 -W -C 'ADDRESS,USER' -Q"
	// argString := "-H 192.168.1.1 -W --output json -C 'ADDRESS,USER' -Q"
	argString := "-H 192.168.1.1 -W --output json -C 'ADDRESS,USER' -Q -P 123-123 -I"
	// argString := "-H 192.168.1.1 -W --output json -C 'ADDRESS,USER' -Q -P 123-123 -F 'ADDRESS = 1.1.1.1 || USER = user' -I"
	fmt.Println(argString)
	//
	//var reg = regexp.MustCompile(`\s+(')?.*!?\s+(')?`) //\s+(')?.*(')?
	//var elements []string = reg.Split(strings.TrimSpace(argString), -1)

	var elements []string = strings.Split(argString, " ")
	var mapResult, err = StringArrayToMap(elements)
	if err != nil {
		t.Errorf("call ArrayToMap: %s", err)
	} else {
		fmt.Println(len(mapResult))
		fmt.Println(mapResult)
	}
}

func TestDirectJoin(t *testing.T) {
	// argString := "root@192.168.1.1 -X 12345"
	// argString := "root@192.168.1.1 -X 12345 -P 22"

	// argString := "-X 12345 root@192.168.1.1 "
	argString := "-X 12345 root@192.168.1.1 -P 22"

	//bad
	// argString := "-X root@192.168.1.1"
	// argString := "-K root@192.168.1.1"
	// argString := "-X 123456"
	//argString := "root@192.168.1.1 -A 123 -U abc -P 22"
	// argString := "-A 123 -U abc -P 22"
	fmt.Println(argString)
	var mapResult, err = DirectStringToMap(argString)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
	} else {
		fmt.Println(mapResult)
	}
}

func TestManyIsEmpty(t *testing.T) {
	argString := "-H x -j 12 -F GROUP=group1 -X"
	var elements []string = strings.Split(argString, " ")
	var mapResult, err = StringArrayToMap(elements)
	if err != nil {
		t.Errorf("call ArrayToMap: %s", err)
	}
	fmt.Println(mapResult)
	fmt.Println(ManyIsEmpty(mapResult["-H"])) //nil
	fmt.Println(ManyIsEmpty(mapResult["-I"])) //element emtpy
}

func TestTrim(t *testing.T) {
	fmt.Println(strings.Trim(" 192.168.1.1 ", " "))

	fmt.Println(strings.Contains("H-u", "-"))
	fmt.Println(strings.Index("H-u", "-"))

	fmt.Println(strings.Contains("-H-u", "-"))
	fmt.Println(strings.Index("-H-u", "-"))

	var optionMaps map[string]string = map[string]string{}
	optionMaps["A"] = "A"
	optionMaps["B"] = ""

	fmt.Println(optionMaps["A"] == "")
	val, ok := optionMaps["B"]
	fmt.Println(val, ok)
	val, ok = optionMaps["C"]
	fmt.Println(val, ok)
}

func TestReflect(t *testing.T) {

	// type Student struct {
	// 	Fname  string
	// 	Lname  string
	// 	City   string
	// 	Mobile int64
	// }
	type OptionTest struct {
		Address  string
		Port     int32
		Username string
		Password string
		Key      []byte
		Title    string
		Filter   []string
		Quiet    bool
		Output   string
		Columns  []string
		Wide     bool
		Group    string
		Sort     string
	}
	//s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	s := OptionTest{Address: "Chetan", Title: "Kumar", Group: "Bangalore"}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()
	structfield, bo := typeOfS.FieldByName("Address")
	if bo {
		fmt.Println("xx:" + structfield.Name)
		fmt.Printf("yy:%s\n", v.FieldByName(structfield.Name).Interface())
	}
	for i := 0; i < v.NumField(); i++ {
		var element = fmt.Sprintf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		fmt.Print(element)
	}
}

func TestOS(t *testing.T) {
	fmt.Println(os.PathSeparator)
	fmt.Println(os.PathListSeparator)
	fmt.Println(filepath.Join("a", "b", "c"))
}

func TestInsertDataPathJsonFile(t *testing.T) {
	host1 := cli.Host{
		Address:  "1.2.3.4",
		Username: "Alice",
		Port:     22,
		Password: "123456",
		Group:    "Alice",
		Title:    "Alice",
		Authtype: "-X",
	}

	host2 := cli.Host{
		Address:  "10.11.12.13",
		Username: "Bob",
		Port:     100,
		Key:      "~/b.pem",
		Group:    "Bob",
		Title:    "Bob",
		Authtype: "-K",
	}
	host3 := cli.Host{
		Address:  "100.101.102.103",
		Username: "Cici",
		Port:     50,
		Key:      "~/a.pem",
		Group:    "Cici",
		Title:    "Cici",
		Authtype: "-K",
	}
	err := CreateDataPathJsonFile(host3, host1, host2)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
}

func TestSearchDataPathJsonFile(t *testing.T) {
	var argfilter = ""
	argfilter = "ADDRESS =  100.101.102.103"
	argfilter = "USERNAME =  Cici"
	argfilter = "PORT =  50"
	argfilter = "AUTHTYPE =  -K "
	argfilter = "GROUP =  Cici"
	argfilter = "TITLE =  Cici"
	argfilter = "ADDRESS = 10.11.12.13 && USERNAME = Bob"
	argfilter = "ADDRESS = 1.2.3.4 || USERNAME = Cici"

	kushjson, err := GetKushJson()
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	bytes, err := ReadFileToBytes(kushjson)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	diskhosts, err := ByteToHosts(bytes)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	filtermap, err := ArgFilterToMap(argfilter)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	searchhostmaps, err := HostMapsFilter(filtermap, diskhosts)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	searchhostsmapscolumns, err := HostMapsReturnColumns(searchhostmaps, "USERNAME")
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	jsonStr, err := json.MarshalIndent(searchhostsmapscolumns, "", "")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}

func TestSortMap(t *testing.T) {
	basket := map[string]int{"orange": 5, "apple": 7, "mango": 3, "strawberry": 9}
	keys := make([]string, 0, len(basket))
	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys) // asc
	// sort.Sort(sort.Reverse(sort.StringSlice(keys))) // desc
	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
	fmt.Printf("\n")
	var searchhostmaps []map[string]string = []map[string]string{}
	var searchhostmap1 map[string]string = map[string]string{
		"Address":  "1.2.3.4",
		"Username": "Alice",
		"Port":     "22",
		"Password": "123456",
		"Group":    "Alice",
		"Title":    "Alice",
		"Authtype": "-X",
	}
	var searchhostmap2 map[string]string = map[string]string{
		"Address":  "10.11.12.13",
		"Username": "Bob",
		"Port":     "100",
		"Key":      "~/b.pem",
		"Group":    "Bob",
		"Title":    "Bob",
		"Authtype": "-K",
	}
	var searchhostmap3 map[string]string = map[string]string{
		"Address":  "100.101.102.103",
		"Username": "Cici",
		"Port":     "50",
		"Key":      "~/a.pem",
		"Group":    "Cici",
		"Title":    "Cici",
		"Authtype": "-K",
	}
	searchhostmaps = append(searchhostmaps, searchhostmap3, searchhostmap1, searchhostmap2)
	jsonStr, err := json.MarshalIndent(searchhostmaps, "", "")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("sort before: " + string(jsonStr))
	}
	fmt.Printf("\n")

	sort.SliceStable(searchhostmaps, func(i, j int) bool {
		return searchhostmaps[i]["Username"] > searchhostmaps[j]["Username"] // desc
		// return searchhostmaps[i]["Username"] < searchhostmaps[j]["Username"] // asc
	})

	jsonStr, err = json.MarshalIndent(searchhostmaps, "", "")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("sort after: " + string(jsonStr))
	}
}

func TestContains(t *testing.T) {
	find := ""
	find = "192.10"
	find = "xxx"
	find = "30"
	find = "Bob"
	fmt.Println(find)

	// host1 := cli.Host{
	// 	Address:  "1.2.3.4",
	// 	Username: "Alice",
	// 	Port:     22,
	// 	Password: "123456",
	// 	Group:    "Alice",
	// 	Title:    "Alice",
	// 	Authtype: "-X",
	// }

	// host2 := cli.Host{
	// 	Address:  "10.11.12.13",
	// 	Username: "Bob",
	// 	Port:     100,
	// 	Key:      "~/b.pem",
	// 	Group:    "Bob",
	// 	Title:    "Bob",
	// 	Authtype: "-K",
	// }
	// host3 := cli.Host{
	// 	Address:  "100.101.102.103",
	// 	Username: "Cici",
	// 	Port:     50,
	// 	Key:      "~/a.pem",
	// 	Group:    "Cici",
	// 	Title:    "Cici",
	// 	Authtype: "-K",
	// }
	// var elements []string = []string{}
	// elements = append(elements, host1.String())
	// elements = append(elements, host2.String())
	// elements = append(elements, host3.String())
	// for _, elememt := range elements {
	// 	if strings.Contains(elememt, find) {
	// 		fmt.Println(elememt)
	// 	}
	// }

	//
	kushjson, err := GetKushJson()
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	bytes, err := ReadFileToBytes(kushjson)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	diskhosts, err := ByteToHosts(bytes)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	filtermap, err := ArgFilterToMap("")
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	searchhostmaps, err := HostMapsFilter(filtermap, diskhosts)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	searchhostsmapscolumns, err := HostMapsReturnColumns(searchhostmaps, "USERNAME")
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	jsonStr, err := json.MarshalIndent(searchhostsmapscolumns, "", "")
	if err != nil {
		fmt.Printf("err:%s\n", err)
	} else {
		fmt.Println(string(jsonStr))
	}
	var matchhostmaps []map[string]string = []map[string]string{}
	for _, searchhostsmapscolumn := range searchhostsmapscolumns {
		var linestring = MapValueToString(searchhostsmapscolumn, ",")
		if strings.Contains(linestring, find) {
			matchhostmaps = append(matchhostmaps, searchhostsmapscolumn)
		}
	}
	matchlen := len(matchhostmaps)
	switch matchlen {
	case 0:
		// no match
	case 1:
		// ssh
	default:
		// match much
	}
}

func Xyz() {
	//right
	// kushjsona, errjson := ReadFileToBytes("kushjson")

	// kushjsonb, errjson := ReadFileToBytes("kushjson")

	// kushjsonc, errjson := ReadFileToBytes("kushjson")

	// //right
	// var kushjsond, errjsond = ReadFileToBytes("kushjson")

	// var kushjsone, errjsone = ReadFileToBytes("kushjson")

	// var kushjsonf, errjsonf = ReadFileToBytes("kushjson")

	// //error
	// var kushjsong, errjson := ReadFileToBytes("kushjson")

	// var kushjsonh, errjson := ReadFileToBytes("kushjson")

	// var kushjsonj, errjson := ReadFileToBytes("kushjson")

}

func TestRemove(t *testing.T) {

	slice := []string{"a", "b", "c", "d", "e", "f", "g"}

	index := 0
	slice = Remove(slice, index)
	fmt.Printf("slice:%v,index:%d\n", slice, index)

	index = 4
	slice = Remove(slice, index)
	fmt.Printf("slice:%v,index:%d\n", slice, index)

	index = 6
	slice = Remove(slice, index)
	fmt.Printf("slice:%v,index:%d\n", slice, index)
}
