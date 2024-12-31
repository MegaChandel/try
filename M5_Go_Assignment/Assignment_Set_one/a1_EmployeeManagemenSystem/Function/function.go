package function

import (
	"A1_EmployeeManagementSystem/model"
	"fmt"
)

func AddEmployee(Database []model.Employee) []model.Employee {
	var newemployee model.Employee

	fmt.Println("Enter the ID")
	fmt.Scan(&newemployee.ID)

	fmt.Println("Enter the Name")
	fmt.Scan(&newemployee.Name)

	fmt.Println("Enter the Age")
	fmt.Scan(&newemployee.Age)

	fmt.Println("Enter the Department")
	fmt.Scan(&newemployee.Department)

	Database = append(Database, newemployee)

	return Database

}

func SearchEmployee(Database []model.Employee, id int) ([]model.Employee, error) {
	if id <= 0 {
		return nil, fmt.Errorf("it cant be zero")
	}
	var ResultEmployee []model.Employee
	for _, emp := range Database {
		if emp.ID == id {
			ResultEmployee = append(ResultEmployee, emp)
		}

	}
	if len(ResultEmployee) == 0 {
		return nil, fmt.Errorf("there is no such employee ")
	}
	return ResultEmployee, nil

}

func FilterByDepartment(Database []model.Employee, department string) ([]model.Employee, error) {
	var Filtered []model.Employee

	for _, emp := range Database {
		if emp.Department == department {
			Filtered = append(Filtered, emp)
		}
	}
	if len(Filtered) == 0 {
		return nil, fmt.Errorf("no such employee ")
	}
	return Filtered, nil

}

func CountEmployee(Database []model.Employee, departmentName string) (int, error) {
	var Count int
	for _, emp := range Database {
		if emp.Department == departmentName {
			Count = Count + 1
		}
	}
	if Count == 0 {
		num := 0
		return num, fmt.Errorf("no such employee of this department")

	}
	return Count, nil

}
