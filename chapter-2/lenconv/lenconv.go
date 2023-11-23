// Package lenconv выполняет вычисления длины в Футах и Метрах
package lenconv

type (
	Foot  float64
	Meter float64
)

// FoM преобразует Foot в Meter
func FoM(f Foot) Meter {
	return Meter(f / 3.28084) //nolint:gomnd
}

// MoF преобразует Meter в Foot
func MoF(m Meter) Foot {
	return Foot(m * 3.28084) //nolint:gomnd
}
