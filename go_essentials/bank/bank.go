package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go bank")
	var bankBalance, readErr = fileops.GetFloatFromFile(accountBalanceFile)
	if readErr != nil {
		fmt.Println("ERROR")
		fmt.Println(readErr)
		fmt.Println("----------------")
	}
appLoop:
	for {
		userSelection := getUserSelection()

		switch userSelection {
		case 1:
			fmt.Printf("Your bank balance: %v\n\n", bankBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64 = 0
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Print("Enter a valid deposint amount! \n\n")
				continue
			}
			bankBalance += depositAmount
			fileops.WriteFloatToFile(bankBalance, accountBalanceFile)
			fmt.Printf("Your bank balance: %v\n\n", bankBalance)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64 = 0
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > bankBalance {
				fmt.Printf("You can withdraw more then you have!\nYour bank baance: %v\n\n", bankBalance)
				continue
			}
			bankBalance -= withdrawAmount
			fileops.WriteFloatToFile(bankBalance, accountBalanceFile)
			fmt.Printf("Your bank balance: %v\n\n", bankBalance)
		case 4:
			fmt.Print("Exiting...\n\n")
			break appLoop
		default:
			fmt.Println("Enter a valid input!")
		}
	}
}
