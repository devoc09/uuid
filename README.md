# UUID

A command line tool for generating uuid to Stdout.

## Installation

-	Build from this Sources.

Author's env `go version go1.17 darwin/amd64`

## How to use

```
$ uuid
```

if you use macOS

```
$ uuid | pbcopy
```

```
uuid is a tool to generate uuid.
Usage: 
    uuid [arguments]

Args:
    -h      Print Help message
    -v      Print the version of this tool
```

## Vim plugin

This tool is also available as a vim plugin.

provides command and keymap(normal-mode) for generating UUID.

-	command

```
:UUID
```

-	keymap(normal-mode)

```
<C-S-u>
```

these put the UUID after the cursor.
