package gonewrepo

import "fmt"



func MyFunc() {
	type Address struct {
		Name    string
		Street  string
		City    string
		State   string
		Pincode int
	}
	r := Address{Name: "raj", Street: "nrroad", City: "talala", State: "gujrat", Pincode: 362150}
	fmt.Println(r)

	
}