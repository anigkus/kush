// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/anigkus/kush/cli"
	"github.com/anigkus/kush/sys"
)

// InArray check if an element is in an array
func InArray(target string, arrays []string) bool {
	for _, element := range arrays {
		if target == element {
			return true
		}
	}
	return false
}

// DirectStringToMap
func DirectStringToMap(argString string) (map[string]string, error) {
	if err := ManyIsEmpty(argString); err != nil {
		return nil, fmt.Errorf("%s%s", "error: ", "arguments empty")
	}
	argString = strings.Trim(argString, " ")
	var elements []string = strings.Split(argString, " ")
	return DirectArrayToMap(elements)
}

// CheckArrayContainsSubstr
func CheckArrayContainsSubstr(elements []string, substr string) bool {
	for _, element := range elements {
		if strings.Contains(element, sys.AT) {
			return true
		}
	}
	return false
}

// DirectArrayToMap
func DirectArrayToMap(elements []string) (map[string]string, error) {
	if len(elements) > 5 || len(elements) < 3 {
		return nil, fmt.Errorf("%s%s", "error: ", "arguments length illegal")
	}
	//-x -k -p
	// required contains `@`
	if !CheckArrayContainsSubstr(elements, sys.AT) {
		return nil, fmt.Errorf("%s%s", "error: ", "arguments format illegal")
	}
	var mapResult = map[string]string{}
	var keys []string
	var values []string
	var i, size = 0, len(elements)
	for i < size {
		var element, step = elements[i], 1
		if strings.Index(element, "-") == 0 {
			keys = append(keys, element)
			var index = i + 1
			if index >= size {
				index = size - 1
			}
			var nextElement = elements[index]
			if nextElement[:1] != "-" && !strings.Contains(element, sys.AT) {
				values = append(values, nextElement)
				step = 2
			} else {
				values = append(values, "")
			}
		} else {
			if strings.Contains(element, sys.AT) {
				var userhost = strings.Split(element, sys.AT)
				keys = append(keys, sys.OPTION_S_USERNAME)
				keys = append(keys, sys.OPTION_S_ADDRESS)
				values = append(values, userhost[0])
				values = append(values, userhost[1])
			} else {
				values = append(values, element)
			}
		}
		i += step
	}
	var keysLen = len(keys)
	var valuesLen = len(values)
	for index, key := range keys {
		//key and value not equal
		if keysLen != valuesLen && (index+1) > valuesLen {
			mapResult[key] = ""
		} else {
			mapResult[key] = values[index]
		}
	}
	return mapResult, nil
}

// StringToMap convert string to map type
func StringToMap(argString string) (map[string]string, error) {
	var mapResult = map[string]string{}
	if err := ManyIsEmpty(argString); err != nil {
		return mapResult, errors.New("empty")
	}
	argString = strings.Trim(argString, " ")
	var elements []string = strings.Split(argString, " ")
	return StringArrayToMap(elements)
}

// StringArrayToMap convert array to map type
func StringArrayToMap(elements []string) (map[string]string, error) {
	var mapResult = map[string]string{}
	if len(elements) <= 0 {
		return mapResult, errors.New("empty")
	}
	var keys []string
	var values []string
	var i, size = 0, len(elements)
	for i < size {
		var element, step = elements[i], 1
		if strings.Index(element, "-") == 0 {
			keys = append(keys, element)
			var index = i + 1
			if index >= size {
				index = size - 1
			}
			var nextElement = elements[index]
			if nextElement[:1] != "-" {
				values = append(values, nextElement)
				step = 2
			} else {
				values = append(values, "")
			}
		} else {
			values = append(values, element)
		}
		i += step
	}
	var keysLen = len(keys)
	var valuesLen = len(values)
	for index, key := range keys {
		//key and value not equal
		if keysLen != valuesLen && (index+1) > valuesLen {
			mapResult[key] = ""
		} else {
			mapResult[key] = values[index]
		}
	}
	return mapResult, nil
}

// CurrentUser return current execute user
func CurrentUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

// CurrentUserHome return `user.HomeDir`
func CurrentUserHome() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}

// ExitPrintln terminal console output
func ExitPrintln(printText any) {
	fmt.Println(printText)
	os.Exit(0)
}

// ManyIsEmpty multi element is empty check
func ManyIsEmpty(elements ...string) error {
	if len(elements) <= 0 {
		return errors.New("emtpy")
	}
	for _, element := range elements {
		if element == "" || len(element) <= 0 {
			return errors.New("element emtpy")
		}
	}
	return nil
}

// ManyIsNotEmpty multi element is not empty check
func ManyIsNotEmpty(elements ...string) error {
	if len(elements) <= 0 {
		return errors.New("emtpy")
	}
	for _, element := range elements {
		if len(element) > 0 {
			return nil
		}
	}
	return errors.New("all emtpy")
}

// StringArrayToString  arrry to string eg:[A B C] =>A,B,C
func StringArrayToString(elements []string) string {
	if elements == nil || len(elements) <= 0 {
		return ""
	}
	var returnText = ""
	for _, element := range elements {
		returnText += element + ","
	}
	return returnText[:len(returnText)-1]
}

// GetEnv get environment variable value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetDataPathRootHome get current user home path or $KUSH_HOME value to assign data path home
func GetDataPathRootHome() (string, error) {
	var datapathroothome = ""
	if datapathroothome = GetEnv(sys.KUSH_HOME_VAR, datapathroothome); len(datapathroothome) <= 0 {
		var value, err = CurrentUserHome()
		if err != nil {
			return datapathroothome, errors.New("error: get the current user home exception")
		}
		datapathroothome = value
	}
	var fileinfo, erros = os.Stat(datapathroothome)
	if errors.Is(erros, os.ErrNotExist) {
		return datapathroothome, errors.New("error: \"" + datapathroothome + "\" data no such file or directory")
	}
	if !fileinfo.Mode().IsDir() {
		return datapathroothome, errors.New("error: \"" + datapathroothome + "\" not dir")
	}
	// if err := syscall.Access(datapathroothome, syscall.O_RDWR); err != nil {
	// 	return datapathroothome, errors.New("error: " + user + " has no read and write permissions on the \"" + datapathroothome + "\" ")
	// }
	var _, errpath = os.Stat(filepath.Join(datapathroothome, sys.KUSH_HOME_PATH))
	if errors.Is(errpath, os.ErrNotExist) {
		//mkdir
		err := os.Mkdir(filepath.Join(datapathroothome, sys.KUSH_HOME_PATH), sys.PERM_FILE_MODE)
		if err != nil {
			return datapathroothome, fmt.Errorf("%s%s%s", "error: ", sys.KUSH_HOME_PATH, " create fail")
		}
	}
	return datapathroothome, nil
}

// CreateDataPathJsonFile create `$KUSH_HOME/.kush/.kush.json` by host
func CreateDataPathJsonFile(addhosts ...cli.Host) error {
	if len(addhosts) <= 0 {
		return fmt.Errorf("%s%s", "error: ", "append data empty for .kush.json ")
	}
	var kushjson, err = GetKushJson()
	if err != nil {
		return err
	}
	var _, errpath = os.Stat(kushjson)
	if errors.Is(errpath, os.ErrNotExist) {
		_, err := os.Create(kushjson)
		if err != nil {
			return fmt.Errorf("%s%s%s", "error: ", sys.KUSH_PATH_JSON, " create fail")
		}
	}

	// read
	bytehosts, err := ReadFileToBytes(kushjson)
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "read .kush.json exception")
	}
	diskhosts := []cli.Host{}
	if len(bytehosts) > 0 {
		err = json.Unmarshal([]byte(bytehosts), &diskhosts)
		if err != nil {
			return fmt.Errorf("%s%s", "error: ", "unmarshal .kush.json exception")
		}
	}
	// check exist
	for _, addhost := range addhosts {
	lable:
		for j, diskhost := range diskhosts {
			if HostEquals(addhost, diskhost) {
				addhosts = Remove(addhosts, j)
				break lable
			}
		}
	}
	if len(addhosts) <= 0 {
		return fmt.Errorf("%s%s", "error: ", " append host exist or empty")
	}
	// append
	diskhosts = append(diskhosts, addhosts...)
	// write
	data, err := json.MarshalIndent(diskhosts, "", "")
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "MarshalIndent .kush.json exception")
	}
	err = WriteFileToBytes(kushjson, data)
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "write .kush.json exception")
	}
	return nil
}

// RemoveDataPathJsonFile remove `$KUSH_HOME/.kush/.kush.json` by host
func RemoveDataPathJsonFile(delhosts ...cli.Host) error {
	if len(delhosts) <= 0 {
		return fmt.Errorf("%s%s", "error: ", "host empty")
	}
	kushjson, err := GetKushJson()
	if err != nil {
		return err
	}
	bytehosts, err := ReadFileToBytes(kushjson)
	if err != nil {
		return err
	}
	diskhosts, err := ByteToHosts(bytehosts)
	if err != nil {
		return err
	}
	for _, delhost := range delhosts {
	lable:
		for j, diskhost := range diskhosts {
			if HostEquals(delhost, diskhost) {
				diskhosts = Remove(diskhosts, j)
				break lable
			}
		}
	}
	// overwrite
	data, err := json.MarshalIndent(diskhosts, "", "")
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "MarshalIndent .kush.json exception")
	}
	err = WriteFileToBytes(kushjson, data)
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "write .kush.json exception")
	}
	return nil
}

// UpdateDataPathJsonFile update `$KUSH_HOME/.kush/.kush.json` by host
func UpdateDataPathJsonFile(updatehosts []cli.Host) error {
	kushjson, err := GetKushJson()
	if err != nil {
		return err
	}
	// overwrite
	data, err := json.MarshalIndent(updatehosts, "", "")
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "MarshalIndent .kush.json exception")
	}
	err = WriteFileToBytes(kushjson, data)
	if err != nil {
		return fmt.Errorf("%s%s", "error: ", "write .kush.json exception")
	}
	return nil
}

// ExportDataPathFile export host to file stream
func ExportDataPathFile(output string, outhosts ...map[string]string) error {
	switch output {
	case "json":
		{
			bytes, err := json.MarshalIndent(outhosts, "", "")
			if err != nil {
				return fmt.Errorf("%s%s", "error: ", "MarshalIndent outhosts exception")
			}
			ExitPrintln(string(bytes))
		}
	}
	return fmt.Errorf("%s%s", "error: ", "only supports json format")
}

func ImportDataPathFile(inputbytes []byte, output string) error {
	switch output {
	case "json":
		{
			inputhosts := []cli.Host{}
			err := json.Unmarshal(inputbytes, &inputhosts)
			if err != nil {
				return fmt.Errorf("%s%s", "error: ", "json format illegal")
			}
			if len(inputhosts) <= 0 {
				return fmt.Errorf("%s%s", "error: ", "json length illegal")
			}
			for _, inputhost := range inputhosts {
				var address = inputhost.Address
				var port = inputhost.Port
				var username = inputhost.Username
				var authtype = inputhost.Authtype
				var password = inputhost.Password
				var key = inputhost.Key
				if _, err = CheckAddress(address); err != nil {
					return err
				}
				if _, err = CheckPortRange(strconv.Itoa(int(port))); err != nil {
					return err
				}
				if _, err = CheckUsername(username); err != nil {
					return err
				}
				switch authtype {
				case sys.AUTHTYPE_K:
					if _, err = CheckPublicKey(key); err != nil {
						return err
					}
				case sys.AUTHTYPE_X:
					if _, err = CheckPassword(password); err != nil {
						return err
					}
				default:
					return fmt.Errorf("%s%s", "error: ", "authtype required choose one( -X | -K )")
				}
			}
			err = CreateDataPathJsonFile(inputhosts...)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("%s%s", "error: ", "only supports json format")
}

// Remove remove any slice array by index
func Remove[T any](slice []T, s int) []T {
	if len(slice) > 0 && (s+1) >= len(slice) {
		return slice[:len(slice)-1]
	}
	return append(slice[:s], slice[s+1:]...)
}

// HostEquals compare two hosts for equality
func HostEquals(source cli.Host, target cli.Host) bool {
	if source.Address == target.Address &&
		source.Username == target.Username &&
		source.Port == target.Port &&
		source.Authtype == target.Authtype &&
		source.Key == target.Key &&
		source.Password == target.Password {
		return true
	}
	return false
}

// ArgFilterToMap arguments filter to map[string]string
func ArgFilterToMap(argfilter string) (map[string]string, error) {
	var filtermap map[string]string = map[string]string{}
	if len(argfilter) > 0 {
		// -f 'ADDRESS = 1.1.1.1 && USERNAME = user || GROUP = user'
		if strings.Contains(argfilter, sys.LOGICAL_OPERATORS_AND) && strings.Contains(argfilter, sys.LOGICAL_OPERATORS_OR) {
			return nil, fmt.Errorf("%s%s", "error: ", "logical operators && and || can only use one")
		}
		var operators = ""
		var filtersexpressions []string = []string{}
		var operatorsand = strings.Split(argfilter, sys.LOGICAL_OPERATORS_AND)
		var operatorsor = strings.Split(argfilter, sys.LOGICAL_OPERATORS_OR)
		// only two operator
		if len(operatorsand) > 2 || len(operatorsor) > 2 {
			return nil, fmt.Errorf("%s%s", "error: ", "logical operators can only have one")
		}
		if strings.Contains(argfilter, sys.LOGICAL_OPERATORS_AND) {
			operators = sys.LOGICAL_OPERATORS_AND
			filtersexpressions = append(filtersexpressions, operatorsand...)
		} else if strings.Contains(argfilter, sys.LOGICAL_OPERATORS_OR) {
			operators = sys.LOGICAL_OPERATORS_OR
			filtersexpressions = append(filtersexpressions, operatorsor...)
		} else {
			// only one condition
			filtersexpressions = append(filtersexpressions, argfilter)
		}

		// special key
		if len(operators) > 0 {
			filtermap[sys.OPERATORS_KEY] = operators
		}
		for _, filtersexpression := range filtersexpressions {
			var expressions = strings.Split(filtersexpression, "=")
			if len(expressions) != 2 {
				return nil, fmt.Errorf("%s%s%s%s", "error: ", "filter [", filtersexpression, "] format illegal")
			}
			var expressionname = strings.Trim(expressions[0], " ")
			if !InArray(expressionname, sys.OPTION_FILTERS) {
				return nil, fmt.Errorf("%s%s%s%s", "error: ", "filter [", filtersexpression, "] expression illegal")
			}
			var expressionValue = strings.Trim(expressions[1], " ")
			filtermap[expressionname] = expressionValue
		}
	}
	return filtermap, nil
}

// UpperCaseFirstLetter return after convert word first letter to upper case
//
// eg: word->Word , WORD->Word
func UpperCaseFirstLetter(word string) (string, error) {
	if err := ManyIsEmpty(word); err != nil || len(word) < 1 {
		return "", fmt.Errorf("%s%s%s", "error: ", word, " size required > 1 ")
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:]), nil
}

// HostToMap struct to map
func HostToMap(host cli.Host) (map[string]string, error) {
	if (host == cli.Host{}) {
		return nil, fmt.Errorf("%s%s", "error: ", "host data empty")
	}
	var hostmap map[string]string = map[string]string{}
	valueof := reflect.ValueOf(host)
	vtype := valueof.Type()
	for i := 0; i < valueof.NumField(); i++ {
		var fieldname, fieldvalue = vtype.Field(i).Name, valueof.Field(i).Interface()
		hostmap[fieldname] = fmt.Sprint(fieldvalue)
	}
	return hostmap, nil
}

// ByteToHosts parse byte to host
func ByteToHosts(bytes []byte) ([]cli.Host, error) {
	diskhosts := []cli.Host{}
	err := json.Unmarshal([]byte(bytes), &diskhosts)
	if err != nil {
		return nil, fmt.Errorf("%s%s", "error: ", "unmarshal json exception")
	}
	return diskhosts, nil
}

// GetKushJson return .kush.json absolute path
func GetKushJson() (string, error) {
	var datapathroothome, err = GetDataPathRootHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(datapathroothome, sys.KUSH_HOME_PATH, sys.KUSH_PATH_JSON), nil
}

// AppendSearchHostMaps append not exist elements
func AppendSearchHostMaps(searchhostmaps []map[string]string, hostmap map[string]string) []map[string]string {
	var matchElement bool = true
	for _, searchhostmap := range searchhostmaps {
		var searchhosttext = fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s",
			searchhostmap[sys.HOST_ADDRESS],
			searchhostmap[sys.HOST_PORT],
			searchhostmap[sys.HOST_USERNAME],
			searchhostmap[sys.HOST_AUTHTYPE],
			searchhostmap[sys.HOST_PASSWORD],
			searchhostmap[sys.HOST_KEY],
			searchhostmap[sys.HOST_GROUP],
			searchhostmap[sys.HOST_TITLE])
		var hostmaptext = fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s",
			hostmap[sys.HOST_ADDRESS],
			hostmap[sys.HOST_PORT],
			hostmap[sys.HOST_USERNAME],
			hostmap[sys.HOST_AUTHTYPE],
			hostmap[sys.HOST_PASSWORD],
			hostmap[sys.HOST_KEY],
			hostmap[sys.HOST_GROUP],
			hostmap[sys.HOST_TITLE])
		if hostmaptext == searchhosttext {
			matchElement = false
			break
		}
	}
	if matchElement {
		searchhostmaps = append(searchhostmaps, hostmap)
	}
	return searchhostmaps
}

// HostMapsFilter filter host map collections
func HostMapsFilter(filtermap map[string]string, diskhosts []cli.Host) ([]map[string]string, error) {
	var searchhostmaps []map[string]string = []map[string]string{}
	// filters
	if len(filtermap) > 0 {
		// special key
		var operators, okOperators = filtermap[sys.OPERATORS_KEY]
		if okOperators && len(operators) > 0 {
			//not normal expression
			delete(filtermap, sys.OPERATORS_KEY)
		}
		var filtersize = len(filtermap)
		if filtersize > 0 {
			// much expression
			for _, diskhost := range diskhosts { // 4
				valueof := reflect.ValueOf(diskhost)
				vtype := valueof.Type()
				var matchsize = filtersize
				for key, value := range filtermap { // 2
					upperletter, err := UpperCaseFirstLetter(key)
					if err != nil {
						return nil, err
					}
					structfield, _ := vtype.FieldByName(upperletter)
					hostValue := fmt.Sprint(valueof.FieldByName(structfield.Name).Interface())
					//-f 'ADDRESS = 1.1.1.1 && USERNAME = test'
					// &&
					if operators == sys.LOGICAL_OPERATORS_AND && hostValue == value {
						matchsize--
					} else if operators == sys.LOGICAL_OPERATORS_OR && hostValue == value {
						// ||
						hostmap, _ := HostToMap(diskhost)
						searchhostmaps = AppendSearchHostMaps(searchhostmaps, hostmap)
					} else if len(operators) <= 0 && hostValue == value {
						hostmap, _ := HostToMap(diskhost)
						searchhostmaps = AppendSearchHostMaps(searchhostmaps, hostmap)
					}
				}
				if matchsize == 0 {
					hostmap, _ := HostToMap(diskhost)
					searchhostmaps = AppendSearchHostMaps(searchhostmaps, hostmap)
				}
			}
		}
	} else {
		for _, diskhost := range diskhosts {
			hostmap, _ := HostToMap(diskhost)
			searchhostmaps = AppendSearchHostMaps(searchhostmaps, hostmap)
		}
	}
	return searchhostmaps, nil
}

// HostMapsSort sort host map collections
func HostMapsSort(searchhostmaps []map[string]string, sortfield string) ([]map[string]string, error) {
	// sort
	if len(sortfield) > 0 {
		if !InArray(sortfield, sys.OPTION_SORTS) {
			return nil, fmt.Errorf("%s%s%s%s", "error: ", "sort [", sortfield, "] illegal")
		}
		upperletter, err := UpperCaseFirstLetter(sortfield)
		if err != nil {
			return nil, err
		}
		sort.SliceStable(searchhostmaps, func(i, j int) bool {
			// return searchhostmaps[i][upperletter] > searchhostmaps[j][upperletter] // desc
			return searchhostmaps[i][upperletter] < searchhostmaps[j][upperletter] // asc
		})
	}
	return searchhostmaps, nil
}

// HostMapsReturnColumns search collections show required columns
func HostMapsReturnColumns(searchhostmaps []map[string]string, columns ...string) ([]map[string]string, error) {
	var searchhostsmapscolumns []map[string]string = []map[string]string{}
	// columns
	if len(columns) > 0 {
		for _, column := range columns {
			if !InArray(column, sys.OPTION_COLUMNS) {
				return nil, fmt.Errorf("%s%s%s%s", "error: ", "columns [", column, "] illegal")
			}
		}
		for _, searchhostmap := range searchhostmaps {
			var searchhostsmapcolumn map[string]string = map[string]string{}
			for _, column := range columns {
				upperletter, err := UpperCaseFirstLetter(column)
				if err != nil {
					return nil, err
				}
				// auth
				var value = searchhostmap[upperletter]
				if column == sys.COLUMN_AUTH {
					var authtype = searchhostmap[sys.HOST_AUTHTYPE]
					switch authtype {
					case sys.AUTHTYPE_X:
						value = searchhostmap[sys.HOST_PASSWORD]
					case sys.AUTHTYPE_K:
						value = searchhostmap[sys.HOST_KEY]
					}
				}
				searchhostsmapcolumn[upperletter] = value
			}
			searchhostsmapscolumns = append(searchhostsmapscolumns, searchhostsmapcolumn)
		}
	}
	return searchhostsmapscolumns, nil
}

// ReadFileToByte parse file to bytes
func ReadFileToBytes(file string) ([]byte, error) {
	if err := ManyIsEmpty(file); err != nil {
		return nil, fmt.Errorf("file is empty")
	}
	fileinfo, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("error: \"" + file + "\" file no such file or directory")
	}
	if !fileinfo.Mode().IsRegular() {
		return nil, errors.New("error: \"" + file + "\" not regular")
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("error: read file key error")
	}
	//decode
	sDec, err := base64.StdEncoding.DecodeString(string(bytes))
	if err != nil {
		return nil, errors.New("error: decode file error")
	}
	return sDec, nil
}

// WriteFileToBytes bytes to file
func WriteFileToBytes(file string, data []byte) error {
	//encode
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return os.WriteFile(file, []byte(sEnc), sys.PERM_FILE_MODE)
}

// MapValueToString the map value is converted into a string and join using sep
//
// map[string]string
// {
// "key1":"value1",
// "key2":"value2"
// }
// => return "value1,value2"
func MapValueToString(arrayMaps map[string]string, sep string) string {
	if len(arrayMaps) <= 0 {
		return ""
	}
	var elements []string
	for _, value := range arrayMaps {
		elements = append(elements, value)
	}
	return strings.Join(elements, sep)
}

// CheckAddress validata ssh address
func CheckAddress(address string) (string, error) {
	var err = ManyIsEmpty(address)
	if err != nil {
		return "", errors.New("error: address cannot be empty")
	}
	return address, nil
}

// CheckUsername validata ssh username
func CheckUsername(username string) (string, error) {
	var err = ManyIsEmpty(username)
	if err != nil {
		return "", errors.New("error: username cannot be empty")
	}
	return username, nil
}

// CheckPassword validata ssh password
func CheckPassword(password string) (string, error) {
	var err = ManyIsEmpty(password)
	if err != nil {
		return "", errors.New("error: password cannot be empty")
	}
	return password, nil
}

// CheckPortRange validata ssh port
func CheckPortRange(port string) (int32, error) {
	var err = ManyIsEmpty(port)
	if err != nil {
		return -1, errors.New("error: port cannot be empty")
	}
	var int64port int64 = -1
	if int64port, err = strconv.ParseInt(port, 0, 0); err != nil {
		return -1, errors.New("error: port not number")
	}
	if int64port < 0 || int64port > 65535 {
		return -1, errors.New("error: port out of range 1-65535")
	}
	return int32(int64port), nil
}

// CheckPublicKey validata ssh publicKey
func CheckPublicKey(key string) (string, error) {
	var err = ManyIsEmpty(key)
	if err != nil {
		return "", errors.New("error: key cannot be empty")
	}
	fileinfo, err := os.Stat(key)
	if errors.Is(err, os.ErrNotExist) {
		return "", errors.New("error: \"" + key + "\" key no such file or directory")
	}
	if !fileinfo.Mode().IsRegular() {
		return "", errors.New("error: \"" + key + "\" not regular")
	}
	return key, nil
}

// SearchHostToTerminal search host to terminal struct
func SearchHostToTerminal(searchhostmap map[string]string) (*cli.Terminal, error) {
	if len(searchhostmap) <= 0 {
		return nil, fmt.Errorf("%s%s", "error: ", "map to option exception")
	}
	var address, okAddress = searchhostmap[sys.HOST_ADDRESS]
	var username, okUsername = searchhostmap[sys.HOST_USERNAME]
	var port = searchhostmap[sys.HOST_PORT]
	var authtype = searchhostmap[sys.HOST_AUTHTYPE]
	if !okAddress || len(address) <= 0 {
		return nil, fmt.Errorf("%s%s", "error: ", "address cannot be empty")
	}
	if !okUsername || len(username) <= 0 {
		return nil, fmt.Errorf("%s%s", "error: ", "username cannot be empty")
	}
	var iport, err = CheckPortRange(port)
	if err != nil {
		ExitPrintln(fmt.Sprint(err))
	}
	var password, okPassword = searchhostmap[sys.HOST_PASSWORD]
	var key, okKey = searchhostmap[sys.HOST_KEY]
	if !okPassword && !okKey {
		return nil, errors.New("error: required choose one( -x | -k )")
	}
	if len(strings.Trim(password, " ")) > 0 && len(strings.Trim(key, " ")) > 0 {
		return nil, errors.New("error: can only choose one( -x | -k )")
	}
	var auth = ""
	if okPassword && len(strings.Trim(password, " ")) > 0 {
		var password, err = CheckPassword(password)
		if err != nil {
			return nil, err
		}
		auth = password
	}
	if okKey && len(strings.Trim(key, " ")) > 0 {
		var key, err = CheckPublicKey(key)
		if err != nil {
			return nil, err
		}
		auth = key
	}

	var terminal = new(cli.Terminal)
	switch authtype {
	case sys.AUTHTYPE_X:
		terminal = cli.New(address, username, auth, "", iport)
	case sys.AUTHTYPE_K:
		terminal = cli.New(address, username, "", auth, iport)
	default:
		return nil, fmt.Errorf("%s%s", "error: ", "required choose one( -x | -k )")
	}
	return terminal, nil
}

// RenderTable render the table
func RenderTable(searchhostmaps []map[string]string, columns ...string) {
	var valuespaces map[string]int = HostMapsByValueSpace(searchhostmaps, columns)
	var headspaces, headtexts = RenderTableHead(valuespaces, columns...)
	if len(headtexts) > 0 {
		fmt.Println(headtexts)
	}

	var rowtexts = RenderTableRow(searchhostmaps, headspaces, columns...)
	if len(rowtexts) > 0 {
		fmt.Println(rowtexts)
	}
}

// RenderTableRow render the table row
func RenderTableRow(searchhostmaps []map[string]string, headspaces map[string]int, columns ...string) string {
	var rows []string = []string{}
	if len(searchhostmaps) > 0 {
		var searchhostsmapscolumns, err = HostMapsReturnColumns(searchhostmaps, columns...)
		if err != nil {
			ExitPrintln(fmt.Sprintf("error: %v", err))
		}
		for _, searchhostsmapscolumn := range searchhostsmapscolumns {
			var row []string = []string{}
			for _, column := range columns {
				upperletter, _ := UpperCaseFirstLetter(column)
				var value = searchhostsmapscolumn[upperletter]
				var curcolumncount = len(value)
				var count = headspaces[column] - curcolumncount
				if count < 0 {
					count = 0 //only one column
				}
				row = append(row, fmt.Sprintf("%s%s", value, strings.Repeat(" ", count)))
			}
			var rowtext = strings.Join(row, "")
			rows = append(rows, rowtext)
		}
	}
	return strings.Join(rows, "\n")
}

// RenderTableHead render the table header and return the size of each column in the header
func RenderTableHead(valuespaces map[string]int, columns ...string) (map[string]int, string) {
	var heads []string = []string{}
	var headspaces map[string]int = map[string]int{}
	var columnslen = len(columns)
	if columnslen > 0 {
		for i, column := range columns {
			var curcolumncount = len(column)
			var maxcolumncount = valuespaces[column]
			var count = (maxcolumncount - curcolumncount) + 3
			if maxcolumncount <= curcolumncount || count <= 0 {
				count = 3
			}
			if i == (columnslen - 1) {
				count = 0
			}
			var headtext = fmt.Sprintf("%s%s", column, strings.Repeat(" ", count))
			heads = append(heads, headtext)
			headspaces[column] = len(headtext)
		}
	}
	return headspaces, strings.Join(heads, "")
}

// HostMapsByValueSpace calculate the maximum length of the value of the specified key in the set
func HostMapsByValueSpace(searchhostmaps []map[string]string, columns []string) map[string]int {
	var valuespaces map[string]int = map[string]int{}
	if len(columns) > 0 {
		for _, column := range columns {
			var valuespace = 0
			for _, searchhostmap := range searchhostmaps {
				upperletter, _ := UpperCaseFirstLetter(column)
				var value = searchhostmap[upperletter]
				var tempvs = len(value)
				if tempvs > valuespace {
					valuespace = tempvs
				}
			}
			valuespaces[column] = valuespace
		}
		// calc AUTH space
		var valuespace = 0
		for _, searchhostmap := range searchhostmaps {
			var authtype = searchhostmap[sys.HOST_AUTHTYPE]
			var tempvs = 0
			switch authtype {
			case sys.AUTHTYPE_X:
				tempvs = len(searchhostmap[sys.HOST_PASSWORD])
			case sys.AUTHTYPE_K:
				tempvs = len(searchhostmap[sys.HOST_KEY])
			}
			if tempvs > valuespace {
				valuespace = tempvs
			}
			valuespaces[sys.COLUMN_AUTH] = valuespace
		}
	}
	return valuespaces
}

// cmd call system command
func Cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// clear clear the terminal screen
func Clear() {
	switch runtime.GOOS {
	case "darwin":
		Cmd("clear")
	case "linux":
		Cmd("clear")
	case "windows":
		Cmd("cmd", "/c", "cls")
	default:
		Cmd("clear")
	}
}

func PasswordWarning() {
	fmt.Printf("kush: [Warning] Using a password on the command line interface can be insecure.\n")
}
