// Конвертирует числовой аргумент в температуру, длину и вес
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/skosovsky/go-book/chapter-2/lenconv"
	"github.com/skosovsky/go-book/chapter-2/tempconv"
	"github.com/skosovsky/go-book/chapter-2/weightconv"
)

func main() {
	var userNum []string

	if len(os.Args) == 1 {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		text = text[:len(text)-1]
		userNum = strings.Split(text, " ")
	} else {
		userNum = os.Args[1:]
	}

	for _, arg := range userNum {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fT := tempconv.Fahrenheit(val)
		cT := tempconv.Celsius(val)
		kT := tempconv.Kelvin(val)
		fmt.Printf("%.2f°F = %.2f°C, %.2f°C = %.2f°F, %.2fK = %.2f°C\n",
			fT, tempconv.FoC(fT), cT, tempconv.CoF(cT), kT, tempconv.KoC(kT))

		fL := lenconv.Foot(val)
		mL := lenconv.Meter(val)
		fmt.Printf("%.2fm. = %.2ff., %.2ff. = %.2fm.\n",
			fL, lenconv.FoM(fL), mL, lenconv.MoF(mL))

		pW := weightconv.Pound(val)
		kW := weightconv.Kilogram(val)
		fmt.Printf("%.2flb. = %.2fkg., %.2fkg. = %.2flb.\n",
			pW, weightconv.PoK(pW), kW, weightconv.KoP(kW))
	}
}
