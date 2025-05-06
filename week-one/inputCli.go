package main

import (
	"fmt"
)

func main(){
	var name string

	fmt.Println("Masukkan Nama Anda: ")
	fmt.Scanln(&name)

	fmt.Println("Halo,", name)
}