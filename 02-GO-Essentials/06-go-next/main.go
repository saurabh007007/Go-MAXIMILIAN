package main

import "fmt"

func main() {
	a := 100

	fmt.Print("hi there ")
	fmt.Printf("This is a simple Go backtich examples %v", a)
	fmt.Printf(" this ")
	formatedv := fmt.Sprintf("futuer value is %f", 100.0)
	fmt.Println(formatedv)
}

// %f     default width, default precision
// %9f    width 9, default precision
// %.2f   default width, precision 2
// %9.2f  width 9, precision 2
// %9.f   width 9, precision 0

// %v	the value in a default format
// 	when printing structs, the plus flag (%+v) adds field names
// %#v	a Go-syntax representation of the value
// 	(floating-point infinities and NaNs print as ±Inf and NaN)
// %T	a Go-syntax representation of the type of the value
// %%	a literal percent sign; consumes no value
