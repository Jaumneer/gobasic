package main

import (
	"fmt"
	"os"
)

func writeToFile(profit float64, expenses float64, revenue float64, taxRate float64) {
	result := fmt.Sprintf("Profit: %.2f\nExpenses: %.2f\nRevenue: %.2f\nTax Rate: %.2f%%\n", profit, expenses, revenue, taxRate)
	os.WriteFile("database.txt", []byte(result), 0644)
}

func main() {
	revenue := getUserInput("Revenue :")
	expenses := getUserInput("Expenses :")
	taxRate := getUserInput("Tax Rate :")

	profit := (revenue - expenses) * (1 - taxRate/100)

	fmt.Printf("Profit :%.2f\n", profit)
	writeToFile(profit, expenses, revenue, taxRate)

}

func getUserInput(text string) float64 {
	var userInput float64 = 0.0
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("PLEASE INPUT POZITIVE NUMBERS")
		fmt.Println("--------------------------------")
		os.Exit(21)
	}

	return userInput

}
