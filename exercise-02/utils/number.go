package utils

import (
	"strconv"
)

func ParseFloat(val string) (bool, float64) {
	value, err := strconv.ParseFloat(val, 10)

	if err != nil {
		return false, 0
	}

	return true, value
}

func IsString(strVar string) bool {
	isFloat, _ := ParseFloat(strVar)
	if isFloat {
		return false
	}

	return true
}
