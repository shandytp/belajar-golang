package main

import "fmt"

func main() {
	// Cara 3, For looping tanpa kondisi. Seperti cara 1 tapi di breakdown di dalam for loop nya.
	var i int = 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 10 {
			break
		}
	}
}
