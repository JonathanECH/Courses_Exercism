package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	code := strings.ReplaceAll(isbn, "-", "")
	code = strings.ToUpper(code)
	if len(code) <= 9 || len(code) > 10 {
		return false
	}
	sum := 0
	for i := 0 ; i < len(code); i++{
		digit := 0
		switch {
		case code[i] == 'X':
            if i != len(code)-1 {
                return false
            }
            digit = 10
		case code[i] >= '0' && code[i] <= '9':
			digit = int(code[i] - '0')
		default:
			return false
		}
		sum += digit * (10 - i)
	}
	return  sum%11 == 0
}