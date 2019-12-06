package main

func LevenshteinModified(token1 string, token2 string, maxDist int) int {

  // Convert to rune slice to handle Non-ASCII characters in the string
  token1Rune := []rune(token1)
  token2Rune := []rune(token2)

  token1Length := len(token1Rune)
  token2Length := len(token2Rune)

  // Initialise and assign values to first column
  column := make([]int, token1Length+1)
  for index := 1; index <= token1Length; index++ {
    column[index] = index
  }

  for row := 1; row <= token2Length; row++ {
    column[0] = row
    lastKey := row - 1
    minEditDistance := int(^uint(0) >> 1)
    for col := 1; col <= token1Length; col++ {
      oldKey := column[col]
      cost := 0
      if token1Rune[col-1] != token2Rune[row-1] {
        cost = 1
      }
      column[col] = Minimum(column[col], column[col-1], lastKey) + cost
      // Set current min edit distance
      if column[col] < minEditDistance {
        minEditDistance = column[col]
      }
      lastKey = oldKey
    }
    // Check if min edit distance > max distance allowed
    if minEditDistance > maxDist {
      return maxDist + 1
    }
  }

  return column[token1Length]

}
