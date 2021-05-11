package main

import "fmt"

func main() {
	// Contoh penerapan array
	// Inisialisasi Cara 1

	/*
		var nama [7]string // inisialisasi = var nama_var [size_array]data_type
		nama[0] = "Poppin' Party"
		nama[1] = "Afterglow"
		nama[2] = "Pastel*Palletes"
		nama[3] = "Roselia"
		nama[4] = "Hello Happy World"
		nama[5] = "Morfonica"
		nama[6] = "RAISE A SUILEN"

		fmt.Println(nama[0], nama[1], nama[2], nama[3], nama[4], nama[5], nama[6])
	*/

	// Cara 2, inisialisasi Nilai awal Array Horizontal
	/*
		var nama = [7]string{"Poppin' Party", "Afterglow", "Pastel*Palletes", "Roselia", "Hello Happy World", "Morfonica", "RAISE A SUILEN"}
		fmt.Println("Element size :", len(nama))
		fmt.Println("isi elemen :", nama)
	*/

	// Cara 3, inisialisasi Nilai awal Array Vertikal
	/*
		nama := [7]string{
			"Poppin' Party",
			"Afterglow",
			"Pastel*Palletes",
			"Roselia",
			"Hello Happy World",
			"Morfonica",
			"RAISE A SUILEN",
		}
		fmt.Println("Element size :", len(nama))
		fmt.Println("isi elemen :", nama)
	*/

	// Cara 4, inisialisasi dynamic size array

	var numbers = [...]int{2, 3, 2, 1, 4, 5, 1, 9, 19}

	fmt.Println("Element size :", len(numbers))
	fmt.Println("isi elemen :", numbers)
}
