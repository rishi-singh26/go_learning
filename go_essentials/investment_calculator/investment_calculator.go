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

    // fmt.Println("Future Value: ", futureValue)
    fmt.Printf("Future Value: %v\nFuture Value: (Adjusted for inflation): %v", futureValue, futureRealValue)
    // fmt.Println("Future Value: (Adjusted for inflation) ", futureRealValue)
}