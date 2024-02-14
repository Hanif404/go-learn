package db

type Department struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Employee struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	DepartmentID int64  `db:"department_id"`
}
