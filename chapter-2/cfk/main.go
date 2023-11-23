// Конвертирует числовой аргумент в температуру по Цельсию, Фаренгейту и Кельвину
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/skosovsky/go-book/chapter-2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%.2f°F = %.2f°C, %.2f°C = %.2f°F, %.2fK = %.2f°C\n",
			f, tempconv.FoC(f), c, tempconv.CoF(c), k, tempconv.KoC(k))
	}
}
