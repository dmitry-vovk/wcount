package countformatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	input := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"foo":   15,
		"bar":   15,
	}
	expected := []WordCount{
		{
			Word:  "bar",
			Count: 15,
		},
		{
			Word:  "foo",
			Count: 15,
		},
		{
			Word:  "three",
			Count: 3,
		},
		{
			Word:  "two",
			Count: 2,
		},
		{
			Word:  "one",
			Count: 1,
		},
	}
	assert.Equal(t, expected, GetWordCounts(input))
}

func TestStringer(t *testing.T) {
	testCases := []struct {
		input  WordCount
		expect string
	}{
		{
			input: WordCount{
				Word:  "hello",
				Count: 15,
			},
			expect: "hello : 15",
		},
		{
			input: WordCount{
				Word:  "foo bar",
				Count: 1,
			},
			expect: "foo bar : 1",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.expect, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.input.String())
		})
	}
}
