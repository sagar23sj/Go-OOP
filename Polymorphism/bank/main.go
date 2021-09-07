package main

import (
	"fmt"

	"github.com/sagar23sj/Go-OOP/Polymorphism/bank/banks"
)

func main() {
	//Note:  All this initialization can be done while setting up app dependencies
	//Setting up driver variable for fetching respective bank drivers at runtime

	//initializing initAMEX with AmericanExpress type
	initAMEX := banks.NewAmericanExpressDriver()

	//initializing initCITI with CitiBank type
	initCITI := banks.NewCitiBankDriver()

	//initializing driver object with Driver interface
	driver := banks.NewDriver(initAMEX, initCITI)

	fmt.Println("\n----------------Transactions on Citi Bank----------------")
	selectedBank_1 := "CITI-BANK"

	//loading correct CheckbalanceDriver at runtime based on Bank Code
	checkBalanceDriver, _ := driver.GetCheckBalanceDriver(selectedBank_1)

	//Request Body for CheckBalance method
	checkBalanceRequest := banks.CheckBalanceRequest{
		AccountNumber: 1000100101,
		Pin:           1234,
	}

	//Calling CheckBalance method on fetched bank driver
	checkBalanceResponse := checkBalanceDriver.CheckBalance(checkBalanceRequest)

	fmt.Println("\nBalance : ", checkBalanceResponse.Balance)

	//loading correct TransferMoneyDriver at runtime based on Bank Code
	transferMoneyDriver, _ := driver.GetTransferMoneyDriver(selectedBank_1)

	//Request Body for TransferMoney method
	transferMoneyRequest := banks.TransferMoneyRequest{
		AccountNumber: 1000100101,
		Pin:           1234,
		Amount:        111111,
	}

	//Calling TransferMoney method on fetched bank driver
	transferMoneyResponse := transferMoneyDriver.TransferMoney(transferMoneyRequest)

	fmt.Printf("\nTransfer Status : %s ------- Balance : %f", transferMoneyResponse.Status, transferMoneyResponse.Balance)

	fmt.Println("\n\n----------------Transactions on American Express Bank----------------")

	selectedBank_2 := "AMERICAN-EXPRESS"

	checkBalanceDriver, _ = driver.GetCheckBalanceDriver(selectedBank_2)

	checkBalanceRequest = banks.CheckBalanceRequest{
		AccountNumber: 200020020,
		Pin:           5678,
	}

	checkBalanceResponse = checkBalanceDriver.CheckBalance(checkBalanceRequest)

	fmt.Println("\nBalance : ", checkBalanceResponse.Balance)

	transferMoneyDriver, _ = driver.GetTransferMoneyDriver(selectedBank_2)

	transferMoneyRequest = banks.TransferMoneyRequest{
		AccountNumber: 200020020,
		Pin:           5678,
		Amount:        222222,
	}

	transferMoneyResponse = transferMoneyDriver.TransferMoney(transferMoneyRequest)

	fmt.Printf("\nTransfer Status : %s ------- Balance : %f", transferMoneyResponse.Status, transferMoneyResponse.Balance)

}
