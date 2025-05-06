package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("Number %d is even\n", i)
		} else {
			fmt.Printf("Number %d is odd\n", i)
		}
	}
}
