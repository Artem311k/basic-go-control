package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

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
