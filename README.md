# go-indexer

[![Tag](https://img.shields.io/github/tag/wealdtech/go-indexer.svg)](https://github.com/wealdtech/go-ecodec/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-indexer.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-indexer?status.svg)](https://godoc.org/github.com/wealdtech/go-ecodec)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-indexer.svg)](https://travis-ci.org/wealdtech/go-ecodec)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-indexer.svg)](https://codecov.io/github/wealdtech/go-ecodec)
[![Go Report Card](https://goreportcard.com/badge/github.com/wealdtech/go-indexer)](https://goreportcard.com/report/github.com/wealdtech/go-ecodec)

Go module providing simple index of strings to UUIDs.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-indexer` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-indexer
```

## Usage

### Example

```go
package main

import (
    "fmt"

    indexer "github.com/wealdtech/go-indexer"
)

func main() {
    indexer := index.New()

}
```
## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-indexer/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
