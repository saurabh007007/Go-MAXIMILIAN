package main

import (
	"fmt"

	"example.com/pointers/sum"
)

func main() {
	a := 0
	b := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	val := sum.Sum(a, b)
	fmt.Println("Sum is", val)
	fmt.Println("Subtract is", sum.Subtract(a, b))
}
