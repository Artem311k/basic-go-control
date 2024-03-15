package main

import (
	"fmt"
)

func main() {

	exec("account.json")

}

func exec(fileName string) {
	accountBalance := initBalance(fileName)
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
			exit(fileName, accountBalance)
			return
		} else {
			fmt.Println("Not valid choice! Try again!")
		}
	}
}

func exit(fileName string, balance float64) {
	writeBalanceTofile(fileName, balance)
	fmt.Println("Goodbye!")
}
