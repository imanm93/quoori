# Quoori Programming Challenge

This is a solution to the Quoori programming challenge

## Project Files

There are

## Test

To run the tests please enter the following command from your terminal.

```
go test
```

### With Benchmarks

To view the average execution time of each variant of the Levenshtein distance algorithms for each test case, please run the following command from your terminal.

```
go test -bench=.
```

## Usage

Please create a `main.go` file in the root directory and access the public functions `Levenshtein` or `LevenshteinModified` to use them in your program. (An example is shown below)

```go

package main

import "fmt"

func main() {

  stringOne := "ExampleOne"
  stringTwo := "ExampleTwo"
  maxDist := 1

  resultLevenshtein := Levenshtein(stringOne, stringTwo)
  fmt.Printf("The Levenshtein distance between '%v' and '%v' without a max limit is %v\n", stringOne, stringTwo, resultLevenshtein)

  resultLevenshteinModified := LevenshteinModified(stringOne, stringTwo, maxDist)
  fmt.Printf("The Levenshtein distance between '%v' and '%v' with a max limit is %v\n", stringOne, stringTwo, resultLevenshteinModified)

}

```
