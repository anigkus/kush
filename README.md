
# Kush README

Kush is a cross-platform command-line SSH remote address connection tool.

# Installation

Kush is available on Linux, macOS and Windows platforms.

* Binaries for Linux, Windows and Mac are available as tarballs in the [release](https://github.com/anigkus/kush/releases) page.

* Via Bash for Linux and macOS
```
curl -sS https://raw.githubusercontent.com/anigkus/kush/main/install.sh | bash
```

OR

```
curl -sS https://raw.githubusercontent.com/anigkus/kush/main/install.sh | bash -s -- 0.0.9
```

# Command Help

```
Usage:  kush COMMAND [OPTIONS]

COMMAND:
  create    Create the kush for host
  remove    Remove the kush for host
  update    Update the kush for host
  search    Search the kush for host
  export    Export the kush to  file
  import    Import the kush for file
 
OPTIONS:
  -A, --address         [required] Address 
  -P, --port            [optional] Port, DEFAULT:22
  -U, --username        [optional] Username, DEFAULT:`whoami`}
  -X, --password        [required] Password, Can only choose one( -X | -K )
  -K, --key             [required] Key, Can only choose one( -X | -K )
  -T, --title           [optional] Title
  -F, --filter          [optional] Filter( ADDRESS , USERNAME , PORT , AUTHTYPE , GROUP , TITLE )
  -Q, --quiet           [optional] Quiet
  -O, --output          [optional] Output, DEFAULT:json
  -C, --columns         [optional] columns( ADDRESS , USERNAME , PORT ,AUTHTYPE , AUTH , GROUP , TITLE )
  -W, --wide            [optional] Wide
  -G, --group           [optional] Group, DEFAULT:default
  -S, --sort            [optional] Sort[ASC]( ADDRESS , USERNAME , PORT , AUTHTYPE , GROUP ), DEFAULT:ADDRESS
  -V, --version         version
  -H, --help            help

Run 'kush COMMAND --help' for more information on a command.

To get more help with kush, check out our guides at https://kush.github.com/go/guides/
```

# Special Parameter

## --filter

```
-F 'ADDRESS = 1.1.1.1'
-F 'USERNAME = user'
-F 'PORT = 1'
-F 'AUTHTYPE = [ -X | -K ]'
-F 'GROUP = group'
-F 'TITLE = test'
-F 'ADDRESS = 1.1.1.1 && USERNAME = user'
-F 'ADDRESS = 1.1.1.1 || USERNAME = user'

```

## --columns 

```
-C 'ADDRESS , USERNAME , PORT ,AUTHTYPE , AUTH , GROUP , TITLE'
```

## --sort 

```
-S ADDRESS 
—S USERNAME 
—S PORT 
—S AUTHTYPE 
—S GROUP
```

# Data Path

## Default

```
~/.kush/.kush.json
```

## Environment

`~/.bash_profile `,Append below:

```
#kush
KUSH_HOME="/home/root"
export PATH=$KUSH_HOME/bin:$PATH
```

# Data Format

## Default

| ADDRESS |
| :--- |
| 192.168.1.1 |
| 192.168.1.2 |

```
> 1.2
```

## Full Table

| ADDRESS | USERNAME | PORT | AUTHTYPE | AUTH | GROUP | TITLE |
| :--- | :---  | :---  | :--- | :--- | :---  | :---  |
| 192.168.1.1 | root | 22 | -X | 123456 | group1 | xxx1  |
| 192.168.1.2 | root | 22 | -K | ~/id_rsa.pub | group3  | xxx1  |

```
> 1.2
```

# Direct Connection

```
kush root@192.168.1.10 -X 123456
```

# Create Host

## `create` command format

```
kush create [OPTIONS]
```

## `create` use case

### add a password-based host, interactive mode

```
kush create -A 192.168.1.1 -X 123456
```

### add a password-based host, non-interactive mode

```
kush create -A 192.168.1.1 -X 123456 -Q
```

### add a identity-based host, interactive mode

```
kush create -A 192.168.1.1 -K ~/.ssh/id_rsa_github.pub
```

### add a password-based and group name use: default, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -X 123456
```

### add a password-based and group name use: default and title name use: default, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -X 123456 -T test
```

### add a identity-based and group name use: default and title name use: test, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -K ~/.ssh/id_rsa_github.pub -T test
```

### add a password-based and group name use: default and title name use: test, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -X 123456 -T test
```

### add a identity-based and group name use: default and title name use: test, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -K ~/.ssh/id_rsa_github.pub -T test
```

### add a password-based and group name use: Group1 and title name use: test, interactive mode

```
kush create -A 192.168.1.1 -U root -P 22 -X 123456 -T test -G Group1
```

# Remove Host

## `remove` command format

```
kush remove [OPTIONS]
```
## `remove` use case

### `remove` according to the combination of address、username、port, interactive mode

```
kush remove -A 192.168.1.1 -U root -P 22 
```

### `remove` according to the combination of address、username、port, non-interactive mode

```
kush remove -A 192.168.1.1 -U root -P 22 -Q
```

### `remove` by [ address ] condition, non-interactive mode

```
kush remove -A 192.168.1.1 -Q
```

### `remove` according to the combination of [address] and filter conditions (GROUP=group1), non-interactive mode

```
kush remove -A 192.168.1.1 -F 'GROUP=group1' -Q
```

# Update Host

## `update` command format 

```
kush update [OPTIONS]
```
## `update` use case 

### `update` password according to `address`, interactive mode

```
kush update -A 192.168.1.1 -X 123456
```

### `update` identity key according to `address`, interactive mode

```
kush update -A 192.168.1.1 -K ~/.ssh/id_rsa_github.pub
```

### `update` title according to `address`, interactive mode

```
kush update -A 192.168.1.1 -T xxx
```

### `update` username, port according to `address`, interactive mode

```
kush update -A 192.168.1.1 -U root -P 22
```

### `update` username, port, group name according to `address`, interactive mode

```
kush update -A 192.168.1.1 -U root -P 22 -G Group1
```

# Search Host

## `search` command format

```
kush search [OPTIONS]
```

## `search` use case

### `search` data according to `address`

```
kush search -A 192.168.1.1
```

### `search` data according to `address` condition and add filter condition (GROUP=group1), and sort and display according to address

```
kush search -A test.host1.com -F 'GROUP=group1' -S ADDRESS
```

# Export Host

## `export` command format

```
kush export [OPTIONS]
```

## `export` use case 

### `export` json format, default column

```
kush export -o json > /tmp/kush.json
```

### `export` json format, specify(GROUP=group1)

```
kush export -o json -F 'GROUP=group1' > /tmp/kush.json
```

### `export` json format, specify column

```
kush export -o json -C 'ADDRESS,USERNAME' > ~/.kush/kush.json
```

# Import Host

## `import` host command format

```
kush import [OPTIONS]
```

## `import` use case 

### import json format

```
kush import -o json < ~/.kush/kush.json
```

# Version

```
kush [ -V ｜ --version ]
```

# Help

```
kush | kush -H | kush --help
```