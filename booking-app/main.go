package main

import "fmt"

func main() {
	fmt.Println("Welcome!")
	fmt.Println("Print your name:")

	var userName string
	fmt.Scan(&userName)

	fmt.Printf("Okay, your name is %v.", userName)

}

