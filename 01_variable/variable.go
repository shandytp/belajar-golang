package main

import "fmt"

func main() {
	// var makanan, minuman, gorengan string = "pecel", "es teh", "tempe mendoan" // bisa assign value pakai tipe data
	makanan, minuman, gorengan := "pecel", "es teh", "tempe mendoan" // bisa assign value langsung tanpa inisialisasi tipe data
	hargaMakanan, hargaMinuman, hargaGorengan := 10000, 3000, 2000

	fmt.Println("Nama makanannya", makanan, "Nama minumannya", minuman, "Nama gorengannya", gorengan)
	fmt.Println("Harga makanannya", hargaMakanan, "Harga minumannya", hargaMinuman, "Harga gorengannya", hargaGorengan)
}
