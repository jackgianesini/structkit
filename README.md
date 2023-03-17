[![Go](https://github.com/lab210-dev/structator/actions/workflows/go.yml/badge.svg)](https://github.com/lab210-dev/structator/actions/workflows/go.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lab210-dev/structator)
[![Go Report Card](https://goreportcard.com/badge/github.com/lab210-dev/structator)](https://goreportcard.com/report/github.com/lab210-dev/structator)
[![codecov](https://codecov.io/gh/lab210-dev/structator/branch/main/graph/badge.svg?token=3JRL5ZLSIH)](https://codecov.io/gh/lab210-dev/structator)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lab210-dev/structator/blob/main/LICENSE)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://GitHub.com/Naereen/ama)
[![Github tag](https://badgen.net/github/release/lab210-dev/structator)](https://github.com/lab210-dev/structator/releases)

# Overview

StructKit is simple tool for copy specific fields a new struct, including metadata.

## Features

- [x] Simple
- [x] Coverage
- [x] Compatible with slice
- [x] Tag copy

### Usage

```go
package main

import (
    "github.com/lab210-dev/structkit"
    "log"
)

type Foo struct {
    Counter int    `json:"int"`
    Value   string `json:"value"`
    Struct  Bar    `json:"struct"`
    Slice   []Bar  `json:"Slice"`
}

type Bar struct {
    Value string `json:"value"`
}

func main() {
    payload := Foo{Value: "foo", Struct: Bar{Value: "bar"}}
    log.Printf("%v", structkit.From(payload, "Value", "Struct.Bar"))
}
```

## 🤝 Contributions
Contributors to the package are encouraged to help improve the code.