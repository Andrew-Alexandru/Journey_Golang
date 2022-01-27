package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	//applies to variables not to constants
	//var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	//%T - stands for Type and %v stands for value
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for name

	userName = "Harry"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
