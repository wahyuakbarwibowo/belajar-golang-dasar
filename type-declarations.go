package main

import "fmt"

func typeDeclaration() {
	type NoKTP = string
	type Married = bool

	var noKtpEko NoKTP = "12312312140128031"
	var marriedStatus Married = false

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
