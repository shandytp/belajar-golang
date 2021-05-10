package main

import "fmt"

func main() {
	var nilai float32 = 9.5

	if nilai == 10 {
		fmt.Println("Lulus dengan sempurna")
	} else if nilai < 10 {
		fmt.Println("Cukup mantap")
	} else if nilai == 7 {
		fmt.Println("Mungkin bisa belajar lagi")
	} else {
		fmt.Println("You stupid. Nilai mu", nilai)
	}
}
