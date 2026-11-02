package luhn

import "strings"

func Valid(id string) bool {
	digits := strings.ReplaceAll(id, " ", "")
	if len(digits) <= 1 {
		return false
	}
	sum := 0
	double := false
	for i := len(digits) - 1; i >= 0; i-- {
		char := digits[i]
		if char < '0' || char > '9' {
			return false
		}
		digit := int(char - '0')
		if double{
			digit *= 2
			if digit > 9{
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}