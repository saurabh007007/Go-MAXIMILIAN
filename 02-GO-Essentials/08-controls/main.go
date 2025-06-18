package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func WriteBalanceTofile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644) // 0644 read and write by owner otherr can read only
}

func getBalanceFromFile() float64 {
	var balanceText string
	var balance float64
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText = string(data)
	balance, _ = strconv.ParseFloat(balanceText, 64)
	return balance

}

func main() {
	var balance = getBalanceFromFile()
	for i := 0; i < 200; i++ {
		fmt.Println("Welcome to bank ::  ")
		fmt.Println("Please select an option:")
		fmt.Println("1.Check the balance ")
		fmt.Println("2.Deposit money")
		fmt.Println("3.Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println("Enter you choice :: ")
		var choice int
		// Initial balance
		fmt.Scan(&choice)
		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please select a valid option between 1 and 4.")
			return
		}
		switch choice {
		case 1:
			{
				fmt.Println("your cuurrent balance is ", balance)
				break
			}
		case 2:
			{
				var amoutadd int
				fmt.Println("Enter the amount to add : ")
				fmt.Scan(&amoutadd)
				fmt.Println("Balance added successfully")
				balance += float64(amoutadd)
				WriteBalanceTofile(float64(balance))
				break
			}

		case 3:
			{
				var amoutWithdraw int
				fmt.Print("Enter the amount to withdraw :")
				fmt.Scan(&amoutWithdraw)
				fmt.Println("Balance withdrawn successfully")
				balance = balance - float64(amoutWithdraw)1

				WriteBalanceTofile(float64(balance))
				break
			}
		case 4:
			{
				fmt.Print("Exiting the bank system. Thank you for using our services!")
				return
			}

		default:
			fmt.Errorf("Invalid choice. Please select a valid option between 1 and 4.")
		}

	}

}

// panic() // it is used to stop the execution of the program and print the error message
