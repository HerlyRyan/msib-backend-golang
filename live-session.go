package main

import (
	"fmt"
)

func triangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := i; k <= n; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	triangle(10)
}