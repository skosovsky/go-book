// Выводит текст и количество появлений каждой строки, которая появляется
// в консоли или файлах более 1 раза
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fileInfo struct {
	count    int
	filename []string
}

func main() {
	counts := make(map[string]fileInfo)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, f.Name(), counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, strings.Join(n.filename, " "))
		}
	}
}

func countLines(f *os.File, fname string, counts map[string]fileInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}

		currCount := counts[input.Text()]
		currCount.count += 1

		if len(currCount.filename) == 0 {
			currCount.filename = append(currCount.filename, fname)
		} else {
			nameCount := 0
			for _, name := range currCount.filename {
				if name == fname {
					nameCount++
				}
			}
			if nameCount == 0 {
				currCount.filename = append(currCount.filename, fname)
			}

		}

		counts[input.Text()] = currCount
	}
}

//func main() {
//	counts := make(map[string]int)
//
//	for _, filename := range os.Args[1:] {
//		data, err := os.ReadFile(filename)
//
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
//			continue
//		}
//
//		for _, line := range strings.Split(string(data), "\n") {
//			counts[line]++
//		}
//	}
//
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}
