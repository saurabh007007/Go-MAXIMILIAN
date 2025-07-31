package main

import "fmt"

func main() {

	//1. array ]
	arr := [3]string{"saurabh", "cricker", "footbal"}
	fmt.Print(arr)
	fmt.Println(arr[1], arr[0])
	fmt.Println(arr[0:]) // slice

	arr2 := arr[0:2]
	fmt.Println(arr2)
	fmt.Println(cap(arr2))

}
