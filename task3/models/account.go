package models

type Account struct {
	Id      int
	Balance float64
}

func (Account) TableName() string {
	return "account"
}
