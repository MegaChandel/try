package model

type BankStructure struct {
	ID                 int
	Name               string
	Balance            int
	TransactionHistory []string
}
