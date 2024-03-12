package main

import "fmt"

func main() {
	// this is comment
    // print output "Hello World!"
    // fmt.Println("Hello World!");

	/*
        this is comment
        print output "Hello World!"
        fmt.Println("Hello World!");
    */

	fmt.Println("Hello Go!")
	fmt.Println("Herly", 21, true) // Multiple output

	// Output without enter after
	fmt.Print("This is output Print ")
    fmt.Print("Hello Go!")
	fmt.Print("This is output Print\n") // \n meaning enter or add new line after this line

	// Output with format
	fmt.Printf("My name is: %s \n", "Herly Riyanto Hidayat")
	fmt.Printf("My age is: %d tahun\n", 21,)

	// Input with scan
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello, ", name) // Output from input user

	// Multiple input 
	var nama, address string
    fmt.Print("Enter your name and address : ")
    fmt.Scan(&nama, &address) // set input to variable 'name' and 'address'

    fmt.Println("Hello", nama)
    fmt.Println("Your address", address)

	// Scanln
	var yourName, yourAddress string
    fmt.Print("Enter your name : ")
    fmt.Scanln(&yourName)
	fmt.Print("\n")
    fmt.Print("Enter your address : ")
    fmt.Scanln(&yourAddress)

    fmt.Println("Hello", yourName)
    fmt.Println("Your address", yourAddress)

	// Scanf
	var namamu, alamatmu string

    fmt.Print("Enter your name and address : ")
    fmt.Scanf("%s %s", &namamu, &alamatmu) // set input with format string to variable name with multiple input

    fmt.Println("Hello", namamu)
    fmt.Println("Your address", alamatmu)
}