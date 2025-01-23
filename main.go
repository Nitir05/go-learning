package main

import (
	"fmt"
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

	// Declare variables to store user input for booking details
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask the user for their details (first name, last name, email, number of tickets)
	fmt.Println("📝 Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("📝 Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("📧 Please enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("🎟️ How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

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
	fmt.Printf("📝 These are all our current bookings: %v\n", bookings)
	fmt.Println("══════════════════════════════════════════════════════")
}
