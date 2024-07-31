package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func roro() {
	var names [3]string

	fmt.Println(names)
	fmt.Println(len(names))

	var values = [3]int16{
		90,
		92,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	name := "Wahjo"
	increment := func() {
		fmt.Println(name)
		name = "wayu"
	}

	increment()

	fmt.Println(name)

	var contohError error = errors.New("sesuatu error")
	fmt.Println(contohError)

	hasil, err := Pembagian(10, 0)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
