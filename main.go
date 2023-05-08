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

	// fmt.Printf("conferenceTickets is %T\n", conferenceTickets)
	greetUsers(conferenceName, conferenceTickets, remaningTickets)

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidOrder := userTickets > 0 && userTickets <= remaningTickets

		if isValidName && isValidEmail && isValidOrder {
			remaningTickets = remaningTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remain for %v\n", remaningTickets, conferenceName)

			printFirstNames(bookings)

			noTicketsLeft := remaningTickets == 0

			if noTicketsLeft {
				// end program
				fmt.Println("Our conference is sold out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name or last name are too short.")
			}

			if !isValidEmail {
				fmt.Println("Your emaill address is invalid")
			}

			if !isValidOrder {
				fmt.Println("Number of tickets you enter is invalid")
			}

		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}

	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}
