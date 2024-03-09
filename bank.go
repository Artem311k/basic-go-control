package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func writeBalanceTofile(fileName string, balance float64) {
	if fileName != "" {
		var accountInfo AccountInfo
		accountInfo.AccountAmount = balance
		content, _ := json.Marshal(accountInfo)
		err := os.WriteFile("account.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func initBalance(fileNameFile string) float64 {
	if fileNameFile != "" {
		return readBalanceFromFile(fileNameFile)
	}
	var startbalance string
	fmt.Print("Init start balance: ")
	fmt.Scan(&startbalance)
	if startbalance == "Exit" {
		exit(fileNameFile, 0)
	}
	f, err := strconv.ParseFloat(startbalance, 64)
	if err != nil {
		fmt.Println("Not valid balance. Try again or type Exit!")
		initBalance(fileNameFile)
	}
	return f
}

func exit(fileName string, balance float64) {
	writeBalanceTofile(fileName, balance)
	fmt.Println("Goodbye!")
}

func showBalance(balance float64) {
	fmt.Println("Current balance is ", balance)
}

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

func readBalanceFromFile(fileNme string) float64 {
	if _, err := os.Stat("account.json"); err != nil {
		writeBalanceTofile(fileNme, 0)
	}

	file, err := os.Open(fileNme)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened ", file.Name())

	byteValue, _ := io.ReadAll(file)

	var accountInfo AccountInfo

	json.Unmarshal(byteValue, &accountInfo)

	defer file.Close()

	return accountInfo.AccountAmount
}

type AccountInfo struct {
	AccountAmount float64
}
