package main

import (
    "fmt"
    "math"
)

func main() {
    const inflationRate = 2.5
    var investmentAmount, expectedReturnRate, years float64

    fmt.Print("Enter the investment amount: ")
    fmt.Scan(&investmentAmount)
    fmt.Print("Enter the epected return rate: ")
    fmt.Scan(&expectedReturnRate)
    fmt.Print("Enter the number of years: ")
    fmt.Scan(&years)

    futureValue := investmentAmount * math.Pow(1 + (expectedReturnRate / 100), years)
    futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
    
    formatedFv := fmt.Sprintf("Future Value: %.1f\n", futureValue)
    formatedRfv := fmt.Sprintf("Future Value: (Adjusted for inflation): %.1f\n", futureRealValue)
    
    fmt.Print(formatedFv, formatedRfv)
    // fmt.Println("Future Value: ", futureValue)
    // fmt.Printf("Future Value: %.1f\nFuture Value: (Adjusted for inflation): %.1f\n", futureValue, futureRealValue)
    // fmt.Println("Future Value: (Adjusted for inflation) ", futureRealValue)
}