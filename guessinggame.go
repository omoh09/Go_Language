package main

import "fmt"

func main() {
	var input int
	fmt.Println("Input decimal")
	fmt.Scan(&input)
	fmt.Printf("The binary is %b\n", input)
}