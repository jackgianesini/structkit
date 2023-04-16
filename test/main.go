package main

import (
	"github.com/lab210-dev/structkit"
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
