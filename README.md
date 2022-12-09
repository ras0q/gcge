# gcg

Go Constructor Generator

## About

"gcg" generates constructors automatically.
Please check [example](./example) directory for more information.

## Installation

```sh
$ go install github.com/ras0q/gcg@latest
```

## Usage

```txt
$ gcg -h
Usage:
  gcg [flags]
  gcg [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  gen         Generate constructors
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.gcg.yml)
  -h, --help            help for gcg
  -v, --version         version for gcg

Use "gcg [command] --help" for more information about a command.
```

```txt
$ gcg gen -h
Generate constructors

Usage:
  gcg gen [flags]

Flags:
  -h, --help            help for gen
  -o, --output string   Output file
  -p, --private         Generate private constructors

Global Flags:
      --config string   config file (default is $HOME/.gcg.yml)
```
