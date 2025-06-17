// input from the user

package main

import "fmt"

func main() {
	var investment float64 // they are declartion
	var rate float64
	var time float64
	fmt.Print("Enter the investement amount ")
	fmt.Scan(&investment) //it expectes the pointer to populate the vaule in the invertement
	fmt.Print("Enter the rate of the inrests :")
	fmt.Scan(&rate)
	fmt.Print("enter the time ")
	fmt.Scan(&time)
	var futureValue float64 = investment * rate * time / 100

	fmt.Print("here the future valsaue of the rerturn SI ", futureValue)
}
