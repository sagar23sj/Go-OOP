package banks

type CheckBalanceDriver interface {
	CheckBalance(CheckBalanceRequest) CheckBalanceResponse
}

type TransferMoneyDriver interface {
	TransferMoney(TransferMoneyRequest) TransferMoneyResponse
}
