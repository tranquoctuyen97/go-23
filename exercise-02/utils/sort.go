package utils

import (
	"github.com/tranquoctuyen97/go-23/exercise-02/contants"
	"sort"
	"strconv"
)

func DetectSortOrder(sortRawOrder string) string {
	var sortOrder string
	if sortRawOrder == "-" {
		sortOrder = contants.DESC
	} else {
		sortOrder = contants.ASC
	}

	return sortOrder
}

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

func SortListStringBySortOrder(chars []string, sortOrder string) {
	if sortOrder == contants.DESC {
		sort.Sort(sort.Reverse(sort.StringSlice(chars)))
	} else if sortOrder == contants.ASC {
		sort.Strings(chars)
	}
}

func SortListNumberBySortOrder(numbers []float64, sortOrder string) {
	if sortOrder == contants.DESC {
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[j] < numbers[i]
		})
	} else {
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
	}
}

func ConvertListFloatToString(numbers []float64) []string {
	var rs []string

	for _, value := range numbers {
		rs = append(rs, strconv.FormatFloat(value, 'f', -1, 64))
	}

	return rs
}
