package main

import "fmt"

func Looping(input int) float64 {
	result := 0.0
	for i := 1.0; i <= float64(input); i+= 0.5 {
		result += i
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(Looping(10))

	total := 0
	for i := 1; i <= 5 ; i++ { 
        total += i // 1 + 2 + 3 + 4 + 5 = 15
    }
	fmt.Println(total)

	// 
	i := 0 // init statement
    for i < 5 {
        fmt.Println("Hello Go!", i+ 1)
        i++ // post statement
    }
}
