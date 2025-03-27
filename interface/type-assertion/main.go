package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
		case email:
            return v.toAddress, v.cost()
        case sms:
            return v.toPhoneNumber, v.cost()
        default:
            return "", 0.0
	}
}

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}

func (p PlainText) Format() string {
	return p.message
}

func (b Bold) Format() string {
	return fmt.Sprintf("**%s**", b.message);
}

func (c Code) Format() string {
	return fmt.Sprintf("`%s`", c.message);
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
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