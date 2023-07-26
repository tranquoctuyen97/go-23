package utils

import (
	"sort"
	"strconv"
)

func GetListNumberFromSlice(slice []string) []float64 {
	var result []float64
	for _, value := range slice {
		isFloat, floatValue := ParseFloat(value)
		if isFloat {
			result = append(result, floatValue)
		}

	}

	return result
}

func GetListStringFromSlice(slice []string) []string {
	var result []string
	for _, value := range slice {
		if IsString(value) == true {
			result = append(result, value)
		}
	}

	return result
}

func SortListStringBySortOrder(chars []string) {
	sort.Strings(chars)
}

func SortListNumberBySortOrder(numbers []float64) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
}

func ConvertListFloatToString(numbers []float64) []string {
	var rs []string

	for _, value := range numbers {
		rs = append(rs, strconv.FormatFloat(value, 'f', -1, 64))
	}

	return rs
}
