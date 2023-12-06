package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	newData := []string{"one", "one", "two", "three", "three", "three"}
	fmt.Printf("%q\n", noduplicates(newData))
	fmt.Printf("%q\n", newData)

	bytes := []byte("Hello,   world!")
	fmt.Printf("%q\n", nospaces(bytes))
	fmt.Printf("%q\n", bytes)

}

func nonempty(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noduplicates(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	i := 0
	for j := 1; j <= len(strings)-1; j++ {
		if strings[i] != strings[j] {
			i++
			strings[i] = strings[j]
		}
	}

	return strings[:i+1]
}

func nospaces(bytes []byte) []byte {
	if len(bytes) == 0 {
		return bytes
	}

	result := bytes[:0]
	var lastRune rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			result = append(result, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(lastRune) {
			result = append(result, ' ')
		}
		lastRune = r
		i += s
	}
	return result
}
