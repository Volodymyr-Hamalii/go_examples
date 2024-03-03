package main

import "fmt"

func GetName() string {
	fmt.Println("Print your name:")

	var userName string
	fmt.Scan(&userName)

	//fmt.Printf("Okay, your name is %v.", userName)

	return userName
}