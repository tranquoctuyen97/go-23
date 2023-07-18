package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Print("The data you entered is not enough")
		return
	}

	countryCode := args[len(args)-1]
	var isValid = ValidateCountryCode(countryCode)
	if isValid == false {
		return
	}

	firstName := args[0]
	lastName := args[1]
	midName := strings.Join(args[2:len(args)-1], " ")

	orderedNames := []string{firstName, lastName, midName}
	sort.Strings(orderedNames)

	fmt.Println(strings.Join(orderedNames, " "))
}

func ValidateCountryCode(countryCode string) bool {
	var isValid bool
	switch countryCode {
	case "VN", "US":
		isValid = true
	default:
		fmt.Println("The country code is not supported")
		isValid = false
	}

	return isValid
}
