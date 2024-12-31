package main

import (
	"a2_BankTransactionSystem/handlers"
	"a2_BankTransactionSystem/model"
	"fmt"
)

func main() {

	Database := []model.BankStructure{
		{ID: 1, Name: "Megha Chandel", Balance: 100000, TransactionHistory: []string{}},
		{ID: 2, Name: "Shiva Chandel", Balance: 200000, TransactionHistory: []string{}},
		{ID: 3, Name: "Kamlesh Chandel", Balance: 300000, TransactionHistory: []string{}},
	}

	var option int

	fmt.Println("Enter the operation you want to Perform\n 1.Deposit\n2.Withdraw \n3.TransactinHistory ")
	fmt.Scan(&option)

	switch option {
	case 1:
		var Amnt int
		var id int
		fmt.Println("Enter the id")
		fmt.Scan(&id)
		fmt.Println("Enter the amount you want to deposit")
		fmt.Scan(&Amnt)

		UpdatedDatabase, err := handlers.DepositFunction(Database, Amnt, id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(UpdatedDatabase)

	case 2:
		var Amnt int
		var id int
		fmt.Println("Enter the id")
		fmt.Scan(&id)
		fmt.Println("Enter the amount you want to withdraw")
		fmt.Scan(&Amnt)

		UpdatedDatabase, err := handlers.WithdrawFunction(Database, Amnt, id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(UpdatedDatabase)

	case 3:
		var id int
		fmt.Println("Enter the id")
		fmt.Scan(&id)

		handlers.TransHistory(Database, id)

	}

}
