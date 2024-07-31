package main

import "fmt"

func getUserSelection() (selection int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check banlance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Scan(&selection)
	return
}
