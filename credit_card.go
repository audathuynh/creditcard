package creditcard

import (
	"strconv"
	"strings"
)

// Check if a credit card number is valid by using Luhn algorithm.
// Spaces or tabs can be in the card number.
// Return true if the card number is valid, false otherwise.
func LuhnCheck(cardNumber string) bool {
	cardNumber = strings.Replace(cardNumber, "\t", "", -1)
	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	sum := 0
	flag := false // false: untouch a digit; true: double a digit.
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}
		if flag {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + 1
			}
		}
		sum += digit
		flag = !flag // toggle the flag to double or untouch a digit
	}
	return sum%10 == 0
}

const AMEX = "AMEX"
const VISA = "VISA"
const MASTERCARD = "MasterCard"
const DISCOVER = "Discover"
const UNKNOWN = "Unknown"

// Get the type of a credit card number.
// Spaces or tabs can be in the card number.
// +============+=============+===============+
// | Card Type  | Begins With | Number Length |
// +============+=============+===============+
// | AMEX       | 34 or 37    | 15            |
// +------------+-------------+---------------+
// | VISA       | 4           | 13 or 16      |
// +------------+-------------+---------------+
// | MasterCard | 51-55       | 16            |
// +------------+-------------+---------------+
// | Discover   | 6011        | 16            |
// +------------+-------------+---------------+
func GetCardType(cardNumber string) string {
	cardNumber = strings.Replace(cardNumber, "\t", "", -1)
	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	l := len(cardNumber)
	// cards with the prefix '3'
	if cardNumber[0] == '3' && l == 15 {
		c := cardNumber[1]
		if c == '4' || c == '7' {
			return AMEX
		}
		return UNKNOWN
	}
	// cards with the prefix '4'
	if cardNumber[0] == '4' {
		if l == 13 || l == 16 {
			return VISA
		}
		return UNKNOWN
	}
	// cards with the prefix '5'
	if cardNumber[0] == '5' && l == 16 {
		c := cardNumber[1]
		if '1' <= c && c <= '5' {
			return MASTERCARD
		}
		return UNKNOWN
	}
	// cards with the prefix '6'
	if l == 16 && cardNumber[0:4] == "6011" {
		return DISCOVER
	}
	return UNKNOWN
}
