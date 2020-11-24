package wordcounter

import (
	"io"
	"sync"
	"unicode"

	"github.com/dmitry-vovk/wcount/stream"
)

type Counter struct {
	stream      stream.InputStream // Source of runes
	charC       chan rune
	currentWord string
	words       map[string]int
	wordsM      sync.RWMutex
}

// New returns new word counter for given InputStream
func New(stream stream.InputStream) *Counter {
	r := Counter{
		stream: stream,
		charC:  make(chan rune),
		words:  make(map[string]int),
	}
	go r.charProcessor()
	return &r
}

// Run reads incoming runes and handles potential errors
func (wc *Counter) Run() error {
	defer func() {
		close(wc.charC)
		wc.stream.Dispose()
	}()
	for {
		r, err := wc.stream.TakeChar()
		if err != nil {
			// Getting EOF is ok, any other error gets reported to the caller
			if err != io.EOF {
				return err
			}
			break
		}
		wc.charC <- normalize(r)
	}
	return nil
}

// GetWordCounts returns list of words and number of occurrences
func (wc *Counter) GetWordCounts() map[string]int {
	wordCounts := make(map[string]int)
	for word, count := range wc.words {
		wordCounts[word] = count
	}
	return wordCounts
}

// charProcessor handles stream of runes
func (wc *Counter) charProcessor() {
	for r := range wc.charC {
		if isWordCharacter(r) { // within a word
			wc.currentWord += string(r)
		} else { // word boundary
			if wc.currentWord != "" {
				wc.wordsM.Lock()
				if _, ok := wc.words[wc.currentWord]; ok {
					wc.words[wc.currentWord]++
				} else {
					wc.words[wc.currentWord] = 1
				}
				wc.wordsM.Unlock()
				wc.currentWord = ""
			}
		}
	}
}

// normalize processes a rune returning normal form (here, lowercase)
func normalize(r rune) rune {
	return unicode.ToLower(r)
}

// isWordCharacter tells if the rune is a word character
func isWordCharacter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
