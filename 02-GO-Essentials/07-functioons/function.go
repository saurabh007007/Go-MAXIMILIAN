package main

import "fmt"

func main() {
	Output("Saurabh Yadav")

}

func Output(text string) {
	fmt.Print("Welcome ", text)
}

func ReturnValues(text string, x int) (string, int) {
	return text, x
}

// alternative return values
func ReturnValues2(text string, x int) (fv string, gx int) {
	fv = text
	gx = x

}
