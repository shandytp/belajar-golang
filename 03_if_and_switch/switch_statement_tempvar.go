package main

import "fmt"

func main() {
	// Pada switch statement juga bisa menggunakan Temporary Variable

	switch hari := "Rabu"; hari {
	case "Senin", "Selasa", "Kamis", "Jum'at":
		fmt.Println("Angel wes, wes angel....")
	case "Rabu":
		fmt.Println("It's Wednesday my dudeeeeeeee")
	default:
		fmt.Println("WEKEEEND BOISSSSSSS")
	}
}
