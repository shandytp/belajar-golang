package main

import "fmt"

func main() {
	value1 := "Ini adalah nilai"

	value2 := value1

	// Dua line dibawah menghasilkan output yang sama

	fmt.Println(value1)
	fmt.Println(value2)

	// Tetapi memiliki value memory address yang berbeda
	fmt.Println(&value1) // Output = 0xc000050240
	fmt.Println(&value2) // Output = 0xc000050250
}
