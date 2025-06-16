package main

// int  a no without the decimal places
// float64 a no with the decimal places
// string a text value
// bool a true of false value
import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturn float64 = 5.5
	var years float64 = 10
	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturn/100, years)

	fmt.Println("The future value of the investment is: ", futureValue)

}
