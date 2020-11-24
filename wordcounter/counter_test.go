package wordcounter

import (
	"testing"

	"github.com/dmitry-vovk/wcount/stream"
	"github.com/stretchr/testify/assert"
)

func TestWordCounter(t *testing.T) {
	const expectedNumberOfWords = 155
	r := New(stream.MakeFastReader(stream.ExampleText))
	if assert.NoError(t, r.Run()) {
		wordCounts := r.GetWordCounts()
		assert.Equal(t, expectedNumberOfWords, len(wordCounts))
	}
}

func TestWordCharacters(t *testing.T) {
	var wordChars = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	for _, c := range wordChars {
		assert.True(t, isWordCharacter(c))
	}
	var nonWordChars = " \t\n01234567890.,“”?!-()_"
	for _, c := range nonWordChars {
		assert.False(t, isWordCharacter(c))
	}
}
