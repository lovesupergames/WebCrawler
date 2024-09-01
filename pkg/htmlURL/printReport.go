package htmlURL

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n  REPORT for %s\n=============================\n", baseURL)

	pairs := SortMapByValueAndThenByKey(pages)

	for _, v := range pairs {
		fmt.Printf("Found %d internal links to %s\n", v.Value, v.Key)
	}
}

func SortMapByValueAndThenByKey(m map[string]int) []Pair {
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		// Сначала сортируем по значению (убывание)
		if pairs[i].Value != pairs[j].Value {
			return pairs[i].Value > pairs[j].Value
		}

		// Если значения равны, сортируем по ключу (возрастание)
		return pairs[i].Key < pairs[j].Key
	})

	return pairs
}
