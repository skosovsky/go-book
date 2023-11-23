// Package tempconv выполняет вычисления температур в Цельсиях, Фаренгейтах и Кельвинах
package tempconv

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
)

// CoF преобразует Celsius в Fahrenheit
func CoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) //nolint:gomnd
}

// CoK преобразует Celsius в Kelvin
func CoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15) //nolint:gomnd
}

// FoC преобразует Fahrenheit в Celsius
func FoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) //nolint:gomnd
}

// FoK преобразует Fahrenheit в Kelvin
func FoK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15) //nolint:gomnd
}

// KoC преобразует Kelvin в Celsius
func KoC(k Kelvin) Celsius {
	return Celsius(k - 273.15) //nolint:gomnd
}

// KoF преобразует Kelvin в Fahrenheit
func KoF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32) //nolint:gomnd
}
