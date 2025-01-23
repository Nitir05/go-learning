package main

import (
	"fmt"
)

func main() {
	conferenceName := "🎉 Go Conference 🎉"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Positive integer only

	fmt.Println("══════════════════════════════════════════════════════")
	fmt.Printf("✨ Welcome to the %v booking application! ✨\n", conferenceName)
	fmt.Printf("🎫 We have a total of %v tickets, and %v are still available. 🎫\n", conferenceTickets, remainingTickets)
	fmt.Println("📢 Get your tickets here to attend this amazing event! 📢")
	fmt.Println("══════════════════════════════════════════════════════")

	bookings := []string{} // Initialize an empty slice

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for details
	fmt.Println("📝 Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("📝 Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("📧 Please enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("🎟️ How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

	// Update remaining tickets and bookings
	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName);

	fmt.Println("══════════════════════════════════════════════════════")
	fmt.Printf("🎉 Thank you, %v %v! 🎉\n", firstName, lastName)
	fmt.Printf("✅ You have successfully booked %v tickets.\n", userTickets)
	fmt.Printf("📧 A confirmation email will be sent to %v.\n", email)
	fmt.Println("══════════════════════════════════════════════════════")
	fmt.Printf("📉 Remaining tickets: %v\n", remainingTickets)
	fmt.Println("══════════════════════════════════════════════════════")
	fmt.Printf("📝 These are all our current bookings: %v\n", bookings)
	fmt.Println("══════════════════════════════════════════════════════")
}