package parallelcounter

import (
	"log"
	"sync"
	"time"

	"github.com/dmitry-vovk/wcount/countformatter"
	"github.com/dmitry-vovk/wcount/wordcounter"
)

type Counter struct {
	counters     []*wordcounter.Counter
	wordCountsC  chan []countformatter.WordCount
	wg           *sync.WaitGroup
	ticker       *time.Ticker
	pollInterval time.Duration
}

const defaultPollInterval = time.Second * 10

// New returns an instance of aggregate word counter
func New(counters ...*wordcounter.Counter) *Counter {
	if len(counters) == 0 {
		// No counters -- probably a bug
		panic("expected at least one counter")
	}
	r := Counter{
		counters:     counters,
		wordCountsC:  make(chan []countformatter.WordCount),
		wg:           &sync.WaitGroup{},
		pollInterval: defaultPollInterval,
	}
	return &r
}

func (c *Counter) WithPollInterval(interval time.Duration) *Counter {
	if interval <= 0 {
		panic("invalid poll interval")
	}
	c.pollInterval = interval
	return c
}

// Run initiates counting and returns a chan for intermediate results
func (c *Counter) Run() chan []countformatter.WordCount {
	c.ticker = time.NewTicker(defaultPollInterval)
	c.runner()
	go c.poller()
	return c.wordCountsC
}

// Wait blocks until all the counters are done
func (c *Counter) Wait() {
	c.wg.Wait()
	c.ticker.Stop()
}

// CurrentCounts returns aggregated counts for all counters
func (c *Counter) CurrentCounts() []countformatter.WordCount {
	return countformatter.GetWordCounts(c.collectCounts())
}

// runner starts the counters
func (c *Counter) runner() {
	for i := range c.counters {
		c.wg.Add(1)
		go func(counter *wordcounter.Counter) {
			if err := counter.Run(); err != nil {
				log.Fatal(err) // Normally we do not expect any error, therefore it is fatal
			}
			c.wg.Done()
		}(c.counters[i])
	}
}

// poller periodically returns aggregated counts
func (c *Counter) poller() {
	for range c.ticker.C {
		c.wordCountsC <- countformatter.GetWordCounts(c.collectCounts())
	}
	close(c.wordCountsC)
}

// collectCounts collects and aggregates word counts from all counters
func (c *Counter) collectCounts() map[string]int {
	allWordsCount := make(map[string]int)
	for _, counter := range c.counters {
		for word, count := range counter.GetWordCounts() {
			if _, ok := allWordsCount[word]; ok {
				allWordsCount[word] += count
			} else {
				allWordsCount[word] = count
			}
		}
	}
	return allWordsCount
}
