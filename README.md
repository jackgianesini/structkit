[![Go](https://github.com/kitstack/structkit/actions/workflows/coverage.yml/badge.svg)](https://github.com/kitstack/structkit/actions/workflows/coverage.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kitstack/structkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/kitstack/structkit)](https://goreportcard.com/report/github.com/kitstack/structkit)
[![codecov](https://codecov.io/gh/kitstack/structkit/branch/main/graph/badge.svg?token=3JRL5ZLSIH)](https://codecov.io/gh/kitstack/structkit)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kitstack/structkit/blob/main/LICENSE)
[![Github tag](https://badgen.net/github/release/kitstack/structkit)](https://github.com/kitstack/structkit/releases)

# 🚀 Overview

StructKit is simple tool for : 

- [x] Copy specific fields a new struct.
  - Tag copy 
- [x] Get value of specific field.
    - Slice
    - Struct
    - Pointer
- [x] Set value of specific field.
    - Slice (Append, Replace or Update)
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
    "github.com/kitstack/structkit"
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
    "github.com/kitstack/structkit"
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


### Set

```go
package main

import (
    "github.com/kitstack/structkit"
    "log"
)

type Foo struct {
    Value   string
    Struct  Bar
    StructP *Bar
    Slice   []Bar
}

type Bar struct {
    Value string
}

func main() {
    payload := Foo{}
  
    err := structkit.Set(&payload, "Value", "foo")
    if err != nil {
      panic(err)
    }
    log.Print(payload.Value) // foo
  
    err = structkit.Set(&payload, "Struct.Value", "bar")
    if err != nil {
      panic(err)
    }
    log.Print(payload.Struct.Value) // bar
  
    err = structkit.Set(&payload, "StructP.Value", "bar")
    if err != nil {
      panic(err)
    }
    log.Print(payload.StructP.Value) // bar
  
    err = structkit.Set(&payload, "Slice.[*]", Bar{Value: "bar"})
    if err != nil {
      panic(err)
    }
    log.Print(payload.Slice) // [{bar}]
  
    err = structkit.Set(&payload, "Slice.[0].Value", "bar updated")
    if err != nil {
      panic(err)
    }
    log.Print(payload.Slice) // [{bar updated}]
}

```


### 💪 Benchmark

```bash
goos: darwin
goarch: arm64
pkg: github.com/kitstack/structkit
BenchmarkGet
BenchmarkGet-10                      	 6346312	       194.0 ns/op
BenchmarkGetEmbeddedValue
BenchmarkGetEmbeddedValue-10         	 5543814	       209.5 ns/op
BenchmarkGetEmbeddedSliceValue
BenchmarkGetEmbeddedSliceValue-10    	 4526838	       262.4 ns/op
```

## 🤝 Contributions
Contributors to the package are encouraged to help improve the code.