package main

import "fmt"

func main() {

	fmt.Println("Welcome to bank ::  ")
	fmt.Println("Please select an option:")
	fmt.Println("1.Chect the balance ")
	fmt.Println("2.Deposit money")
	fmt.Println("3.Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("Enter you choice :: ")
	var choice int
	var balance int = 1000 // Initial balance
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
			balance += amoutadd
			break
		}

	case 3:
		{
			var amoutWithdraw int
			fmt.Print("Enter the amount to withdraw :")
			fmt.Scan(&amoutWithdraw)
			fmt.Println("Balance withdrawn successfully")
			balance = balance - amoutWithdraw
			break
		}
	case 5:
		{
			fmt.Print("Exiting the bank system. Thank you for using our services!")
			return
		}

	default:
		fmt.Errorf("Invalid choice. Please select a valid option between 1 and 4.")
	}
}
