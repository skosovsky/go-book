package main

import "testing"

var a uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount(a)
	}
}

func BenchmarkPopCountStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountStandard(a)
	}
}
func BenchmarkPopCountFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountFull(a)
	}
}
func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountByClearing(a)
	}
}
func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountByShifting(a)
	}
}
func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitCount(a)
	}
}
