package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO conference"
	const conferenceTickets int = 50
	var remaningTickets uint = 50
	bookings := []string{} // slice of strings

	fmt.Printf("conferenceTickets is %T\n", conferenceTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remaningTickets {
			fmt.Printf("We have only %v tickets, so you can't book %v tickets\n", remaningTickets, userTickets)
			continue
		}

		remaningTickets = remaningTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remain for %v\n", remaningTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsLeft := remaningTickets == 0

		if noTicketsLeft {
			// end program
			fmt.Println("Our conference is sold out. Come back next year")
			break
		}
	}
}
