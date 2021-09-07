package banks

import "fmt"

type CitiBank struct{}

func NewCitiBankDriver() *CitiBank {
	return &CitiBank{}
}

func (c *CitiBank) CheckBalance(req CheckBalanceRequest) CheckBalanceResponse {
	fmt.Printf("\n--------Check Balance Service--------\nBank: CITI BANK, Account Number : %d, Pin : %d", req.AccountNumber, req.Pin)
	return CheckBalanceResponse{
		Balance: 2222222,
	}
}

func (c *CitiBank) TransferMoney(req TransferMoneyRequest) TransferMoneyResponse {
	fmt.Printf("\n--------Transfer Money Service--------\nBank: CITI BANK, Account Number : %d, Pin : %d, Amount : %f", req.AccountNumber, req.Pin, req.Amount)
	return TransferMoneyResponse{
		Status:  "success",
		Balance: 2222222 - req.Amount,
	}
}
