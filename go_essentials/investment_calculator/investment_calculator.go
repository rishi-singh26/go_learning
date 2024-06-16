package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, expectedReturnRate, years float64

	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter the epected return rate: ")
	outputText("Enter the epected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter the number of years: ")
	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	// formatedFv := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formatedRfv := fmt.Sprintf("Future Value: (Adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Print(formatedFv, formatedRfv)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Printf(`
Future Value: %.1f
Future Value: (Adjusted for inflation): %.1f
`, futureValue, futureRealValue)
	// fmt.Println("Future Value: (Adjusted for inflation) ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
