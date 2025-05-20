package main

import (
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
    lastIndex := make(map[rune]int) // menyimpan indeks terakhir kemunculan karakter
    maxLen := 0
    start := 0 // indeks awal window substring tanpa duplikat

    for i, ch := range s {
        if lastPos, found := lastIndex[ch]; found && lastPos >= start {
            start = lastPos + 1
        }
        lastIndex[ch] = i
        if i - start + 1 > maxLen {
            maxLen = i - start + 1
        }
    }
    return maxLen
}

func main() {
    testStrings := []string{
        "abcabcbb",
        "bbbbb",
        "pwwkew",
        "",
        "dvdf",
    }

    for _, s := range testStrings {
        fmt.Printf("Input: %q, Panjang substring tanpa duplikat terpanjang: %d\n", s, lengthOfLongestSubstring(s))
    }
}
