package modules

import (
	"fmt"
	"sb3/level4/db"
)

func GetDepartmentByID(id int64) (db.Department, error) {
	conn, err := db.Initialize()
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal konek database : %v", err)
	}
	q := db.New(conn)
	data, err := q.GetDepartmentByID(id)
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal read : %v", err)
	}
	return data, nil
}

func PostDepartment(name string) (db.Department, error) {
	conn, err := db.Initialize()
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal konek database : %v", err)
	}

	q := db.New(conn)
	id, err := q.PostDepartment(name)
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal insert : %v", err)
	}

	data, err := GetDepartmentByID(id)
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal ambil data : %v", err)
	}
	return data, nil
}

func PutDepartment(id int64, name string) (db.Department, error) {
	conn, err := db.Initialize()
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal konek database : %v", err)
	}

	q := db.New(conn)
	err = q.PutDepartment(id, name)
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal update : %v", err)
	}

	data, err := GetDepartmentByID(id)
	if err != nil {
		return db.Department{}, fmt.Errorf("gagal ambil data : %v", err)
	}
	return data, nil
}

func DelDepartment(id int64) error {
	conn, err := db.Initialize()
	if err != nil {
		return fmt.Errorf("gagal konek database : %v", err)
	}

	q := db.New(conn)
	err = q.DelDepartment(id)
	if err != nil {
		return fmt.Errorf("gagal delete : %v", err)
	}
	return nil
}
