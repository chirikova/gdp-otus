package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}

	splitPattern := regexp.MustCompile(`[\p{P}\p{S}]*\s+`)

	type wordOccurrence struct {
		word  string
		count int
	}

	// считаем количество вхождений каждого слова
	wordsCount := make(map[string]int, 0)

	for _, val := range splitPattern.Split(str, -1) {
		val = strings.ToLower(val)

		if len(val) == 0 {
			continue
		}

		wordsCount[val]++
	}

	// формируем срез из слов с количеством вхождений
	words := make([]wordOccurrence, 0)

	for word, count := range wordsCount {
		words = append(words, wordOccurrence{word, count})
	}

	// сортируем по убыванию и лексикографически
	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return strings.Compare(words[i].word, words[j].word) <= 0
		}

		return words[i].count > words[j].count
	})

	// формируем первые топ 10 или меньше
	topCount := 10
	if len(words) < 10 {
		topCount = len(words)
	}

	result := make([]string, 0)

	for _, item := range words[0:topCount] {
		result = append(result, item.word)
	}

	return result
}
