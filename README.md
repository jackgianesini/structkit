[![Go](https://github.com/lab210-dev/structkit/actions/workflows/go.yml/badge.svg)](https://github.com/lab210-dev/structkit/actions/workflows/coverage.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lab210-dev/structkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/lab210-dev/structkit)](https://goreportcard.com/report/github.com/lab210-dev/structkit)
[![codecov](https://codecov.io/gh/lab210-dev/structkit/branch/main/graph/badge.svg?token=3JRL5ZLSIH)](https://codecov.io/gh/lab210-dev/structkit)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lab210-dev/structkit/blob/main/LICENSE)
[![Github tag](https://badgen.net/github/release/lab210-dev/structkit)](https://github.com/lab210-dev/structkit/releases)

# 🚀 Overview

StructKit is simple tool for : 

- [x] Copy specific fields a new struct.
  - Tag copy 
- [x] Get value of specific field.
    - Slice
    - Struct
    - Pointer
- [x] Coverage
    - 100% tested code 🤓
- [x] Benchmark
    - Get (Optimized with cache)


### Copy

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
    log.Printf("%v", structkit.Copy(payload, "Value", "Struct.Value")) // {foo {bar}}
}
```

### Get

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
    payload := Foo{Value: "foo", Struct: Bar{Value: "bar"}, Slice: []Bar{{Value: "baz"}}}
	
    log.Printf("%v", structkit.Get(payload, "Value")) // foo
    log.Printf("%v", structkit.Get(payload, "Struct/Value", &structkit.Option{Delimiter: "/"})) // bar
    log.Printf("%v", structkit.Get(payload, "Slice.[0].Value")) // baz
}
```

### 💪Benchmark

```bash
goos: darwin
goarch: arm64
pkg: github.com/lab210-dev/structkit
BenchmarkGet
BenchmarkGet-10                      	 6346312	       194.0 ns/op
BenchmarkGetEmbeddedValue
BenchmarkGetEmbeddedValue-10         	 5543814	       209.5 ns/op
BenchmarkGetEmbeddedSliceValue
BenchmarkGetEmbeddedSliceValue-10    	 4526838	       262.4 ns/op
```

## 🤝 Contributions
Contributors to the package are encouraged to help improve the code.