package main

import (
	"fmt"
	"math"
)

// investement tracker for the personal

func main() {
	var investementAmount = 1000
	var expectedReturn = 5.5
	var years = 10
	var futureValue = float64(investementAmount) * math.Pow(1+expectedReturn/100, float64(years))

	fmt.Printf("The future value of the investement is: %.2f\n", futureValue)

}

// go run .  // it runs the current directory
// go run main.go  // it runs the main file
