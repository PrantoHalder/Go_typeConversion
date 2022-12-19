package main 

import "fmt"

func intToFloat(){
	a := 10
	b := 4.5 

	c := float64(a) + b //conversion using float64()
	fmt.Println("Value after converting int to float ",c)
}