package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, revErr := getUserInput("revenue")
	if revErr != nil {
		fmt.Println("ERROR")
		fmt.Println(revErr)
		return
	}
	expenses, expErr := getUserInput("expenses")
	if expErr != nil {
		fmt.Println("ERROR")
		fmt.Println(expErr)
		return
	}
	taxRate, taxErr := getUserInput("tax rate")
	if taxErr != nil {
		fmt.Println("ERROR")
		fmt.Println(taxErr)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateFiniantials(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", earningsBeforeTax)
	fmt.Printf("%.2f\n", earningsAfterTax)
	fmt.Printf("%.2f\n", ratio)

	txt := fmt.Sprintf("Earnings Before Tax: %.2f\nnEarnings After Tax:  %.2f\nRatio: %.2f", earningsBeforeTax, earningsAfterTax, ratio)
	writeToFile(txt)
}

func getUserInput(text string) (inp float64, err error) {
	fmt.Print("Enter your " + text + ": ")
	fmt.Scan(&inp)
	if inp <= 0 {
		return 0, errors.New("Invalid " + text + " provided")
	}
	return inp, nil
}

func calculateFiniantials(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat
	return
}

func writeToFile(text string) {
	os.WriteFile("profit_calculation.txt", []byte(text), 0644)
}
