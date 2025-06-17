package main

import "fmt"

// :=  is used to declare and initialize a variable

func main() {
	x := 10 //Decalre the value with
	fmt.Println(x)
	const saurabh = 10 // this is the constatnt value for the go

	var a, y float64 = 10, 20 //float assigned at the end not at the first
	// decalring the multiple values

	num := 30.0 // it will assume to be float64
	fmt.Print(a, y, num)
}

func add(a int64, b int64) {
	c := a + b
	return c
}
