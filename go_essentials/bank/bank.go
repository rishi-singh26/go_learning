package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go bank")
	var bankBalance, readErr = readFromFile()
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
			writeToFile(bankBalance)
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
			writeToFile(bankBalance)
			fmt.Printf("Your bank balance: %v\n\n", bankBalance)
		case 4:
			fmt.Print("Exiting...\n\n")
			break appLoop
		default:
			fmt.Println("Enter a valid input!")
		}
	}
}

func writeToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}

func readFromFile() (float64, error) {
	data, readErr := os.ReadFile(accountBalanceFile)
	if readErr != nil {
		return 0, errors.New("failed to find balance file")
	}
	balanceTxt := string(data)
	balance, parseErr := strconv.ParseFloat(balanceTxt, 64)
	if parseErr != nil {
		return 0, errors.New("failed to parse file content")
	}
	return balance, nil
}

func getUserSelection() (selection int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check banlance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Scan(&selection)
	return
}
