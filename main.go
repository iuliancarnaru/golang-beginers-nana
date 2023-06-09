package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// global variables
const conferenceTickets int = 50

var conferenceName string = "GO conference"

var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) // a list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wgr = sync.WaitGroup{}

func main() {

	// fmt.Printf("conferenceTickets is %T\n", conferenceTickets)
	greetUsers()

	// for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidOrder := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidOrder {
		bookTicket(userTickets, firstName, lastName, email)

		// concurrency mode
		wgr.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsLeft := remainingTickets == 0

		if noTicketsLeft {
			// end program
			fmt.Println("Our conference is sold out. Come back next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("Your name or last name are too short.")
		}

		if !isValidEmail {
			fmt.Println("Your email address is invalid")
		}

		if !isValidOrder {
			fmt.Println("Number of tickets you enter is invalid")
		}

	}
	wgr.Wait()
}

// }

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user (supports only one type which is initially declared)
	// var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("bookings %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remain for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-------------------------------")
	fmt.Printf("Sending ticket:\n %v \nto %v\n", ticket, email)
	fmt.Println("-------------------------------")
	wgr.Done()
}
