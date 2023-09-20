
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
curl -sS https://raw.githubusercontent.com/anigkus/kush/main/install.sh | bash -s -- 0.0.1
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
  -a, --address         [required] Address 
  -p, --port            [optional] Port, DEFAULT:22
  -u, --username        [optional] Username, DEFAULT:`whoami`
  -x, --password        [required] Password, Can only choose one( -x | -k )
  -k, --key             [required] Key, Can only choose one( -x | -k )
  -t, --title           [optional] Title
  -f, --filter          [optional] Filter( ADDRESS , USERNAME , PORT , AUTHTYPE , GROUP , TITLE )
  -q, --quiet           [optional] Quiet
  -o, --output          [optional] Output, DEFAULT:json
  -c, --columns         [optional] columns( ADDRESS , USERNAME , PORT ,AUTHTYPE , AUTH , GROUP , TITLE )
  -w, --wide            [optional] Wide
  -g, --group           [optional] Group, DEFAULT:default
  -s, --sort            [optional] Sort[ASC]( ADDRESS , USERNAME , PORT , AUTHTYPE , GROUP ), DEFAULT:ADDRESS
  -v, --version         version
  -h, --help            help

Run 'kush COMMAND --help' for more information on a command.

Direct Connection:

kush user@host -x 123456
kush user@host -k ~/.ssh/id_rsa_github.pub

To get more help with kush, check out our guides at https://github.com/anigkus/kush
```

# Special Parameter

## --filter

```
-f 'ADDRESS = 1.1.1.1'
-f 'USERNAME = user'
-f 'PORT = 1'
-f 'AUTHTYPE = [ -x | -k ]'
-f 'GROUP = group'
-f 'TITLE = test'
-f 'ADDRESS = 1.1.1.1 && USERNAME = user'
-f 'ADDRESS = 1.1.1.1 || USERNAME = user'

```

## --columns 

```
-c 'ADDRESS , USERNAME , PORT ,AUTHTYPE , AUTH , GROUP , TITLE'
```

## --sort 

```
-s ADDRESS 
—s USERNAME 
—s PORT 
—s AUTHTYPE 
—s GROUP
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
| 192.168.1.1 | root | 22 | -x | 123456 | group1 | xxx1  |
| 192.168.1.2 | root | 22 | -k | ~/id_rsa.pub | group3  | xxx1  |

```
> 1.2
```

# Direct Connection

```
kush root@192.168.1.10 -x 123456
```

# Create Host

## `create` command format

```
kush create [OPTIONS]
```

## `create` use case

### add a password-based host, interactive mode

```
kush create -a 192.168.1.1 -x 123456
```

### add a password-based host, non-interactive mode

```
kush create -a 192.168.1.1 -x 123456 -q
```

### add a identity-based host, interactive mode

```
kush create -a 192.168.1.1 -k ~/.ssh/id_rsa_github.pub
```

### add a password-based and group name use: default, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -x 123456
```

### add a password-based and group name use: default and title name use: default, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -x 123456 -t test
```

### add a identity-based and group name use: default and title name use: test, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -k ~/.ssh/id_rsa_github.pub -t test
```

### add a password-based and group name use: default and title name use: test, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -x 123456 -t test
```

### add a identity-based and group name use: default and title name use: test, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -k ~/.ssh/id_rsa_github.pub -t test
```

### add a password-based and group name use: Group1 and title name use: test, interactive mode

```
kush create -a 192.168.1.1 -u root -p 22 -x 123456 -t test -g Group1
```

# Remove Host

## `remove` command format

```
kush remove [OPTIONS]
```
## `remove` use case

### `remove` according to the combination of address、username、port, interactive mode

```
kush remove -a 192.168.1.1 -u root -p 22 
```

### `remove` according to the combination of address、username、port, non-interactive mode

```
kush remove -a 192.168.1.1 -u root -p 22 -q
```

### `remove` by [ address ] condition, non-interactive mode

```
kush remove -a 192.168.1.1 -q
```

### `remove` according to the combination of [address] and filter conditions (GROUP=group1), non-interactive mode

```
kush remove -a 192.168.1.1 -f 'GROUP=group1' -q
```

# Update Host

## `update` command format 

```
kush update [OPTIONS]
```
## `update` use case 

### `update` password according to `address`, interactive mode

```
kush update -a 192.168.1.1 -x 123456
```

### `update` identity key according to `address`, interactive mode

```
kush update -a 192.168.1.1 -k ~/.ssh/id_rsa_github.pub
```

### `update` title according to `address`, interactive mode

```
kush update -a 192.168.1.1 -t xxx
```

### `update` username, port according to `address`, interactive mode

```
kush update -a 192.168.1.1 -u root -p 22
```

### `update` username, port, group name according to `address`, interactive mode

```
kush update -a 192.168.1.1 -u root -p 22 -g Group1
```

# Search Host

## `search` command format

```
kush search [OPTIONS]
```

## `search` use case

### `search` data according to `address`

```
kush search -a 192.168.1.1
```

### `search` data according to `address` condition and add filter condition (GROUP=group1), and sort and display according to address

```
kush search -a test.host1.com -f 'GROUP=group1' -s ADDRESS
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
kush export -o json -f 'GROUP=group1' > /tmp/kush.json
```

### `export` json format, specify column

```
kush export -o json -c 'ADDRESS,USERNAME' > ~/.kush/kush.json
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
kush [ -v ｜ --version ]
```

# Help

```
kush | kush -h | kush --help
```