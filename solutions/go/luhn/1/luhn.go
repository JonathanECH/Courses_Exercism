package luhn

import (
	"errors"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	numberInString := strings.ReplaceAll(id, " ", "")
	if len(numberInString) <= 1 {
		return false
	}
	var numbers []int
	for _, char := range numberInString {
		digit, err := strconv.Atoi(string(char))
		numbers = append(numbers, digit)
		if err != nil {
			errors.New("Numero no valido!")
			return false
		}
	}
	sum := 0
	for i, n := range numbers {
		offsetFromEnd := len(numberInString) - i
		if offsetFromEnd%2 == 1 {
			sum += n
		} else if n < 5 {
			sum += 2 * n
		} else {
			sum += 2*n - 9
		}
	}
	return sum%10 == 0
}

