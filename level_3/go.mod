module level3/main

go 1.21.6

replace level3/dbClient => ./db_client

require (
	github.com/golang-one/batching v1.0.1
	github.com/stretchr/testify v1.8.3
	level3/dbClient v0.0.0-00010101000000-000000000000
	level3/department v0.0.0-00010101000000-000000000000
	level3/employee v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace level3/department => ./department

replace level3/employee => ./employee
