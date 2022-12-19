package main

import "fmt"

func intTostringW2(){
	a := 10
	b := 5

	c := fmt.Sprintf("%d",a) // assign the value as string
	d := fmt.Sprintf("%d",b) // assign the value as string

	e := c+d

	fmt.Println("the value after converting int to string  way2 ",e)

}
