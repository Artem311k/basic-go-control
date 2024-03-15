package main

import (
	"fmt"
	"strconv"
)

func processBalanceOperation(balance float64, choice int) float64 {
	money := initAmount()

	switch choice {
	case 2:
		fmt.Println("Depositing ", money, " to balance...")
		balance += money
	case 3:
		fmt.Println("Withdrowing ", money, " from balance...")
		balance -= money
		if balance < 0 {
			balance = 0
			fmt.Print("Not enough money... Withdrowing ", balance)
		}
	default:
		fmt.Println("Invalid opration choice ", choice)
		return balance
	}
	fmt.Println("Current balance is ", balance)
	return balance
}

func initAmount() float64 {
	var amount string
	var parsedAmount float64
	fmt.Print("Enter amount to proceed: ")
	fmt.Scan(&amount)
	var err error
	parsedAmount, err = strconv.ParseFloat(amount, 64)
	for parsedAmount < 0 || err != nil {
		fmt.Println("Invalid amount! Try again!")
		fmt.Print("Enter amount to proceed: ")
		fmt.Scan(&amount)
		parsedAmount, err = strconv.ParseFloat(amount, 64)
	}
	return parsedAmount
}
