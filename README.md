# Big Fat Converter

Converts from one encoding to another.

Supported formats:

- HCL
- JSON
- YAML
- TOML
- XML (only decoding for now)
- plist

## Download

```
https://github.com/hoshsadiq/big-fat-converter/releases
```

## usage

```
Convert from one format to another

Usage:
  bfc [command]

Available Commands:
  convert     Convert formats from an input type to an output type
  help        Help about any command
  version     Print the version number of the Big Fat Converter

Flags:
  -h, --help   help for bfc

Use "bfc [command] --help" for more information about a command.
```

Files can be read from arguments or stdin, however, the result will _always_ be output to stdout.
Use the shell's redirection features to output to a file if needed.


### examples

```
$ bfc convert input.yaml -o json                    # Convert input.yaml to json
$ bfc convert input.hcl -o json                     # Convert input.hcl to json
$ cat input.yaml | bfc convert -i yaml -o json      # Convert stdin (yaml) to json
```
