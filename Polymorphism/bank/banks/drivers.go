package banks

import "errors"

type Driver interface {
	GetCheckBalanceDriver(bankCode string) (CheckBalanceDriver, error)
	GetTransferMoneyDriver(bankCode string) (TransferMoneyDriver, error)
}

func NewDriver(amex *AmericanExpress, citi *CitiBank) Driver {
	return &driver{
		americanExpress: amex,
		citiBank:        citi,
	}
}

type driver struct {
	americanExpress *AmericanExpress
	citiBank        *CitiBank
}

func (d *driver) GetCheckBalanceDriver(bankCode string) (CheckBalanceDriver, error) {
	switch bankCode {
	case "AMERICAN-EXPRESS":
		return d.americanExpress, nil
	case "CITI-BANK":
		return d.citiBank, nil
	default:
		return nil, errors.New("bank not found")
	}
}

func (d *driver) GetTransferMoneyDriver(bankCode string) (TransferMoneyDriver, error) {
	switch bankCode {
	case "AMERICAN-EXPRESS":
		return d.americanExpress, nil
	case "CITI-BANK":
		return d.citiBank, nil
	default:
		return nil, errors.New("bank not found")
	}
}
