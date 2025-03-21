package main

import (
	"fmt"
	"strings"
)

// main is the entry point of the ticket booking application.
// It allows users to book tickets for the "Go Conference" and manages ticket reservations.
func main() {
	// Define conference details
	conferenceName := "🎉 Go Conference 🎉"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Positive integer only

	// Display introductory message
	fmt.Println("══════════════════════════════════════════════════════")
	fmt.Printf("✨ Welcome to the %v booking application! ✨\n", conferenceName)
	fmt.Printf("🎫 We have a total of %v tickets, and %v are still available. 🎫\n", conferenceTickets, remainingTickets)
	fmt.Println("📢 Get your tickets here to attend this amazing event! 📢")
	fmt.Println("══════════════════════════════════════════════════════")

	// Initialize an empty slice for storing bookings
	bookings := []string{}

	for{
		

		// Declare variables to store user input for booking details
		var firstName string
		var lastName string
		var email string
		var city string
		var userTickets uint

		// Ask the user for their details (first name, last name, email, number of tickets)
		fmt.Println("📝 Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("📝 Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("📧 Please enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Please enter your city: ")
		fmt.Scan(&city)
		fmt.Println("🎟️ How many tickets would you like to book? ")
		fmt.Scan(&userTickets)

		//Validate username and email
		var isValidName bool= len(firstName) >= 2 && len(lastName) >= 1
		if !isValidName {
			fmt.Println("Please enter a valid name, First name should be 2 or more characters and Last name should be 1 or more characters")
			continue
		}
		
		var isValidAddress bool = strings.Contains(email, "@")
		if(!isValidAddress) {
			fmt.Println("Please enter a valid address, Email address should contain @ in the address");
			continue
		}

		var isValidTicketNumber bool = userTickets > 0 && userTickets<= remainingTickets
		if(!isValidTicketNumber) {
			fmt.Println("Please enter valid ticket number, tickets should be more than 1 and less than avaialable tickets")
		}

		isInvalidCity := city != "Singapore" || city != "London"
		if(!isInvalidCity) {
			fmt.Println("Please enter valid city, Our conference is only available in Singapore and London")
		}

		// Validate if the tickets are available 
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v remaining tickets, you cannot book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		// Update remaining tickets and append the booking details
		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		// Display confirmation message
		fmt.Println("══════════════════════════════════════════════════════")
		fmt.Printf("🎉 Thank you, %v %v! 🎉\n", firstName, lastName)
		fmt.Printf("✅ You have successfully booked %v tickets.\n", userTickets)
		fmt.Printf("📧 A confirmation email will be sent to %v.\n", email)
		fmt.Println("══════════════════════════════════════════════════════")
		
		// Display updated ticket count and current bookings
		fmt.Printf("📉 Remaining tickets: %v\n", remainingTickets)
		fmt.Println("══════════════════════════════════════════════════════")

		firstNames := []string{}

		for _, booking := range bookings {
			var firstname  = strings.Fields(booking)[0];
			firstNames = append(firstNames, firstname)
		}

	
		fmt.Printf("📝 These are all our current bookings: %v\n", firstNames)
		fmt.Println("══════════════════════════════════════════════════════")

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out, Come back next year!!")
			break
		}

	}
}
