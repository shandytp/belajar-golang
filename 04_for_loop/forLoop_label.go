package main

import "fmt"

func main() {
	// Diberi label, namanya terserah
outerLoop:
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
