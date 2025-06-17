package main

import "fmt"

func main() {
	fmt.Println("This is the profit calculator program / CLI")
	var amount float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter the amount of revenu : ")
	fmt.Scan(&amount)

	fmt.Print("Enter the expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate ")
	fmt.Scan(&taxrate)

	ebt := amount - expenses
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit

	fmt.Println("Earning before the taxes :", ebt)
	fmt.Println("Here is the profite ", profit)
	fmt.Println("Here is the ratio ", ratio)

}
