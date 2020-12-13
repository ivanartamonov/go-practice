package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	if packedString == "" {
		return "", nil
	}

	var _, err = strconv.Atoi(packedString)
	if err == nil {
		return "", ErrInvalidString
	}

	chars := []rune(packedString)
	charsLength := len(chars)

	if unicode.IsDigit(chars[0]) {
		return "", ErrInvalidString
	}

	var sb = strings.Builder{}
	needEscape := false

	for index, charRune := range chars {
		nextIdx := index + 1
		prevIdx := index - 1
		nextCharExists := nextIdx < charsLength

		if charRune == '\\' && !needEscape {
			needEscape = true
			continue
		}

		if unicode.IsDigit(charRune) && !needEscape {
			isPrevWasEscaped := prevIdx-1 >= 0 && chars[prevIdx-1] == '\\'
			if unicode.IsDigit(chars[prevIdx]) && !isPrevWasEscaped {
				return "", ErrInvalidString
			}
			continue
		}

		if needEscape && unicode.IsLetter(charRune) {
			return "", ErrInvalidString
		}

		needEscape = false

		if nextCharExists {
			repeatCount, err := strconv.Atoi(string(chars[nextIdx]))

			if err == nil {
				appendStr := strings.Repeat(string(charRune), repeatCount)
				sb.WriteString(appendStr)
				continue
			}
		}

		sb.WriteRune(charRune)
	}

	return sb.String(), nil
}
