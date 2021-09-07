package banks

//CheckBalanceRequest struct is input parameter to CheckBalance Method
type CheckBalanceRequest struct {
	AccountNumber int64
	Pin           int
}

//CheckBalanceResponse struct is output parameter from CheckBalance Method
type CheckBalanceResponse struct {
	Balance float64
}

//TransferMoneyRequest struct is input parameter to TransferMoney Method
type TransferMoneyRequest struct {
	AccountNumber int64
	Pin           int
	Amount        float64
}

//TransferMoneyResponse struct is out parameter from TransferMoney Method
type TransferMoneyResponse struct {
	Balance float64
	Status  string
}
