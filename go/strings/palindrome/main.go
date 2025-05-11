package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// Приведение строки к нижнему регистру
	s = strings.ToLower(s)
	// Удаление всех символов, не являющихся буквами или цифрами
	r := []rune(s)
	var cleanedRunes []rune
	for _, char := range r {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleanedRunes = append(cleanedRunes, char)
		}
	}

	// Проверка на палиндром
	for i, j := 0, len(cleanedRunes)-1; i < j; i, j = i+1, j-1 {
		if cleanedRunes[i] != cleanedRunes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("Лёша на полке клопа нашёл")) // true
	fmt.Println(isPalindrome("Never odd or even"))         // true
}
