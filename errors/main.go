package main

import (
	"errors"
	"fmt"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string{
    return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		 err := errors.New(fmt.Sprintf("can not divide %v by zero", divisor))
         return 0.0, err
	}
	return dividend / divisor, nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
    custCost, custErr := sendSMS(msgToCustomer)
    if custErr != nil {
        return 0, custErr
    }

    spouseCost, spouseErr := sendSMS(msgToSpouse)
    if spouseErr != nil {
        return 0, spouseErr
    }

    return custCost + spouseCost, nil
}

func validateStatus(status string) error {
	statusLen := len(status)
    if statusLen == 0 {
        return errors.New("Status cannot be empty")
    }

    if statusLen > 140 {
        return errors.New("Status exceeds 140 characters")
    }

    return nil
}

func sendSMS(message string) (int, error) {
    const maxTextLen = 25
    const costPerChar = 2
    if len(message) > maxTextLen {
        return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
    }
    return costPerChar * len(message), nil
}

func getSMSErrorString(cost float64, recipient string) string {
    return fmt.Sprintf("SMS that costs %v to be sent to %v can not be sent", cost, recipient)
}

func main() {
    // Example 1: Sending a single SMS
    message := "Hello, welcome to our service!"
    cost, err := sendSMS(message)
    if err != nil {
        fmt.Printf("Error sending SMS: %s\n", getSMSErrorString(float64(cost), "Customer"))
    } else {
        fmt.Printf("SMS sent successfully. Cost: %d\n", cost)
    }

    // Example 2: Sending SMS to a couple
    msgToCustomer := "Hi Customer, your order is ready."
    msgToSpouse := "Hi Spouse, your partner's order is ready."
    totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
    if err != nil {
        fmt.Printf("Error sending SMS to couple: %s\n", getSMSErrorString(float64(totalCost), "Couple"))
    } else {
        fmt.Printf("SMS sent to couple successfully. Total Cost: %d\n", totalCost)
    }

    // Example 3: Sending an SMS that exceeds the character limit
    longMessage := "This message is way too long to be sent in a single SMS because it exceeds the character limit."
    cost, err = sendSMS(longMessage)
    if err != nil {
        fmt.Printf("Error sending long SMS: %s\n", getSMSErrorString(float64(cost), "Customer"))
    } else {
        fmt.Printf("Long SMS sent successfully. Cost: %d\n", cost)
    }
}