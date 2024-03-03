// https://youtu.be/yyUHQIec83I?si=S-Ize7JqZ-mR-oxN&t=3761
package main

import "fmt"

func getName() string {
	fmt.Println("Print your name:")

	var userName string
	fmt.Scan(&userName)

	fmt.Printf("Okay, your name is %v.", userName)

	return userName
}

func main() {
	fmt.Println("Welcome!")

	//var userName string = getName()

	const maxTickets uint8 = 50
	var remainingTickets uint = 50

	var bookings [50]string



}

