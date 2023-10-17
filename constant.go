package main

import "fmt"

func main() {
	const firstName string = "Iqbal"
	const lastName = "Fauzi"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	//error
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"
}
