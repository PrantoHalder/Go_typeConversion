package main

import (
	"fmt"
	"strconv"
)

func stringToInt () {
	a := "10"
	b := 20

	e,_ := strconv.Atoi(a)
	d := e+b
	fmt.Println("the value after converitng form stirng to int ",d)
}