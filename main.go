package main

import (
	"fmt"
	"time"

	"github.com/kr/pretty"
)

type MyStruct struct {
	Str string
	I   int
	B   *bool

	T time.Time
}

func main() {
	for i := 0; i <= 50; i++ {
		var s MyStruct
		err := fakeStruct(&s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%# v\n", pretty.Formatter(s))
	}
}
