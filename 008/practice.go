package main

import "fmt"

type Product struct {
	id string

	title string
	price float64
}

func main() {

	//1. array ]
	arr := [3]string{"saurabh", "cricker", "footbal"}
	fmt.Print(arr)
	fmt.Println(arr[1], arr[0])
	fmt.Println(arr[0:]) // slice

	arr2 := arr[0:2]
	fmt.Println(arr2)
	fmt.Println(cap(arr2))

	// dynamic string

	main := []string{"learn go", "learn basics "}

	fmt.Println(main)

	main = append(main, "saurabh bhai ")

	fmt.Println(main)

	//both are best way
	// prod := []Product{Product{"1", "Cream", 99.2}, Product{"2", "Soap", 43.9}}
	prod := []Product{{"1", "Cream", 99.2}, {"2", "Soap", 43.9}}

	fmt.Println(prod)

	prod = append(prod, Product{"3", "cream", 76.4})
	fmt.Println(prod)

	check := []float64{23, 45, 67, 12}

	fmt.Println(check)

	check2 := []float64{23.5, 54.2, 576.56}
	fmt.Println(check2)

	check = append(check, check2...)
	fmt.Println(check)

}
