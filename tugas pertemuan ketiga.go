package main

import "fmt"

func main() {

	var angka1, angka2 float64
	var operator string

	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan angka kedua : ")
	fmt.Scanln(&angka2)

	fmt.Print("Masukkan operator ( + - * /) : ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1+angka2)

	case "-":
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1-angka2)

	case "*":
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1*angka2)

	case "/":
		if angka2 == 0.0 {
			fmt.Println("Hasil tidak terdefinisi")
		} else {
			fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1/angka2)
		}
	default:
		fmt.Println("Invalid Operator")

	}

}
