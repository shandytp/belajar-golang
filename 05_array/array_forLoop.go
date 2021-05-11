package main

import "fmt"

func main() {
	// Cara 1, For Looping biasa

	names := [7]string{
		"Poppin' Party",
		"Afterglow",
		"Pastel*Palletes",
		"Roselia",
		"Hello Happy World",
		"Morfonica",
		"RAISE A SUILEN",
	}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, names[i])
	// }

	// Cara 2, For Looping dengan range

	/*
		for i, name := range names {
			fmt.Printf("elemen %d : %s\n", i, name)
		}
	*/

	// Cara 3, For Looping dengan range variable underscore.
	// Karena range return index dan elemen, jika ingin elemennya saja bisa di pass dengan underscore

	for _, name := range names {
		fmt.Printf("nama band : %s\n", name)
	}

}
