package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func main() {
	var name string = "Alfaizuna"

	fmt.Println(greet(name))
}