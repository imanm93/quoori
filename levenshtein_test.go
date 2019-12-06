package main

import "testing"

func TestLevenshteinOne(t *testing.T) {
  actual := Levenshtein("Haus", "Maus")
  expected := 1
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinTwo(t *testing.T) {
  actual := Levenshtein("Haus", "Mausi")
  expected := 2
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinThree(t *testing.T) {
  actual := Levenshtein("Haus", "Häuser")
  expected := 3
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func TestLevenshteinFour(t *testing.T) {
  actual := Levenshtein("Kartoffelsalat", "Runkelrüben")
  expected := 12
  if actual != expected {
    t.Errorf("Actual value (%d) != Expected value (%d)", actual, expected)
  }
}

func BenchmarkLevenshteinTestCases(b *testing.B) {
  testCases := []struct {
    name string
    token1 string
    token2 string
  }{
    {"Case One", "Haus", "Maus"},
    {"Case Two", "Haus", "Mausi"},
    {"Case Three", "Haus", "Häuser"},
    {"Case Four", "Kartoffelsalat", "Runkelrüben"},
  }

  for _, testCase := range testCases {
    b.Run(testCase.name, func(b *testing.B) {
      for i:=0; i<b.N; i++ {
        Levenshtein(testCase.token1, testCase.token2)
      }
    })
  }

}
