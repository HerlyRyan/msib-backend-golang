package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	var number int32 = 10

	var bigNum = int64(number) // convert int32 to int64

	var floatNum = float32(number) // convert int32 to float32

	fmt.Println(bigNum)
	fmt.Println(floatNum)

	var numbers int32 = 10
    var isMaried bool = false

	// Sprintf
    var numString = fmt.Sprintf("%d", numbers) // convert int32 ke string
    var isMariedString = fmt.Sprintf("%t", isMaried) // convert bool ke string
    var Pi = fmt.Sprintf("%f", math.Pi) // convert float64 dari variable 'Pi' di package 'math' ke string

    fmt.Println(numString)
    fmt.Println(isMariedString)
    fmt.Println(Pi)

    fmt.Println("type numString:", reflect.TypeOf(numString))
    fmt.Println("type isMarriedString:", reflect.TypeOf(isMariedString))
    fmt.Println("type Pi:", reflect.TypeOf(Pi))

	// strconv.Itoa, convert integer to string
	var str = strconv.Itoa(10) // convert string from int

    fmt.Println(str)
    fmt.Println("type:", reflect.TypeOf(str))

	// strconv.Atoi, conver string to integer
	var num, _ = strconv.Atoi("10") // convert string to int

    fmt.Println(num)
    fmt.Println("type:", reflect.TypeOf(num))
}