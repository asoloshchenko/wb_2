package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bytes"
	"errors"
	"strconv"
)

// Unpack unpacks a string containing encoded characters and returns the decoded string.
// Empty string causes an error.
// It accepts an escape sequence as a special character.
// Example: a4bc2d5e => aaaabccddddde
//
//	abcd     => abcd
//	45       => "", error: invalid string
//	""       => "", error: empty string
//	qwe\4\5  => qwe45
//	qwe\45   => qwe44444
//	qwe\\5   => qwe\\\\
//
// It takes a string as input and returns a string and an error.
func Unpack(str string) (string, error) {
	if str == "" {
		return "", errors.New("empty string")
	}

	var buf bytes.Buffer
	runes := []rune(str)
	i := 0
	for i < len(runes) {
		switch {
		case runes[i] == '\\':
			if i+1 >= len(runes) {
				return "", errors.New("invalid string")
			}
			i++
			fallthrough

		case runes[i] < '0' || runes[i] > '9':
			c := runes[i]
			j := i + 1
			for j < len(runes) && runes[j] >= '0' && runes[j] <= '9' {
				j++
			}

			n, _ := strconv.Atoi(string(runes[i+1 : j]))
			for k := 0; k < n; k++ {
				buf.WriteRune(c)
			}
			if n == 0 {
				buf.WriteRune(c)
			}

			i = j

		default:
			return "", errors.New("invalid string")
		}
	}
	return buf.String(), nil
}
