package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go bank")
	fmt.Println("What do you want to do?: ")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1
	if choice == 1 {
		fmt.Printf("Your balance is %0.2f\n", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += +depositAmount
		fmt.Printf("Balance updated! New amount: %0.1f\n", accountBalance)
	}

	fmt.Printf("Your choice: %v\n", choice)
}
