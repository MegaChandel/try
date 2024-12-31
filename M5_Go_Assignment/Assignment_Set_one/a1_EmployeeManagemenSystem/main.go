package main

import (
	function "A1_EmployeeManagementSystem/Function"
	"A1_EmployeeManagementSystem/model"
	"fmt"
)

func main() {
	Database := []model.Employee{
		{ID: 1, Name: "Megha Chandel", Age: 21, Department: "IT"},
		{ID: 2, Name: "Shiva Chandel", Age: 22, Department: "IT"},
		{ID: 3, Name: "Kamlesh Chandel", Age: 45, Department: "HR"},
		{ID: 4, Name: "Ranvijay Singh", Age: 30, Department: "Trainer"},
	}

	var Action int

	fmt.Println("Enter the operation you want to Perform")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Search Employee")
	fmt.Println("1. List Employee By department")
	fmt.Println("2. Count Employee")
	fmt.Scan(&Action)

	switch Action {
	case 1:
		Database = function.AddEmployee(Database)
		fmt.Println(Database)

	case 2:
		fmt.Println("Enter the id of the employee you want to search ")
		var id int
		fmt.Scan(&id)
		result, err := function.SearchEmployee(Database, id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	case 3:
		fmt.Println("Enter the name of the department you would want to filter")
		var dep string
		fmt.Scan(&dep)

		results, err := function.FilterByDepartment(Database, dep)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(results)

	case 4:
		fmt.Println("Enter the name of the department whose count you want ")
		var dep string

		fmt.Scan(&dep)

		results, err := function.CountEmployee(Database, dep)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(results)

	}

}
