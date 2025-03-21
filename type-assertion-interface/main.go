package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	fmt.Println(e.(email))
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
    email := email{
        isSubscribed: false,
        body:         "Hello there",
        toAddress:    "test@example.com",
    }
    
    sms := sms{
        isSubscribed:  true,
        body:         "Hi there",
        toPhoneNumber: "555-555-5555",
    }

    emailType, emailCost := getExpenseReport(email)
    fmt.Printf("Expense type: %s, Cost: $%.2f\n", emailType, emailCost)

    smsType, smsCost := getExpenseReport(sms)
    fmt.Printf("Expense type: %s, Cost: $%.2f\n", smsType, smsCost)
}