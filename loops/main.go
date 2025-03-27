package main

import "fmt"

func countConnections(groupSize int) int{
	connections := 0
	for i := 0; i < groupSize; i++ {
		connections += i
	}
	return connections
}

func fizzbuzz() {

	for i := 1; i<= 100; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Println("fizzbuzz")
			continue
		}

		if i % 3 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i % 5 == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		fee := float64(i) * 0.01
		totalCost += 1.0 + fee
	}

	return totalCost
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance <= 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func main() {
	fizzbuzz()
	numMessages := 5
	cost := bulkSend(numMessages)
	fmt.Printf("Cost of %d messages is %.2f\n", numMessages, cost)

	// Example 2: Using getMaxMessagesToSend
    costMultiplier := 1.5
    maxCostInPennies := 100
    maxMessages := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
    fmt.Printf("Maximum messages that can be sent with a cost multiplier of %.2f and a budget of %d pennies is %d\n", costMultiplier, maxCostInPennies, maxMessages)
}
