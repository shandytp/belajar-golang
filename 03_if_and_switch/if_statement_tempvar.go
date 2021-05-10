package main

import "fmt"

func main() {

	// Hampir sama dengan if statement sebelumnya, tetapi menggunakan Variable Temporary

	const budget int = 20000
	hargaMakanan, hargaMinuman := 10000, 3000

	if totalHarga := hargaMakanan + hargaMinuman; budget < totalHarga {
		fmt.Printf("Nice, kamu sudah berhemat. Sisa uang kamu adalah %d", budget-totalHarga)
	} else if totalHarga >= 10000 && totalHarga <= 15000 {
		fmt.Printf("Hmm, mungkin kamu waktunya berhemat. Karena sisa budget kamu adalah %d", budget-totalHarga)
	} else {
		fmt.Println("Hemat plis")
	}
}
