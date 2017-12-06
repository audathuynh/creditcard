package creditcard_test

import (
	"fmt"
	"github.com/audathuynh/creditcard"
	"testing"
)

func TestLuhnCheck(t *testing.T) {
	fmt.Println("Testing Luhn algorithm")
	var cardNumber string
	var result, expected bool

	// should sort the card numbers by the first digit.
	cardNumber = "378282246310005"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "4408 0412 3456 7893"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "4111111111111111"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}
	cardNumber = "4111111111111"
	result = creditcard.LuhnCheck(cardNumber)
	expected = false
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}
	cardNumber = "4012888888881881"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "5105 1051 0510 5100"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "5105 1051 0510 5106"
	result = creditcard.LuhnCheck(cardNumber)
	expected = false
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "6011111111111117"
	result = creditcard.LuhnCheck(cardNumber)
	expected = true
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

	cardNumber = "9111111111111111"
	result = creditcard.LuhnCheck(cardNumber)
	expected = false
	if result != expected {
		t.Errorf("Error when testing Luhn algorithm with the card number", cardNumber)
	}

}

func TestCardType(t *testing.T) {
	fmt.Println("Testing card types")
	var cardNumber string
	var result, expected string

	// should sort the card numbers by the first digit.
	cardNumber = "378282246310005"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.AMEX
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}
	cardNumber = "4408 0412 3456 7893"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.VISA
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "4111111111111111"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.VISA
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "4111111111111"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.VISA
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "4012888888881881"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.VISA
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "5105 1051 0510 5100"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.MASTERCARD
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "5105 1051 0510 5106"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.MASTERCARD
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "6011111111111117"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.DISCOVER
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

	cardNumber = "9111111111111111"
	result = creditcard.GetCardType(cardNumber)
	expected = creditcard.UNKNOWN
	if result != expected {
		t.Errorf("Error when testing the card type of the card number", cardNumber)
	}

}

func TestProgram(t *testing.T) {
	fmt.Println("Run and show results")
	cardNumbers := []string{
		"4408 0412 3456 7893",
		"4111111111111111",
		"4111111111111",
		"4012888888881881",
		"378282246310005",
		"6011111111111117",
		"5105 1051 0510 5100",
		"5105 1051 0510 5106",
		"9111111111111111"}
	for _, cardNumber := range cardNumbers {
		cardType := creditcard.GetCardType(cardNumber)
		valid := creditcard.LuhnCheck(cardNumber)
		result := cardType + ": " + cardNumber + " "
		if valid {
			result += "(valid)"
		} else {
			result += "(invalid)"
		}
		fmt.Println(result)
	}
}
