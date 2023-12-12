// Вычисляет количество символов Unicode
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Количество символов Unicode
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8
	nums := 0                       // Количество цифр
	letters := 0                    // Количество букв
	other := 0                      // Количество символов других категорий

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны, байтов, ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			nums++
		} else if unicode.IsLetter(r) {
			letters++
		} else {
			other++
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, c := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, c)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
	if nums > 0 {
		fmt.Printf("\n%d цифр UTF-8\n", nums)
	}
	if letters > 0 {
		fmt.Printf("\n%d букв UTF-8\n", letters)
	}
	if other > 0 {
		fmt.Printf("\n%d других символов UTF-8\n", other)
	}
}
