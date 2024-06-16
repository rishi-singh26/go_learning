package main

import "fmt"

func main() {
	revenue := getUserInput("Enter you revenue: ")
	expenses := getUserInput("Enter your expenses: ")
	taxRate := getUserInput("Enter you tax rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculateFiniantials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)
}

func getUserInput(text string) (inp float64) {
	fmt.Print(text)
	fmt.Scan(&inp)
	return
}

func calculateFiniantials(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat
	return
}
