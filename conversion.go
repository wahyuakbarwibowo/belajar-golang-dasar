package main

import "fmt"

func conversion() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	const name = "Wahyu"
	var a uint8 = name[0]
	fmt.Println(string(a))
}
