package employee

import (
	"fmt"
	"level3/dbClient"
)

type Employee struct {
	ID         int64
	Name       string
	DepartId   int64
	DepartName string
}

func GetEmployeeById(id int64) (Employee, error) {
	var emp Employee

	row := db_client.DBClient.QueryRow("SELECT emp.*, dpt.department_name FROM Employee emp LEFT JOIN Department dpt on emp.department_id = dpt.department_id WHERE employee_id = ?", id)
	if err := row.Scan(&emp.ID, &emp.Name, &emp.DepartId, &emp.DepartName); err != nil {
		return emp, fmt.Errorf("failed : %v", err)
	}
	return emp, nil
}

func GetEmployeeByDepartId(departId int64) ([]Employee, error) {
	var employees []Employee

	rows, err := db_client.DBClient.Query("SELECT emp.*, dpt.department_name FROM Employee emp LEFT JOIN Department dpt on emp.department_id = dpt.department_id where emp.department_id =?", departId)
	if err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}

	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.DepartId, &emp.DepartName); err != nil {
			return nil, fmt.Errorf("failed : %v", err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}
	return employees, nil
}

func GetEmployee() ([]Employee, error) {
	var employees []Employee

	rows, err := db_client.DBClient.Query("SELECT emp.*, dpt.department_name FROM Employee emp LEFT JOIN Department dpt on emp.department_id = dpt.department_id")
	if err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}

	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.DepartId, &emp.DepartName); err != nil {
			return nil, fmt.Errorf("failed : %v", err)
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}
	return employees, nil
}

func PostEmployee(name string, departId int64) (int64, error) {
	res, err := db_client.DBClient.Exec("INSERT INTO Employee(employee_name, department_id) VALUES (?, ?);", name, departId)
	if err != nil {
		return 0, fmt.Errorf("failed : %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed : %v", err)
	}
	return id, nil
}

func PutEmployee(id int64, name string, departId int64) (bool, error) {
	_, err := db_client.DBClient.Exec("UPDATE Employee set employee_name=?, department_id=? where employee_id=?;", name, departId, id)
	if err != nil {
		return false, fmt.Errorf("failed : %v", err)
	}
	return true, nil
}

func DelEmployee(id int64) (bool, error) {
	_, err := db_client.DBClient.Exec("DELETE FROM Employee where employee_id=?;", id)
	if err != nil {
		return false, fmt.Errorf("failed : %v", err)
	}
	return true, nil
}
