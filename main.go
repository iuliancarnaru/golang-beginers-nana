package main

import "fmt"

func main() {
	conferenceName := "GO conference"
	const conferenceTickets int = 50
	var remaningTickets uint = 50
	bookings := []string{} // slice of strings

	fmt.Printf("conferenceTickets is %T\n", conferenceTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")

	for i := 0; i < 5; i++ {
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

		remaningTickets = remaningTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remain for %v\n", remaningTickets, conferenceName)

		fistNames := []string{}
		fmt.Printf("There are all our bookings: %v\n", bookings)
	}

}
