package main

import (
	"fmt"
	"github.com/tranquoctuyen97/go-23/exercise-02/contants"
	"github.com/tranquoctuyen97/go-23/exercise-02/utils"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Print("The data you entered is not enough")
		return
	}

	sortRawDirection := strings.Split(args[0], "-")
	var sortOrder = contants.ASC
	var sortType string
	if len(sortRawDirection) == 2 {
		sortType = sortRawDirection[1]
		sortOrder = utils.DetectSortOrder("-")
	} else {
		sortType = sortRawDirection[0]
	}

	isValidSortType := utils.ValidateSortType(sortType)
	if !isValidSortType {
		return
	}

	numberAndStringMixedArr := args[1:]

	if sortType == contants.STRING {
		isValid := utils.ValidateListStringFromSlice(numberAndStringMixedArr)

		if !isValid {
			fmt.Println("Input strings is not valid")
			return
		}
	} else if sortType == contants.NUMBER {
		isValid := utils.ValidateListNumberFromSlice(numberAndStringMixedArr)

		if !isValid {
			fmt.Println("Input numbers is not valid")
			return
		}
	}

	needSortNumbers := utils.GetListNumberFromSlice(numberAndStringMixedArr)
	needSortStrings := utils.GetListStringFromSlice(numberAndStringMixedArr)

	utils.SortListStringBySortOrder(needSortStrings, sortOrder)
	utils.SortListNumberBySortOrder(needSortNumbers, sortOrder)

	convertedNumbers := utils.ConvertListFloatToString(needSortNumbers)

	fmt.Println(strings.Join(convertedNumbers, " "), strings.Join(needSortStrings, " "))
}
