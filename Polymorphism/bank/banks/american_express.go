package banks

import "fmt"

type AmericanExpress struct{}

func NewAmericanExpressDriver() *AmericanExpress {
	return &AmericanExpress{}
}

func (a *AmericanExpress) CheckBalance(req CheckBalanceRequest) CheckBalanceResponse {
	fmt.Printf("\n--------Check Balance Service--------\nBank: AMERICAN EXPRESS BANK, Account Number : %d, Pin : %d", req.AccountNumber, req.Pin)
	return CheckBalanceResponse{
		Balance: 5555555,
	}
}

func (a *AmericanExpress) TransferMoney(req TransferMoneyRequest) TransferMoneyResponse {
	fmt.Printf("\n--------Transfer Money Service--------\nBank: AMERICAN EXPRESS, Account Number : %d, Pin : %d, Amount : %f", req.AccountNumber, req.Pin, req.Amount)
	return TransferMoneyResponse{
		Status:  "success",
		Balance: 5555555 - req.Amount,
	}
}
