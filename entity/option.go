// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package entity

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/anigkus/kush/sys"
	"github.com/anigkus/kush/util"
)

type Option struct {
	Address  string
	Port     int32
	Username string
	Password string
	Key      string
	Title    string
	Filter   map[string]string
	Quiet    bool
	Output   string
	Columns  []string
	Wide     bool
	Group    string
	Sort     string
}

// no parameter constructor: public and instanced function
func OptionDefault() *Option {
	return &Option{
		Output: sys.DEFAULT_OUTPUT,
		Port:   sys.DEFAULT_PORT,
		Group:  sys.DEFAULT_GROUP,
		Sort:   sys.DEFAULT_SORT}
}

// with parameter constructor: public and instanced function
func OptionInstance(
	address string,
	port int32,
	username string,
	password string,
	key string,
	title string,
	filter map[string]string,
	quiet bool,
	output string,
	columns []string,
	wide bool,
	group string,
	sort string) *Option {

	return &Option{
		Address:  address,
		Port:     port,
		Username: username,
		Password: password,
		Key:      key,
		Title:    title,
		Filter:   filter,
		Quiet:    quiet,
		Output:   output,
		Columns:  columns,
		Wide:     wide,
		Group:    group,
		Sort:     sort}

}

// String struct to string
func (option Option) String() string {
	valueof := reflect.ValueOf(option)
	vtype := valueof.Type()
	var elements []string
	for i := 0; i < valueof.NumField(); i++ {
		var element = fmt.Sprintln(vtype.Field(i).Name, "=", valueof.Field(i).Interface())
		elements = append(elements, element)
	}
	//return util.StringArrayToString(elements)
	return strings.Join(elements, "")
}

// get and set
func (option *Option) GetAddress() string {
	return option.Address
}
func (option *Option) SetAddress(address string) {
	option.Address = address
}

func (option *Option) GetPort() int32 {
	return option.Port
}
func (option *Option) SetPort(port int32) {
	option.Port = port
}

func (option *Option) GetUsername() string {
	return option.Username
}
func (option *Option) SetUsername(username string) {
	option.Username = username
}

func (option *Option) GetPassword() string {
	return option.Password
}
func (option *Option) SetPassword(password string) {
	option.Password = password
}

func (option *Option) Getkey() string {
	return option.Key
}
func (option *Option) SetKey(key string) {
	option.Key = key
}

func (option *Option) GetTitle() string {
	return option.Title
}
func (option *Option) SetTitle(title string) {
	option.Title = title
}

func (option *Option) GetFilter() map[string]string {
	return option.Filter
}
func (option *Option) SetFilter(filter map[string]string) {
	option.Filter = filter
}

func (option *Option) GetQuiet() bool {
	return option.Quiet
}
func (option *Option) SetQuiet(quiet bool) {
	option.Quiet = quiet
}

func (option *Option) GetOutput() string {
	return option.Output
}
func (option *Option) SetOutput(output string) {
	option.Output = output
}

func (option *Option) GetColumns() []string {
	return option.Columns
}
func (option *Option) SetColumns(columns []string) {
	option.Columns = columns
}

func (option *Option) GetWide() bool {
	return option.Wide
}
func (option *Option) SetWide(wide bool) {
	option.Wide = wide
}

func (option *Option) GetGroup() string {
	return option.Group
}
func (option *Option) SetGroup(group string) {
	option.Group = group
}

func (option *Option) GetSort() string {
	return option.Sort
}
func (option *Option) SetSort(sort string) {
	option.Sort = sort
}

// ParseArgOptions convert terminal option arguments to map[string]string
//
// If both -a and --address exist, -a is Prefered.
func ParseArgOptions(argOptionsMap map[string]string) (map[string]string, error) {
	if argOptionsMap == nil || len(argOptionsMap) <= 0 {
		return nil, errors.New("error: args empty")
	}
	var optionMaps map[string]string = map[string]string{}
	//Preferred:-a
	var argAddressShortHand, okAddressShortHand = argOptionsMap[sys.OPTION_S_ADDRESS]
	var argAddressOption, okAddressOption = argOptionsMap[sys.OPTION_L_ADDRESS]
	if okAddressShortHand {
		optionMaps[sys.OPTION_S_ADDRESS] = strings.Trim(argAddressShortHand, " ")
	}
	if okAddressOption {
		optionMaps[sys.OPTION_S_ADDRESS] = argAddressOption
	}
	// if err := util.ManyIsEmpty(argAddressShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_ADDRESS] = strings.Trim(argAddressShortHand, " ")
	// } else if err := util.ManyIsEmpty(argAddressOption); err == nil {
	// 	optionMaps[sys.OPTION_S_ADDRESS] = argAddressOption
	// }

	//Preferred:-p
	var argPortShortHand, okPortShortHand = argOptionsMap[sys.OPTION_S_PORT]
	var argPortOption, okPortOption = argOptionsMap[sys.OPTION_L_PORT]
	if okPortShortHand {
		optionMaps[sys.OPTION_S_PORT] = argPortShortHand
	}
	if okPortOption {
		optionMaps[sys.OPTION_S_PORT] = argPortOption
	}
	// if err := util.ManyIsEmpty(argPortShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_PORT] = argPortShortHand
	// } else if err := util.ManyIsEmpty(argPortOption); err == nil {
	// 	optionMaps[sys.OPTION_S_PORT] = argPortOption
	// }

	//Preferred:-u
	var argUsernameShortHand, okUsernameShortHand = argOptionsMap[sys.OPTION_S_USERNAME]
	var argUsernameOption, okUsernameOption = argOptionsMap[sys.OPTION_L_USERNAME]
	if okUsernameShortHand {
		optionMaps[sys.OPTION_S_USERNAME] = argUsernameShortHand
	}
	if okUsernameOption {
		optionMaps[sys.OPTION_S_USERNAME] = argUsernameOption
	}
	// if err := util.ManyIsEmpty(argUsernameShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_USERNAME] = argUsernameShortHand
	// } else if err := util.ManyIsEmpty(argUsernameOption); err == nil {
	// 	optionMaps[sys.OPTION_S_USERNAME] = argUsernameOption
	// }
	//Preferred:-x
	var argPasswordShortHand, okPasswordShortHand = argOptionsMap[sys.OPTION_S_PASSWORD]
	var argPasswordOption, okPasswordOption = argOptionsMap[sys.OPTION_L_PASSWORD]
	if okPasswordShortHand {
		optionMaps[sys.OPTION_S_PASSWORD] = argPasswordShortHand
	}
	if okPasswordOption {
		optionMaps[sys.OPTION_S_PASSWORD] = argPasswordOption
	}
	// if err := util.ManyIsEmpty(argPasswordShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_PASSWORD] = argPasswordShortHand
	// } else if err := util.ManyIsEmpty(argPasswordOption); err == nil {
	// 	optionMaps[sys.OPTION_S_PASSWORD] = argPasswordOption
	// }

	//Preferred:-k
	var argKeyShortHand, okKeyShortHand = argOptionsMap[sys.OPTION_S_KEY]
	var argKeyOption, okKeyOption = argOptionsMap[sys.OPTION_L_KEY]
	if okKeyShortHand {
		optionMaps[sys.OPTION_S_KEY] = argKeyShortHand
	}
	if okKeyOption {
		optionMaps[sys.OPTION_S_KEY] = argKeyOption
	}
	// if err := util.ManyIsEmpty(argKeyShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_KEY] = argKeyShortHand
	// } else if err := util.ManyIsEmpty(argKeyOption); err == nil {
	// 	optionMaps[sys.OPTION_S_KEY] = argKeyOption
	// }

	//Preferred:-t
	var argTitleShortHand, okTitleShortHand = argOptionsMap[sys.OPTION_S_TITLE]
	var argTitleOption, okTitleOption = argOptionsMap[sys.OPTION_L_TITLE]
	if okTitleShortHand {
		optionMaps[sys.OPTION_S_TITLE] = argTitleShortHand
	}
	if okTitleOption {
		optionMaps[sys.OPTION_S_TITLE] = argTitleOption
	}
	// if err := util.ManyIsEmpty(argTitleShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_TITLE] = argTitleShortHand
	// } else if err := util.ManyIsEmpty(argTitleOption); err == nil {
	// 	optionMaps[sys.OPTION_S_TITLE] = argTitleOption
	// }

	//Preferred:-f
	var argFilterShortHand, okFilterShortHand = argOptionsMap[sys.OPTION_S_FILTER]
	var argFilterOption, okFilterOption = argOptionsMap[sys.OPTION_L_FILTER]
	if okFilterShortHand {
		optionMaps[sys.OPTION_S_FILTER] = argFilterShortHand
	}
	if okFilterOption {
		optionMaps[sys.OPTION_S_FILTER] = argFilterOption
	}
	// if err := util.ManyIsEmpty(argFilterShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_FILTER] = argFilterShortHand
	// } else if err := util.ManyIsEmpty(argFilterOption); err == nil {
	// 	optionMaps[sys.OPTION_S_FILTER] = argFilterOption
	// }

	//Preferred:-q
	var _, okQuietShortHand = argOptionsMap[sys.OPTION_S_QUIET]
	var _, okQuietOption = argOptionsMap[sys.OPTION_L_QUIET]

	if okQuietShortHand || okQuietOption {
		optionMaps[sys.OPTION_S_QUIET] = "true"
	}

	//Preferred:-o
	var argOutputShortHand, okOutputShortHand = argOptionsMap[sys.OPTION_S_OUTPUT]
	var argOutputOption, okOutputOption = argOptionsMap[sys.OPTION_L_OUTPUT]
	if okOutputShortHand {
		optionMaps[sys.OPTION_S_OUTPUT] = argOutputShortHand
	}
	if okOutputOption {
		optionMaps[sys.OPTION_S_OUTPUT] = argOutputOption
	}
	// if err := util.ManyIsEmpty(argOutputShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_OUTPUT] = argOutputShortHand
	// } else if err := util.ManyIsEmpty(argOutputOption); err == nil {
	// 	optionMaps[sys.OPTION_S_OUTPUT] = argOutputOption
	// }

	//Preferred:-c
	var argColumnsShortHand, okColumnsShortHand = argOptionsMap[sys.OPTION_S_COLUMNS]
	var argColumnsOption, okColumnsOption = argOptionsMap[sys.OPTION_L_COLUMNS]
	if okColumnsShortHand {
		optionMaps[sys.OPTION_S_COLUMNS] = argColumnsShortHand
	}
	if okColumnsOption {
		optionMaps[sys.OPTION_S_COLUMNS] = argColumnsOption
	}
	// if err := util.ManyIsEmpty(argColumnsShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_COLUMNS] = argColumnsShortHand
	// } else if err := util.ManyIsEmpty(argColumnsOption); err == nil {
	// 	optionMaps[sys.OPTION_S_COLUMNS] = argColumnsOption
	// }

	//Preferred:-w
	var _, okWideShortHand = argOptionsMap[sys.OPTION_S_WIDE]
	var _, okWideOption = argOptionsMap[sys.OPTION_L_WIDE]

	if okWideShortHand || okWideOption {
		optionMaps[sys.OPTION_S_WIDE] = "true"
	}

	//Preferred:-g
	var argGroupShortHand, okGroupShortHand = argOptionsMap[sys.OPTION_S_GROUP]
	var argGroupOption, okGroupOption = argOptionsMap[sys.OPTION_L_GROUP]
	if okGroupShortHand {
		optionMaps[sys.OPTION_S_GROUP] = argGroupShortHand
	}
	if okGroupOption {
		optionMaps[sys.OPTION_S_GROUP] = argGroupOption
	}
	// if err := util.ManyIsEmpty(argGroupShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_GROUP] = argGroupShortHand
	// } else if err := util.ManyIsEmpty(argGroupOption); err == nil {
	// 	optionMaps[sys.OPTION_S_GROUP] = argGroupOption
	// }

	//Preferred:-s
	var argSortShortHand, okSortShortHand = argOptionsMap[sys.OPTION_S_SORT]
	var argSortOption, okSortOption = argOptionsMap[sys.OPTION_L_SORT]
	if okSortShortHand {
		optionMaps[sys.OPTION_S_SORT] = argSortShortHand
	}
	if okSortOption {
		optionMaps[sys.OPTION_S_SORT] = argSortOption
	}
	// if err := util.ManyIsEmpty(argSortShortHand); err == nil {
	// 	optionMaps[sys.OPTION_S_SORT] = argSortShortHand
	// } else if err := util.ManyIsEmpty(argSortOption); err == nil {
	// 	optionMaps[sys.OPTION_S_SORT] = argSortOption
	// }
	return optionMaps, nil
}

// GetOptionMapsByPort port arguments validata, return port
func GetOptionMapsByPort(port string) (int32, error) {
	return util.CheckPortRange(port)
}

// GetOptionMapsByAddress address arguments validata, return address
func GetOptionMapsByAddress(address string) (string, error) {
	return util.CheckAddress(address)
}

// GetOptionMapsByUsername username arguments validata, return username
func GetOptionMapsByUsername(username string) (string, error) {
	return util.CheckUsername(username)
}

// GetOptionMapsByPassword password arguments validata, return password
func GetOptionMapsByPassword(password string) (string, error) {
	return util.CheckPassword(password)
}

// GetOptionMapsByKey key arguments validata, return key
func GetOptionMapsByKey(key string) (string, error) {
	return util.CheckPublicKey(key)
}

// GetOptionMapsByTitle title arguments validata, return title
func GetOptionMapsByTitle(title string) (string, error) {
	var err = util.ManyIsEmpty(title)
	if err != nil {
		return "", errors.New("error: title cannot be empty")
	}
	return title, nil
}

// GetOptionMapsByFilter filter arguments validata, return filter
func GetOptionMapsByFilter(filter string) (map[string]string, error) {
	return util.ArgFilterToMap(filter)
}

// GetOptionMapsByOutput output arguments validata, return output
func GetOptionMapsByOutput(output string) (string, error) {
	var err = util.ManyIsEmpty(output)
	if err != nil {
		return "", errors.New("error: output cannot be empty")
	}
	if !util.InArray(output, sys.OPTION_OUTPUT) {
		return "", errors.New("error: \"" + output + "\" output value illegal")
	}
	return output, nil
}

// GetOptionMapsByGroup group arguments validata, return group
func GetOptionMapsByGroup(group string) (string, error) {
	var err = util.ManyIsEmpty(group)
	if err != nil {
		return "", errors.New("error: group cannot be empty")
	}
	return group, nil
}

// GetOptionMapsBySort sort arguments validata, return sort
func GetOptionMapsBySort(sort string) (string, error) {
	var err = util.ManyIsEmpty(sort)
	if err != nil {
		return "", errors.New("error: sort cannot be empty")
	}
	if !util.InArray(sort, sys.OPTION_SORTS) {
		return "", errors.New("error: \"" + sort + "\" sort value illegal")
	}
	return sort, nil
}

// GetOptionMapsByColumns columns arguments validata, return columns
func GetOptionMapsByColumns(columns string) ([]string, error) {
	var err = util.ManyIsEmpty(columns)
	if err != nil {
		return nil, errors.New("error: columns cannot be empty")
	}
	var errs []string
	var elements []string = []string{}
	for _, element := range strings.Split(columns, ",") {
		element = strings.Trim(element, " ")
		if !util.InArray(element, sys.OPTION_COLUMNS) {
			errs = append(errs, element)
		} else {
			elements = append(elements, element)
		}
	}
	if len(errs) > 0 {
		return nil, errors.New("error: \"" + util.StringArrayToString(errs) + "\" value illegal")
	}
	return elements, nil
}

// OptionMapsToOption arguments Map to Option struct, only enter this parameter to verify
func OptionMapsToOption(optionMaps map[string]string) (Option, error) {
	var option = Option{}

	var mAdddress = optionMaps[sys.OPTION_S_ADDRESS]
	var address, err = GetOptionMapsByAddress(mAdddress)
	if err != nil {
		return Option{}, err
	} else {
		option.SetAddress(address)
	}

	var mPort, okPort = optionMaps[sys.OPTION_S_PORT]
	if okPort {
		var port, err = GetOptionMapsByPort(mPort)
		if err != nil {
			return Option{}, err
		}
		option.SetPort(port)
	} else {
		option.SetPort(sys.DEFAULT_PORT)
	}

	var mUsername, okUsername = optionMaps[sys.OPTION_S_USERNAME]
	if okUsername {
		var username, err = GetOptionMapsByUsername(mUsername)
		if err != nil {
			return Option{}, err
		}
		option.SetUsername(username)
	} else {
		var currentuser, err = util.CurrentUser()
		if err != nil {
			return Option{}, errors.New("error: get current user exception")
		}
		option.SetUsername(currentuser)
	}
	var mPassword, okPassword = optionMaps[sys.OPTION_S_PASSWORD]
	var mKey, okKey = optionMaps[sys.OPTION_S_KEY]
	if !okPassword && !okKey {
		return Option{}, errors.New("error: required choose one( -X | -K )")
	}
	if okPassword && okKey {
		return Option{}, errors.New("error: can only choose one( -X | -K )")
	}
	if okPassword {
		var password, err = GetOptionMapsByPassword(mPassword)
		if err != nil {
			return Option{}, err
		}
		option.SetPassword(password)
	}
	if okKey {
		var key, err = GetOptionMapsByKey(mKey)
		if err != nil {
			return Option{}, err
		}
		option.SetKey(key)
	}

	var mTitle, okTitle = optionMaps[sys.OPTION_S_TITLE]
	if okTitle {
		var title, err = GetOptionMapsByTitle(mTitle)
		if err != nil {
			return Option{}, err
		}
		option.SetTitle(title)
	} else {
		option.SetTitle("")
	}

	var mFilter, okFilter = optionMaps[sys.OPTION_S_FILTER]
	if okFilter {
		var filter, err = GetOptionMapsByFilter(mFilter)
		if err != nil {
			return Option{}, err
		}
		option.SetFilter(filter)
	}

	var _, okQuiet = optionMaps[sys.OPTION_S_QUIET]
	if okQuiet {
		option.SetQuiet(okQuiet)
	}

	var mOutput, okOutput = optionMaps[sys.OPTION_S_OUTPUT]
	if okOutput {
		var output, err = GetOptionMapsByOutput(mOutput)
		if err != nil {
			return Option{}, err
		}
		option.SetOutput(output)
	} else {
		option.SetOutput(sys.DEFAULT_OUTPUT)
	}

	var mColumns, okColumns = optionMaps[sys.OPTION_S_COLUMNS]
	if okColumns {
		var columns, err = GetOptionMapsByColumns(mColumns)
		if err != nil {
			return Option{}, err
		}
		option.SetColumns(columns)
	}

	var _, okWide = optionMaps[sys.OPTION_S_WIDE]
	if okWide {
		option.SetWide(okWide)
	}

	var mGroup, okGroup = optionMaps[sys.OPTION_S_GROUP]
	if okGroup {
		var group, err = GetOptionMapsByGroup(mGroup)
		if err != nil {
			return Option{}, err
		}
		option.SetGroup(group)
	} else {
		option.SetGroup(sys.DEFAULT_GROUP)
	}

	var mSort, okSort = optionMaps[sys.OPTION_S_SORT]
	if okSort {
		var sort, err = GetOptionMapsBySort(mSort)
		if err != nil {
			return Option{}, err
		}
		option.SetSort(sort)
	} else {
		option.SetSort(sys.DEFAULT_SORT)
	}
	return option, nil
}
