package main

import "fmt"

func countWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	ages := map[string]int{
		"Алиса": 25,
		"Боб":   30,
	}
	fmt.Println("Возраст людей:", ages)

	words := []string{"Кот", "Собака", "Кот", "Собака", "Собака", "Кот"}
	wordCount := countWords(words)
	fmt.Println("Количество каждого слова:", wordCount)

	delete(ages, "Боб")
	fmt.Println("После удаления Боба:", ages)
	if age, ok := ages["Боб"]; ok {
		fmt.Println("Боб найден, возраст:", age)
	} else {
		fmt.Println("Боб отсутствует в карте")
	}
}
