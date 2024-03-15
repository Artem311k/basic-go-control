package main

import "fmt"

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

func showBalance(balance float64) {
	fmt.Println("Current balance is ", balance)
}

func initChoice() int {
	var choice int

	fmt.Print("Select an option: ")

	fmt.Scan(&choice)

	return choice
}
