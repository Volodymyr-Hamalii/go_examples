// https://youtu.be/yyUHQIec83I?si=S-Ize7JqZ-mR-oxN&t=3761
package main

import "fmt"

func main() {
	fmt.Println("Welcome!")

	var userName string = GetName()

	const maxTickets uint8 = 50
	var remainingTickets uint8 = 50

	//var bookings [50]string

	var bookedTickets = maxTickets - remainingTickets
	fmt.Printf("%v have booked %v tickets", userName, bookedTickets)


}

