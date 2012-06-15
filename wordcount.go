package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(input string) map[string]int {
  returning := make(map[string]int)

  for _, word := range strings.Fields(input) {
    returning[word]++
  }

  return returning
}

func main() {
  wc.Test(WordCount)
}