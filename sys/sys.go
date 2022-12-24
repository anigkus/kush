// Copyright 2022 The https://github.com/anigkus Authors. All rights reserved.
// Use of this source code is governed by a APACHE-style
// license that can be found in the LICENSE file.

package sys

import (
	"io/fs"
	"time"
)

var VERSION string = "1.0.0"

var COMMANDS []string = []string{
	COMMANDS_CREATE,
	COMMANDS_REMOVE,
	COMMANDS_UPDATE,
	COMMANDS_SEARCH,
	COMMANDS_EXPORT,
	COMMANDS_IMPORT,
}

var COMMANDS_CREATE = "create"
var COMMANDS_REMOVE = "remove"
var COMMANDS_UPDATE = "update"
var COMMANDS_SEARCH = "search"
var COMMANDS_EXPORT = "export"
var COMMANDS_IMPORT = "import"

var OPTIONS []string = []string{
	"-A", "--address",
	"-P", "--port",
	"-U", "--username",
	"-X", "--password",
	"-K", "--key",
	"-C", "--comment",
	"-f", "--filter",
	"-Q", "--quiet",
	"-O", "--output",
	"-C", "--columns",
	"-W", "--wide",
	"-G", "--group",
	"-S", "--sort",
	"-V", "--version",
	"-H", "--help"}

var KUSH_HOME_PATH = ".kush"
var KUSH_PATH_JSON = ".kush.json"
var KUSH_HOME_VAR = "KUSH_HOME"

var DEFAULT_PORT int32 = 22
var DEFAULT_OUTPUT string = "json"
var DEFAULT_GROUP string = "default"
var DEFAULT_SORT string = "ADDRESS"

var AUTHTYPE_K string = "-K"
var AUTHTYPE_X string = "-X"

var OPTION_COLUMNS []string = []string{"ADDRESS", "USERNAME", "PORT", "AUTHTYPE", "AUTH", "GROUP", "TITLE"}
var OPTION_FILTERS []string = []string{"ADDRESS", "USERNAME", "PORT", "AUTHTYPE", "GROUP", "TITLE"}
var OPTION_SORTS []string = []string{"ADDRESS", "USERNAME", "PORT", "AUTHTYPE", "GROUP"}
var OPTION_OUTPUT []string = []string{"json"}

var LOGICAL_OPERATORS_AND string = "&&"
var LOGICAL_OPERATORS_OR string = "||"

var TERMINAL_TIMEOUT time.Duration = 30 * time.Second

var OPERATORS_KEY string = "__OP__"

var AT string = "@"

const (
	COLUMN_ADDRESS  = "ADDRESS"
	COLUMN_USERNAME = "USERNAME"
	COLUMN_PORT     = "PORT"
	COLUMN_AUTHTYPE = "AUTHTYPE"
	COLUMN_AUTH     = "AUTH" //
	COLUMN_GROUP    = "GROUP"
	COLUMN_TITLE    = "TITLE"
)

const (
	HOST_ADDRESS  = "Address"
	HOST_PORT     = "Port"
	HOST_USERNAME = "Username"
	HOST_PASSWORD = "Password"
	HOST_KEY      = "Key"
	HOST_TITLE    = "Title"
	HOST_GROUP    = "Group"
	HOST_AUTHTYPE = "Authtype"
)

const (
	OPTION_ENTITY_ADDRESS  = "address"
	OPTION_ENTITY_PORT     = "port"
	OPTION_ENTITY_USERNAME = "username"
	OPTION_ENTITY_PASSWORD = "password"
	OPTION_ENTITY_AUTH     = "key"
	OPTION_ENTITY_TITLE    = "title"
	OPTION_ENTITY_GROUP    = "group"
)

const (
	// -A, --address
	// -P, --port
	// -U, --username
	// -X, --password
	// -K, --key
	// -T, --title
	// -F, --filter
	// -Q, --quiet
	// -O, --output
	// -C, --columns
	// -W, --wide
	// -G, --group
	// -S, --sort
	// -V, --version
	// -H, --help
	OPTION_S_ADDRESS  = "-A"
	OPTION_L_ADDRESS  = "--address"
	OPTION_S_PORT     = "-P"
	OPTION_L_PORT     = "--port"
	OPTION_S_USERNAME = "-U"
	OPTION_L_USERNAME = "--username"
	OPTION_S_PASSWORD = "-X"
	OPTION_L_PASSWORD = "--password"
	OPTION_S_KEY      = "-K"
	OPTION_L_KEY      = "--key"
	OPTION_S_TITLE    = "-T"
	OPTION_L_TITLE    = "--title"
	OPTION_S_FILTER   = "-F"
	OPTION_L_FILTER   = "--filter"
	OPTION_S_QUIET    = "-Q"
	OPTION_L_QUIET    = "--quiet"
	OPTION_S_OUTPUT   = "-O"
	OPTION_L_OUTPUT   = "--output"
	OPTION_S_COLUMNS  = "-C"
	OPTION_L_COLUMNS  = "--columns"
	OPTION_S_WIDE     = "-W"
	OPTION_L_WIDE     = "--wide"
	OPTION_S_GROUP    = "-G"
	OPTION_L_GROUP    = "--group"
	OPTION_S_SORT     = "-S"
	OPTION_L_SORT     = "--sort"
	OPTION_S_VERSION  = "-V"
	OPTION_L_VERSION  = "--version"
	OPTION_S_HELP     = "-H"
	OPTION_L_HELP     = "--help"
)

// {"-A", "-P", "-U", "-X", "-K", "-T", "-G", "-Q", "-H"}
var CREATE_OPTION []string = []string{OPTION_S_ADDRESS, OPTION_S_PORT, OPTION_S_USERNAME, OPTION_S_PASSWORD, OPTION_S_KEY, OPTION_S_TITLE, OPTION_S_GROUP, OPTION_S_QUIET, OPTION_S_HELP}

// {"-A", "-P", "-U", "-T", "-G", "-F", "-Q", "-H"}
var REMOVE_OPTION []string = []string{OPTION_S_ADDRESS, OPTION_S_PORT, OPTION_S_USERNAME, OPTION_S_TITLE, OPTION_S_GROUP, OPTION_S_FILTER, OPTION_S_QUIET, OPTION_S_HELP}

// {"-A", "-P", "-U", "-X", "-K", "-T", "-G", "-Q", "-H"}
var UPDATE_OPTION []string = []string{OPTION_S_ADDRESS, OPTION_S_PORT, OPTION_S_USERNAME, OPTION_S_PASSWORD, OPTION_S_KEY, OPTION_S_TITLE, OPTION_S_GROUP, OPTION_S_QUIET, OPTION_S_HELP}

// {"-A", "-P", "-U", "-T", "-G", "-F", "-C", "-S", "-H"}
var SEARCH_OPTION []string = []string{OPTION_S_ADDRESS, OPTION_S_PORT, OPTION_S_USERNAME, OPTION_S_TITLE, OPTION_S_GROUP, OPTION_S_FILTER, OPTION_S_COLUMNS, OPTION_S_SORT, OPTION_S_HELP}

// {"-A", "-P", "-U", "-T", "-G", "-F", "-C", "-O", "-H"}
var EXPORT_OPTION []string = []string{OPTION_S_ADDRESS, OPTION_S_PORT, OPTION_S_USERNAME, OPTION_S_TITLE, OPTION_S_GROUP, OPTION_S_FILTER, OPTION_S_COLUMNS, OPTION_S_OUTPUT, OPTION_S_HELP}

// {"-O", "-H"}
var IMPORT_OPTION []string = []string{OPTION_S_OUTPUT, OPTION_S_HELP}

var PERM_FILE_MODE fs.FileMode = 0777
