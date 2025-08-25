package models

type Transaction struct {
	Id            int
	FromBalanceId float64
	ToBalanceId   float64
	Amount        float64
}

func (Transaction) TableName() string {
	return "transaction"
}
