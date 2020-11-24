# Word Count project

The project consists of several packages:

### `wordcounter`

Given a stream (implementation of `stream.InputStream` interface), processes it returning counts for each word encountered.

It handles standard `io.EOF` error internally as it is a common way of letting know the stream of data ended,
returning any other kind of error to be handled by the caller.

There are two helper functions that may be adjusted according to requirements:
 * `normalize(rune) rune` -- accepts a rune and returns its lowercase variant, assuming this is requirement from the example output.
 * `isWordCharacter(rune) bool` -- given a rune, tells if it is a word character. 
 In the current implementation accepts English 26 letters, either lowercase and uppercase.
  There may be a seeming contradiction of requirements with normalize() function, 
  but I have decided to have it this way, so each function handles its own scope of responsibilities,
  and if the requirements for normalization or criteria what is a word character changes, 
  there will be just one change in one place.  

### `parallelcounter`

This is a higher level package, accepting arbitrary non-zero number of `wordcounter.Counter` instances.

The package interface is similar, but with some significant distinctions.
Since the requirement is to periodically display intermediary values, `Run()` method is non-blocking,
but instead it returns a channel into which it sends the values, 
also it has `Wait()` method that blocks until all streams finish.

Also, there is a way to set an optional interval for getting intermediate results,
 provided by method `WithPollInterval(time.Duration)`. If not set, default of 10 seconds will be used.

### `wordcounter`

The package provides a function to prepare word counts for display by sorting the input values and providing standard Stringer interface.

## How to run

### Exercise 1

`go run cmd/exercise_1/main.go`

### Exercise 2

`go run cmd/exercise_2/main.go`
