package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(comma("555555"))
	fmt.Println(commaByte("55555555555555"))
	fmt.Println(commaFloat(-55555555.155555))
}

// comma вставляет запятые в строковое представление неотрицательного числа
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaByte(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i%3 == 0 && i != 0 && i != len(s)-1 {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func commaFloat(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	// определяем отрицательное ли число и точку старта
	isMinus := strings.HasPrefix(s, "-")
	start := 0
	if isMinus {
		start = 1
	}
	// определяем есть ли точка и точку финиша
	dot := strings.Index(s, ".")
	finish := len(s) - 1
	if dot != -1 {
		finish = dot
	}
	// пробегаемся по числу
	var buf bytes.Buffer
	for i, v := range s[start:finish] {
		if i%3 == 0 && i != 0 && i != len(s)-1 {
			buf.WriteString(",")
		}
		// если отрицательное, добавляем минус
		if isMinus && i == 0 {
			buf.WriteString("-")
		}
		buf.WriteRune(v)
	}
	// добавляем значения после точки
	if dot != -1 {
		buf.WriteString(s[dot:])
	}
	return buf.String()
}
