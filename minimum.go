package main

func Minimum(valueOne, valueTwo, valueThree int) int {
  if valueOne < valueTwo {
    if valueOne < valueThree {
      return valueOne
    }
  } else {
    if valueTwo < valueThree {
      return valueTwo
    }
  }
  return valueThree
}
