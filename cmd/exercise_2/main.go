package main

import (
	"fmt"

	"github.com/dmitry-vovk/wcount/countformatter"
	"github.com/dmitry-vovk/wcount/parallelcounter"
	"github.com/dmitry-vovk/wcount/stream"
	"github.com/dmitry-vovk/wcount/wordcounter"
)

func main() {
	const n = 10 // How many streams to run
	// Prepare the list of counters
	var counters []*wordcounter.Counter
	for i := 0; i < n; i++ {
		counters = append(counters,
			wordcounter.New(stream.MakeSlowReader(stream.ExampleText)),
		)
	}
	// Create an instance of parallel counter
	pc := parallelcounter.New(counters...)
	// Run it
	countsC := pc.Run()
	// Consume and display intermediate results
	go func() {
		for c := range countsC {
			fmt.Println("Fresh counts:")
			printCounts(c)
		}
	}()
	// Wait until counting streams finish
	pc.Wait()
	// Display the most recent results
	fmt.Println("Final counts:")
	printCounts(pc.CurrentCounts())
}

// printCounts helps to print a list of word counts
func printCounts(counts []countformatter.WordCount) {
	for _, w := range counts {
		fmt.Printf("%s\n", w)
	}
}
