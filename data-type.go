package main

import "fmt"

func main() {
	// String
	fmt.Println(len("Herly Riyanto Hidayat"))
	fmt.Println("Golang"[0]) // Output with ASCII code
	fmt.Println(string("Golang"[0])) // Output G

	// Numeric
	fmt.Println(123) //int or uint
	fmt.Println(-123) //int (uint can't have minus number)
	fmt.Println(10.5) // float

	// Numeric operation
	fmt.Println(1 + 1)
    fmt.Println(2 - 1)
    fmt.Println(2 * 2)
    fmt.Println(4 / 2)
    fmt.Println(10 % 3) // sisa pembagian 10 dengan 3 adalah 1

	// Augmented assignment
	var a = 10
    a += 10 // a = 10 + 10 = 20
    fmt.Println(a)

	// Unary operator
	var b = 10
    b++ // a = 10 + 1 = 11

    fmt.Println(a)

    var c = 10
    fmt.Println(-c)

	// Boolean
	fmt.Println("Benar =", true)
    fmt.Println("Salah =", false)
}