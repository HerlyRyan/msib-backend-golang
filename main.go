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
}