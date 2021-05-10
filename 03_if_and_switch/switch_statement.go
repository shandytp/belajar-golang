package main

import "fmt"

func main() {
	var hari string = "Rabu"

	switch hari {
	case "Senin", "Selasa", "Kamis", "Jum'at":
		fmt.Println("Angel wes, wes angel...")
	case "Rabu":
		fmt.Println("It's Wednesday my dudeeeeee")
	default:
		fmt.Println("WEEKEND BOIIIIIIIIIIIISSSSSS")
	}

}
