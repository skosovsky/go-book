// Выводит аргументы консоли
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Printf("%vs elapsed\n", time.Since(start).Microseconds()) // fasted > 10 args

	start = time.Now()
	for idx, arg := range os.Args[:] {
		fmt.Println(idx, arg)
	}
	fmt.Printf("%vs elapsed\n", time.Since(start).Microseconds()) // fasted < 10 arg
}
