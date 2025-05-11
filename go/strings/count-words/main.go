package main

import (
	"strings"
)

func CountWords(text string) map[string]int {
	result := make(map[string]int)

	// Разбиение текста на слова с учётом множественных пробелов
	words := strings.Fields(text)

	for _, word := range words {
		// Преобразование слова к нижнему регистру и удаление лишних символов
		word = strings.ToLower(strings.Trim(word, ".,!?"))

		// Увеличение счётчика слова в результате
		result[word]++
	}

	return result
}

func main() {
	// Пример использования
	text := "Hello, world! This is a test."
	wordCounts := CountWords(text)
	for word, count := range wordCounts {
		println(word, ":", count)
	}
}
