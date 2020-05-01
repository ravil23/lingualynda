package wc

import (
	"regexp"
	"sort"
	"strings"
)

type WordCountPair struct {
	Word  string
	Count uint
}

func WordCounts(text string) map[string]uint {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}
	counts := make(map[string]uint)
	words := strings.Split(text, " ")
	for _, word := range words {
		filteredWord := reg.ReplaceAllString(word, "")
		if len(filteredWord) > 0 {
			counts[strings.ToLower(filteredWord)]++
		}
	}
	return counts
}

func SortByValuesAndKeys(collection map[string]uint, reverse bool) []WordCountPair {
	items := make([]WordCountPair, 0, len(collection))
	for word, count := range collection {
		items = append(items, WordCountPair{Word: word, Count: count})
	}
	if reverse {
		sort.Slice(items, func(i, j int) bool {
			return items[i].Count > items[j].Count || (items[i].Count == items[j].Count && items[i].Word > items[j].Word)
		})
	} else {
		sort.Slice(items, func(i, j int) bool {
			return items[i].Count < items[j].Count || (items[i].Count == items[j].Count && items[i].Word < items[j].Word)
		})
	}
	return items
}
