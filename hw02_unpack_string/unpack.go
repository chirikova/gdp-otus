package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isRuneEscapeChar(r rune) bool {
	return string(r) == "\\"
}

func isWritable(r rune, isEscaped bool) bool {
	return (!isEscaped && unicode.IsLetter(r)) || (isEscaped && (unicode.IsDigit(r) || isRuneEscapeChar(r)))
}

func Unpack(packedStr string) (string, error) {
	if len(packedStr) == 0 {
		return "", nil
	}

	var resString strings.Builder
	var prevChar rune
	var charEscaped bool

	for i, char := range packedStr {
		symbolRepeat := 1

		// строка не может начинаться с цифры
		if i == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		// в строке допускаются цифры, а не числа
		if unicode.IsDigit(prevChar) && !charEscaped && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		// экранированных букв не допускается
		if isRuneEscapeChar(prevChar) && !charEscaped && unicode.IsLetter(char) {
			return "", ErrInvalidString
		}

		// если предыдущий рун бэкслеш, отмечаем текущий как экранированный
		if isRuneEscapeChar(prevChar) && !charEscaped && (unicode.IsDigit(char) || isRuneEscapeChar(char)) {
			charEscaped = true

			prevChar = char

			continue
		}

		// задаем количество повторений
		if unicode.IsDigit(char) && isWritable(prevChar, charEscaped) {
			atoi, _ := strconv.Atoi(string(char))

			symbolRepeat = atoi
		}

		// пишем в строку
		if prevChar != 0 && isWritable(prevChar, charEscaped) {
			resString.WriteString(strings.Repeat(string(prevChar), symbolRepeat))

			charEscaped = false
		}

		prevChar = char
	}

	// последний символ также проверяем
	if prevChar != 0 && isWritable(prevChar, charEscaped) {
		resString.WriteString(string(prevChar))
	}

	return resString.String(), nil
}
