package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	exec(initBalance())

}

func exec(accountBalance float64) {

	var isContinue bool = true

	printWelcome()

	for isContinue {

		printChicesList()

		choice := initChoice()

		printUserChoice(choice)

		if choice == 2 || choice == 3 {
			accountBalance = processBalanceOperation(accountBalance, choice)
		} else if choice == 1 {
			showBalance(accountBalance)
		} else if choice == 4 {
			isContinue = false
			exit()
		} else {
			fmt.Println("Not valid choice! Try again!")
		}
	}
}

func initBalance() float64 {
	var startbalance string
	fmt.Print("Init start balance: ")
	fmt.Scan(&startbalance)
	if startbalance == "Exit" {
		exit()
	}
	f, err := strconv.ParseFloat(startbalance, 64)
	if err != nil {
		fmt.Println("Not valid balance. Try again or type Exit!")
		initBalance()
	}
	return f
}

func exit() {
	fmt.Println("Goodbye!")
	os.Exit(3)
}

func showBalance(balance float64) {
	fmt.Println("Current balance is ", balance)
}

func processBalanceOperation(balance float64, choice int) float64 {
	fmt.Print("Enter money value to proceed: ")
	var money float64
	fmt.Scan(&money)

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

func initChoice() int {
	var choice int

	fmt.Print("Select an option: ")

	fmt.Scan(&choice)

	return choice
}

func printWelcome() {
	fmt.Println("Welcom to go Bank!")
}

func printChicesList() {
	fmt.Println("What you want to do?")
	fmt.Println("1. Check balacne")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdrow money")
	fmt.Println("4. Exit")
}

func printChoice(choiceName string) {
	fmt.Println("Your choice is ", choiceName)
}

func printUserChoice(choice int) {
	switch choice {
	case 1:
		printChoice("Check balance")
	case 2:
		printChoice("Deposit money")
	case 3:
		printChoice("Withdrow money")
	case 4:
		printChoice("Exit")
	default:
		printChoice("something is wrong...")
	}
}
