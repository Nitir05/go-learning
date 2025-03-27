package main

import "fmt"

func getMessageWithRetries(primaryMsg string, firstReminder string, lastReminder string) ([3]string, [3]int) {
    messages := [3]string{primaryMsg, firstReminder, lastReminder}
    costs := [3]int{len(primaryMsg), len(primaryMsg) + len(firstReminder), len(primaryMsg) + len(firstReminder) + len(lastReminder)}

    return messages, costs
}

func main() {
    // Example usage of getMessageWithRetries
    primaryMsg := "Your subscription is about to expire."
    firstReminder := "Please renew your subscription soon."
    lastReminder := "This is your final reminder to renew your subscription."

    messages, costs := getMessageWithRetries(primaryMsg, firstReminder, lastReminder)

    fmt.Println("Messages and their cumulative costs:")
    for i := 0; i < len(messages); i++ {
        fmt.Printf("Message %d: %s (Cumulative Cost: %d)\n", i+1, messages[i], costs[i])
    }
}