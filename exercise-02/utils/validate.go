package utils

import (
	"fmt"
	"github.com/tranquoctuyen97/go-23/exercise-02/contants"
	"golang.org/x/exp/slices"
)

func ValidateSortType(sortType string) bool {
	var types = []string{contants.STRING, contants.NUMBER, contants.MIX}

	if slices.Contains(types, sortType) == false {
		println("Sort type: ", sortType, " is not supported")
		return false
	}

	return true
}

func ValidateListNumberFromSlice(slice []string) bool {
	for _, value := range slice {
		isFloat, _ := ParseFloat(value)
		if isFloat == false {
			return false
		}
	}

	return true
}

func ValidateListStringFromSlice(slice []string) bool {
	for _, value := range slice {
		if IsString(value) == false {
			fmt.Print(value, "value")
			return false
		}
	}

	return true
}
