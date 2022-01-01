package main

import "fmt"

func main() {

	var input string
	fmt.Println("Enter input")

	fmt.Scanln(&input) //reference to input variable

	fmt.Printf("Hello World! %s", input)
}
