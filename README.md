# gcg

Go Constructor Generator

"gcg" generates constructors of structs automatically.
Check [example](./example) directory for more information.

## Install

```sh
$ go install github.com/Ras96/gcg@latest
```

## Usage

```txt
$ gcg
Usage:
  gcg [flags]
  gcg [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  gen         Command "gen" generates constructors
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.gcg.yml)
  -h, --help            help for gcg
  -t, --toggle          Help message for toggle

Use "gcg [command] --help" for more information about a command.
```
