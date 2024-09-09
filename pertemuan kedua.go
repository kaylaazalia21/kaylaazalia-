package main

import "fmt"

func main() {

	//deklarasi variabel
	var nama string = "Kayla Azalia"
	var umur int = 18
	var alamat string = "Jalan Kramat I Nomor 1"
	var berkacamata bool = false
	var minuskiri float32 = 0.50
	var minuskanan float32 = 0.50

	//proses
	var hasilminus float32 = minuskiri + minuskanan

	//output
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(alamat)
	fmt.Println(berkacamata)
	fmt.Println(minuskiri)
	fmt.Println(minuskanan)
	fmt.Println(hasilminus)

}
