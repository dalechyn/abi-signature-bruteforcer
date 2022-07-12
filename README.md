# ABISignatureBruteforcer
ETH ABI Signature Bruteforcer

## Build
`go build`

## Usage
```
Usage:
  ABISignatureBruteforcer [flags]

Flags:
  -h, --help            help for ABISignatureBruteforcer
  -s, --name string     name of your function that needs to be bruteforced. use * as placeholder [hello_world_*(uint,address)]
  -t, --target string   target signature, use * as wildcard for any byte. i.e: 0x000000**
```
