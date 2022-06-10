package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	ErrFirstCharIsDigit = errors.New("first char is digit")
	ErrLastCharIsDigit  = errors.New("last char is digit")
	ErrTwoDigits        = errors.New("two digits in a row")
)

func Unpack(str string) (string, error) {
	// Работаем с крайними ситуациями
	if len(str) == 0 {
		return "", nil
	}

	runeStr := []rune(str)

	// Первый символ
	if unicode.IsDigit(runeStr[0]) {
		return "", ErrFirstCharIsDigit
	}

	// Последний символ
	if unicode.IsDigit(runeStr[utf8.RuneCount([]byte(str))-1]) {
		return "", ErrLastCharIsDigit
	}

	// Основной код
	var result string
	var digitCounter int

	for i, char := range runeStr {
		if unicode.IsDigit(char) {
			digitCounter++

			if digitCounter > 1 {
				return "", ErrTwoDigits
			}

			continue
		}

		if i == len(runeStr)-1 {
			result += string(char)
			continue
		}

		nextChar := runeStr[i+1]

		if digit, err := strconv.Atoi(string(nextChar)); err == nil {
			digitCounter = 0

			if digit <= 0 {
				continue
			} else {
				result += strings.Repeat(string(char), digit)
			}
		} else {
			result += string(char)
		}
	}

	return result, nil
}
