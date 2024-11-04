package main

import "fmt"

type celcius struct {
	suhu float64
}
type fahrenheit struct {
	suhu float64
}
type kelvin struct {
	suhu float64
}

func (c celcius) toCelcius() float64 {
	return c.suhu
}

func (c celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32
}

func (c celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f fahrenheit) toFahrenheit() float64 {
	return f.suhu
}

func (f fahrenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

func (f fahrenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

func (k kelvin) toKelvin() float64 {
	return k.suhu
}

func (k kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k kelvin) toFahrenheit() float64 {
	return (k.suhu * (9.0 / 5.0)) - 459.67
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Masukkan Suhu awal: ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")

	fmt.Println("Masukkan suhu awal yang anda inginkan: ")

	var suhuAwal int
	fmt.Scanf("%d\n", &suhuAwal)
	for suhuAwal < 1 || suhuAwal > 3 {
		fmt.Println("Suhu awal tidak valid")
		fmt.Print("Masukkan suhu awal yang anda inginkan: ")
		fmt.Scanf("%d\n", &suhuAwal)
	}

	fmt.Println("Masukkan Suhu akhir: ")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")

	fmt.Println("Masukkan suhu akhir yang anda inginkan: ")

	var suhuAkhir int
	fmt.Scanf("%d\n", &suhuAkhir)
	for suhuAkhir < 1 || suhuAkhir > 3 {
		fmt.Println("Suhu akhir tidak valid")
		fmt.Print("Masukkan suhu akhir yang anda inginkan: ")
		fmt.Scanf("%d\n", &suhuAkhir)
	}

	var suhu float64
	fmt.Print("Masukkan suhu: ")
	fmt.Scanf("%f\n", &suhu)

	var interfaceSuhu hitungSuhu

	switch suhuAwal {
	case 1:
		interfaceSuhu = celcius{suhu}
	case 2:
		interfaceSuhu = fahrenheit{suhu}
	case 3:
		interfaceSuhu = kelvin{suhu}
	}

	var suhuFinal float64

	switch suhuAkhir {
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
	case 2:
		suhuFinal = interfaceSuhu.toFahrenheit()
	case 3:
		suhuFinal = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Suhu akhir adalah %.2f\n", suhuFinal)

}
