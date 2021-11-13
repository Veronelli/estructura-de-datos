package main

import (
	"fmt"
	"strings"

	miPackage "./my-package"
)

func isPalindromo(texto string) {
	texto = strings.ToLower(texto)
	var textReverse string
	for i := len(texto) - 1; i >= 0; i-- {
		textReverse += string(texto[i])
	}
	if texto == textReverse {
		fmt.Println("Es Palidromo")

	} else {
		fmt.Println("No es palidromo")

	}

}

type Car struct {
	brand string
	year  int
}

func main() {
	//Array
	var array [4]int

	for i := 0; i < len(array); i++ {
		array[i] = i

	}

	fmt.Println(array, cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4])

	slice = append(slice, 7)
	slice = append(slice, 8)

	fmt.Println(slice)

	//Append nueva lista descompresion
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)

	fmt.Println(slice)

	slice1 := []string{"hola", "que", "hace"}

	for _ /*i*/, valor := range slice1 {
		fmt.Println(valor)

	}

	isPalindromo("amA")

	//Mapas
	m := make(map[string]int)
	m["Facu"] = 10
	m["Pepito"] = 20

	fmt.Println(m)

	for i, valor := range m {
		fmt.Println(i, valor)
	}

	facu, ok := m["Facu"]
	fmt.Println(facu, ok)

	//Struct
	var car Car = Car{brand: "Fiat", year: 2021}
	fmt.Println(car)

	car.brand = "Ford"
	fmt.Println(car)

	var car2 Car
	car2.brand = "Mercedez"
	car2.year = 2021

	fmt.Println(car2)

	//------------------------ import package

	var myCar miPackage.Car
	myCar.Brand = "Fort"
	myCar.Year = 2000
	fmt.Println(myCar)
	fmt.Println(miPackage.Saludar("Facu"))

}
