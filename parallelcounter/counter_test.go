package parallelcounter

import (
	"testing"
	"time"

	"github.com/dmitry-vovk/wcount/countformatter"
	"github.com/dmitry-vovk/wcount/stream"
	"github.com/dmitry-vovk/wcount/wordcounter"
	"github.com/stretchr/testify/assert"
)

func TestParallelCounter(t *testing.T) {
	assert.Panics(t, func() {
		New()
	})
	assert.Panics(t, func() {
		New().WithPollInterval(-time.Second)
	})
	assert.Panics(t, func() {
		New().WithPollInterval(0)
	})
	c1 := wordcounter.New(stream.MakeFastReader(stream.ExampleText))
	c2 := wordcounter.New(stream.MakeFastReader(stream.ExampleText))
	c := New(c1, c2)
	counts := c.Run()
	go func() {
		for range counts {
			// Drain the channel
		}
	}()
	c.Wait()
	if wordCounts := c.CurrentCounts(); assert.Equal(t, 155, len(wordCounts)) {
		assert.Equal(t, countformatter.WordCount{Word: "the", Count: 30}, wordCounts[0])
	}
}
