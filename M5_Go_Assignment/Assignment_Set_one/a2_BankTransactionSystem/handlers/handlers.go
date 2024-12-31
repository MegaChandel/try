package handlers

import (
	"a2_BankTransactionSystem/model"
	"fmt"
)

func DepositFunction(Database []model.BankStructure, amount int, id int) (model.BankStructure, error) {
	if amount <= 0 {
		return model.BankStructure{}, fmt.Errorf("the amount to be deposited cannot be 0")

	}
	if id <= 0 {
		return model.BankStructure{}, fmt.Errorf("the id cannot be 0 or less")

	}

	for i, acc := range Database {
		if acc.ID == id {
			Database[i].Balance = Database[i].Balance + amount
			transaction := fmt.Sprintf("Amount of %d deposited", amount)
			Database[i].TransactionHistory = append(Database[i].TransactionHistory, transaction)
			return Database[i], nil

		}

	}
	return model.BankStructure{}, nil

}

func WithdrawFunction(Database []model.BankStructure, amount int, id int) (model.BankStructure, error) {
	if amount <= 0 {
		return model.BankStructure{}, fmt.Errorf("amount cannot be Zero")
	}

	var index int
	for i, user := range Database {
		if user.ID == id {
			index = i
			transaction := fmt.Sprintf("Amount of %d withdrawn", amount)
			Database[i].TransactionHistory = append(Database[i].TransactionHistory, transaction)
			break
		}
	}
	Database[index].Balance = Database[index].Balance - amount
	return Database[index], nil

}

func TransHistory(Database []model.BankStructure, id int) error {
	if id <= 0 {
		return fmt.Errorf("id cannot be 0")
	}
	for _, result := range Database {
		if result.ID == id {
			if len(result.TransactionHistory) == 0 {
				return fmt.Errorf("no transaction history")
			}
			fmt.Println(result.TransactionHistory)
			return nil

		}
	}
	return fmt.Errorf("account with %d not found", id)
}
