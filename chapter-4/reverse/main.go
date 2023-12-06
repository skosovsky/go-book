package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var nums = [6]int{1, 2, 3, 4, 5, 6}
	reverse(&nums)
	fmt.Println(nums)

	rightS := rotate([]int{1, 2, 3, 4, 5, 6}, 4, "right")
	fmt.Println(rightS)

	bytes := []byte("Hello, world!")
	reverseBytes(bytes)
	fmt.Printf("%q\n", bytes)
}

func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int, direct string) []int {
	if n > len(s) {
		n = n % len(s)
	}
	if direct == "left" {
		n = len(s) - n
	}

	temp := make([]int, n)
	copy(temp, s[len(s)-n:])
	copy(s, s[:len(s)-n])
	s = append(temp, s[:len(s)-n]...)

	return s
}

func reverseBytes(b []byte) {
	buf := make([]byte, 0, len(b))
	i := len(b)

	for i > 0 {
		_, s := utf8.DecodeLastRune(b[:i])
		buf = append(buf, b[i-s:i]...)
		i -= s
	}
	copy(b, buf)
}
