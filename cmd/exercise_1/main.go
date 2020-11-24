package main

import (
	"fmt"
	"log"

	"github.com/dmitry-vovk/wcount/countformatter"
	"github.com/dmitry-vovk/wcount/stream"
	"github.com/dmitry-vovk/wcount/wordcounter"
)

func main() {
	text := `The cat sat on the mat.`
	// Uncomment the line below to run on the example text
	// text = stream.ExampleText
	// Create an instance of word counter with fast stream reader
	r := wordcounter.New(stream.MakeFastReader(text))
	// Run the counter
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
	// Get the counts
	wordCounts := r.GetWordCounts()
	// Display formatted word counts
	for _, w := range countformatter.GetWordCounts(wordCounts) {
		fmt.Printf("%s\n", w)
	}
}
