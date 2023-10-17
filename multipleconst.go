package main

import "fmt"

func main() {
	const (
		firstName string = "Iqbal"
		lastName         = "Fauzi"
		value            = 1000
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	//error
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"
}
