package main

import "fmt"

func main() {
	var p *int // untuk membuat pointer tambahkan asterisk "*" lalu diikuti tipe data

	i := 42
	p = &i // mengarahkan value ke i
	makanan := "pecel"

	fmt.Println(&makanan) // baca atau access memory address pakai "&", output = 0xc000050240
	fmt.Println(*p)       // membaca variable i dengan pointer p , output = 42

	fmt.Println(&p) // output = 0xc000006028
	*p = 21         // set value baru lewat pointer p

	fmt.Println(i)  // output = 21
	fmt.Println(&p) // output = 0xc000006028
}
