package db

func (queries *Queries) PostDepartment(name string) (int64, error) {
	query := queries.db.Rebind(`INSERT INTO department (name) VALUES (?)`)
	res, err := queries.db.Exec(query, name)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (q *Queries) GetDepartmentByID(id int64) (Department, error) {
	var i Department
	query := q.db.Rebind(`SELECT * FROM department WHERE id = ?`)
	err := q.db.Get(&i, query, id)
	return i, err
}

func (queries *Queries) PutDepartment(id int64, name string) error {
	query := queries.db.Rebind(`UPDATE department SET name = ? WHERE id = ?`)
	_, err := queries.db.Exec(query, name, id)
	return err
}

func (queries *Queries) DelDepartment(id int64) error {
	query := queries.db.Rebind(`DELETE FROM department WHERE id = ?`)
	_, err := queries.db.Exec(query, id)
	return err
}
