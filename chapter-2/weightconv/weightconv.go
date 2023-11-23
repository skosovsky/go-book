// Package weightconv выполняет вычисления веса в Фунтах и Килограммах
package weightconv

type (
	Pound    float64
	Kilogram float64
)

// PoK преобразует Pound в Kilogram
func PoK(p Pound) Kilogram {
	return Kilogram(p / 2.20462) //nolint:gomnd
}

// KoP преобразует Kilogram в Pound
func KoP(k Kilogram) Pound {
	return Pound(k * 2.20462) //nolint:gomnd
}
