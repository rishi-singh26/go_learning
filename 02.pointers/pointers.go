package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	editAgeToAdultYears(&age)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
