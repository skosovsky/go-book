package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int) // Количество символов Unicode

	in := bufio.NewScanner(os.Stdin) // cat sample.txt | go run main.go
	in.Split(bufio.ScanWords)

	for in.Scan() {
		counts[in.Text()]++
	}

	err := in.Err()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for i, v := range counts {
		fmt.Println(i, v)
	}

}
