package main

import (
	"fmt"
	"strconv"
)

func intTostringW1() {
	a := 10
	b := "4.5"

	e := strconv.Itoa(a)
	f := e+b

	fmt.Println("the value after converting int to string way 1", f)
}
