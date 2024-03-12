package main

import "fmt"

func main() {
	var hello string
	hello = "Hello Go! "
	var name string = "Herly Riyanto Hidayat"
	var (
        nama string = "Herly" // can like this
        age int
        isMarried bool
    ) // multiple declaration when different data type
	myname := "Herly Riyanto Hidayat" // short declaration
	const pi = 3.14 // contanta, cannot re-assign and must have a value when first declaration
	var namaku, alamatku string // for declaration with same data type

	fmt.Println(hello)
	fmt.Println(name)
	fmt.Println(nama) // default value = ""
	fmt.Println(age) // default value = 0
	fmt.Println(isMarried) // default value = false
	fmt.Println(myname) 
	fmt.Println(pi) 
	fmt.Println(namaku) 
	fmt.Println(alamatku) 

	// _ this variable use for trash variable, variabel ini tidak menyebabkan error ketika tidak digunakan
}