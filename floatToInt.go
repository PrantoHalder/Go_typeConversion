package main 

import "fmt"

func floatToInt () {
	a := 10
	b := 4.5
	c := a+int(b)
	fmt.Println("Value after changing from Float to int",c)
}