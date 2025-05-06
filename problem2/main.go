package main

import (
	"fmt"
)

func main() {
	
	for i := 1; i <= 5; i++ {
		if i <= 3 {
			fmt.Println(i, "is Small")
		} 

		if i > 3 {
			fmt.Println(i, "is Large")
		}
	}
}

