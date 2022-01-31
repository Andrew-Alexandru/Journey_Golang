package main

import "fmt"

//format package - fmt

func main() {
	conferenceName := "Go Conference"
	//applies to variables not to constants
	//var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookings = [50]string{}
	//example of empty slice var bookings = []string{} or bookings := []string{}
	var bookings []string
	//%T - stands for Type and %v stands for value

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	// example of slices in GO
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole slice %v \n", bookings)
	fmt.Printf("The first value %v \n", bookings[0])
	fmt.Printf("slice type %T\n", bookings)
	fmt.Printf("slice size: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
