package countformatter

import (
	"sort"
	"strconv"
)

type WordCount struct {
	Word  string
	Count int
}

// String implements standard Stringer
func (c WordCount) String() string {
	return c.Word + " : " + strconv.Itoa(c.Count) // concatenation is slightly faster than Sprintf
}

// GetWordCounts converts and sorts word counts
func GetWordCounts(wordCounts map[string]int) []WordCount {
	counts := make([]WordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		counts = append(counts, WordCount{
			Word:  word,
			Count: count,
		})
	}
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].Count == counts[j].Count {
			return counts[i].Word < counts[j].Word
		}
		return counts[i].Count > counts[j].Count
	})
	return counts
}
