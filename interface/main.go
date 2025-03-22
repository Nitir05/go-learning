package main

import (
	"fmt"
	"time"
)

func (e email) cost() int {
    if e.isSubscribed {
        return 2 * len(e.body)
    } else {
        return 5 * len(e.body)
    }
}

func (e email) format() string {
    if e.isSubscribed {
        return fmt.Sprintf("'%s' | Subscribed", e.body)
    } else {
        return fmt.Sprintf("'%s' | Not Subscribed", e.body)
    }
}

type expense interface {
    cost() int
}

type formatter interface {
    format() string
}

type email struct {
    isSubscribed bool
    body         string
}

func sendMessage(msg message) (string, int) {
    msgText := msg.getMessage()
    msgLen := len(msgText)

    return msgText, 3 * msgLen
}

type message interface {
    getMessage() string
}

type employee interface {
    getName() string
    getSalary() int
}

type contractor struct {
    name         string
    hourlyPay    int
    hoursPerYear int
}

func (c contractor) getName() string {
    return c.name
}

func (c contractor) getSalary() int {
    return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
    name   string
    salary int
}

func (ft fullTime) getSalary() int {
    return ft.salary
}

func (ft fullTime) getName() string {
    return ft.name
}

// don't edit below this line

type birthdayMessage struct {
    birthdayTime  time.Time
    recipientName string
}

func (bm birthdayMessage) getMessage() string {
    return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
    reportName    string
    numberOfSends int
}

func (sr sendingReport) getMessage() string {
    return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
    // Create instances of contractor and fullTime
    contractor1 := contractor{name: "John Doe", hourlyPay: 50, hoursPerYear: 2000}
    fullTime1 := fullTime{name: "Jane Smith", salary: 80000}

    // Call getName and getSalary for contractor
    fmt.Printf("Contractor Name: %s\n", contractor1.getName())
    fmt.Printf("Contractor Salary: %d\n", contractor1.getSalary())

    // Call getName and getSalary for fullTime
    fmt.Printf("FullTime Name: %s\n", fullTime1.getName())
    fmt.Printf("FullTime Salary: %d\n", fullTime1.getSalary())

    // Create instances of birthdayMessage and sendingReport
    birthdayMsg := birthdayMessage{birthdayTime: time.Now(), recipientName: "Alice"}
    sendingRpt := sendingReport{reportName: "Monthly Report", numberOfSends: 150}

    // Call sendMessage for birthdayMessage and sendingReport
    birthdayText, birthdayLen := sendMessage(birthdayMsg)
    fmt.Printf("Birthday Message: %s\n", birthdayText)
    fmt.Printf("Birthday Message Length: %d\n", birthdayLen)

    sendingText, sendingLen := sendMessage(sendingRpt)
    fmt.Printf("Sending Report Message: %s\n", sendingText)
    fmt.Printf("Sending Report Message Length: %d\n", sendingLen)

    // Create instance of email
    email1 := email{isSubscribed: true, body: "Welcome to our newsletter!"}

    // Call cost and format for email
    fmt.Printf("Email Cost: %d\n", email1.cost())
    fmt.Printf("Email Format: %s\n", email1.format())
}