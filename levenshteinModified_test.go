package main

import "testing"

func TestLevenshteinModifiedOne(t *testing.T) {
  actual := LevenshteinModified("Haus", "Maus", 2)
  expected := 1
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinModifiedTwo(t *testing.T) {
  actual := LevenshteinModified("Haus", "Mausi", 2)
  expected := 2
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinModifiedThree(t *testing.T) {
  actual := LevenshteinModified("Haus", "Häuser", 2)
  expected := 3
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinModifiedFour(t *testing.T) {
  actual := LevenshteinModified("Kartoffelsalat", "Runkelrüben", 2)
  expected := 3
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func BenchmarkLevenshteinModifiedTestCases(b *testing.B) {
  testCases := []struct {
    name string
    token1 string
    token2 string
    maxDist int
  }{
    {"Case One", "Haus", "Maus", 2},
    {"Case Two", "Haus", "Mausi", 2},
    {"Case Three", "Haus", "Häuser", 2},
    {"Case Four", "Kartoffelsalat", "Runkelrüben", 2},
  }

  for _, testCase := range testCases {
    b.Run(testCase.name, func(b *testing.B) {
      for i:=0; i<b.N; i++ {
        LevenshteinModified(testCase.token1, testCase.token2, testCase.maxDist)
      }
    })
  }

}
