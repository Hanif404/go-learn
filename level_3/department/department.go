package department

import (
	"fmt"
	"level3/dbClient"
)

type Department struct {
	ID   int64
	Name string
}

func GetDepartment() ([]Department, error) {
	var departments []Department

	rows, err := db_client.DBClient.Query("SELECT * FROM Department")
	if err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}

	for rows.Next() {
		var dpt Department
		if err := rows.Scan(&dpt.ID, &dpt.Name); err != nil {
			return nil, fmt.Errorf("failed : %v", err)
		}
		departments = append(departments, dpt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed : %v", err)
	}
	return departments, nil
}

func PostDepartment(name string) (int64, error) {
	res, err := db_client.DBClient.Exec("INSERT INTO Department(department_name) VALUES (?);", name)
	if err != nil {
		return 0, fmt.Errorf("failed : %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed : %v", err)
	}
	return id, nil
}

func PutDepartment(id int, name string) (bool, error) {
	_, err := db_client.DBClient.Exec("UPDATE Department set department_name=? where department_id=?;", name, id)
	if err != nil {
		return false, fmt.Errorf("failed : %v", err)
	}
	return true, nil
}

func DelDepartmentById(id int) (bool, error) {
	_, err := db_client.DBClient.Exec("DELETE FROM Department where department_id=?;", id)
	if err != nil {
		return false, fmt.Errorf("failed : %v", err)
	}
	return true, nil
}

func DelDepartment() (bool, error) {
	_, err := db_client.DBClient.Exec("DELETE FROM Department;")
	if err != nil {
		return false, fmt.Errorf("failed : %v", err)
	}
	return true, nil
}
