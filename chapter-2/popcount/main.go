package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var a uint64 = 11111155
	fmt.Println(popCount(a))
	fmt.Println(popCountStandard(a))
	fmt.Println(popCountFull(a))
	fmt.Println(popCountByClearing(a))
	fmt.Println(popCountByShifting(a))
	fmt.Println(bitCount(a))
}

// dataPrepare рассчитывает количество установленных бит для чисел от 0 до 255
func dataPrepare() [256]byte {
	//	pc[i] - количество единичных битов в i
	var pc [256]byte

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	return pc
}

// popCount возвращает степень заполнения (количество установленных битов) значения x
func popCount(x uint64) int {
	pc := dataPrepare()

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// popCountStandard возвращает степень заполнения (количество установленных битов) значения x стандартной функцией
func popCountStandard(x uint64) int {
	return bits.OnesCount64(x)
}

// popCountFull возвращает степень заполнения (количество установленных битов) значения x циклом
func popCountFull(x uint64) int {
	var pc [256]byte
	result := 0

	for i := 0; i <= 8; i++ {
		for x := 0; x <= int(byte(x>>(i*8))); x++ {
			pc[x] = pc[x/2] + byte(x&1)
		}

		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

// popCountByShifting возвращает степень заполнения (количество установленных битов) значения x
// с помощью сдвига по всем 64 позициям
func popCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// popCountByClearing возвращает степень заполнения (количество установленных битов) значения x
// с помощью использования выражения x&(x-1), которое сбрасывает крайний справа ненулевой бит x
func popCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x &= x - 1 // clear rightmost non-zero bit
		n++
	}
	return n
}

// bitCount возвращает степень заполнения (количество установленных битов) значения x
// с помощью масок из Hacker's Delight
func bitCount(x uint64) int {
	x -= x >> 1 & 0x5555555555555555
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x & 0x7f) //nolint:gomnd
}
