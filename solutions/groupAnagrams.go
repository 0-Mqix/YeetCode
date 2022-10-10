package solutions

import "fmt"

func token(str string) string {
	list := make(map[rune]int)
	for _, c := range str {
		list[c]++
	}
	return fmt.Sprint(list)
}

func GroupAnagrams(words []string) [][]string {
	result := make([][]string, 0)
	list := make(map[string][]string)

	for _, v := range words {
		token := token(v)
		list[token] = append(list[token], v)
	}

	for _, v := range list {
		result = append(result, v)
	}

	return result
}
