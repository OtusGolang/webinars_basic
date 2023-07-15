package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	// Создаем мапу для хранения упоминаний слов
	wordCount := make(map[string]int)

	// Разделяем текст на отдельные слова
	words := strings.Fields(text)

	// Обходим каждое слово
	for _, word := range words {
		// Удаляем пунктуацию и приводим слово к нижнему регистру
		cleanWord := strings.Trim(word, ",.!?")
		cleanWord = strings.ToLower(cleanWord)

		// Увеличиваем счетчик упоминаний слова в мапе
		wordCount[cleanWord]++
	}

	return wordCount
}

func main() {
	text := "Hello, how are you? Are you ready to learn Go? Go is a great language."
	wordCount := countWords(text)

	fmt.Println(wordCount)
}
