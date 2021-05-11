package main

import "fmt"

func main() {

	// Alokasi elemen array menggunakan `make`

	var makanan = make([]string, 2) // make([]data_type, size)
	makanan[0] = "pecel"
	makanan[1] = "ayam goreng"

	fmt.Println(makanan)
}
