package main

import (
	"errors"
	"fmt"
)

func main() {
	title, err := getUserInput("Note title:: ")
	if err != nil {
		fmt.Print(err)
		return

	}
	content, err := getUserInput("Note data :: ")
	if err != nil {
		fmt.Print(err)
		return
	}

}
func getUserInput(prompt string) (string, error) {
	var value string
	fmt.Print(prompt)
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("Invalid option")
	}
	return value, nil
}
