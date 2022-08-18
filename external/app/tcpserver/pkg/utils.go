package pkg

import (
	"strconv"
	"strings"
)

func GetNumbersFromString(data string) ([]int, error) {
	var numbers []int
	var matchNumbers []string

	for _, s := range strings.Split(data, "\r\n") {
		matchNumbers = append(matchNumbers, strings.Split(s, ",")...)
	}

	for _, element := range matchNumbers {
		if element != "" {
			i, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
