package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(anagram("teslaaa", "esaalat"))
}

func anagram(a, b string) bool {
	if len(a) != len(b) || len(a) == 0 {
		return false
	}

	for v := range a {
		s := strconv.Itoa(v)
		if strings.Count(a, s) != strings.Count(b, s) {
			return false
		}
	}

	return true
}
