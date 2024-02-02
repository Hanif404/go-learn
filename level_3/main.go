package main

import (
	"context"
	"fmt"
	"github.com/golang-one/batching"
	"level3/dbClient"
	"level3/department"
	"level3/employee"
	"math/rand"
	"time"
)

func main() {
	db_client.InitialDBConn()
	rand.Seed(time.Now().UnixNano())

	var (
		ctx   = context.Background()
		start = time.Now()
		tasks = []rune("ABCDE")
	)
	batching.Do(ctx, tasks, func(ctx context.Context, batch []rune) error {
		b := make([]rune, 1)
		for i := range b {
			b[i] = tasks[rand.Intn(len(tasks))]
		}
		runTask3a(string(b))
		time.Sleep(time.Second)
		return nil
	}, batching.BatchSize(1), batching.MaxThreads(1), batching.MinSizeForConcurrency(0))
	fmt.Println(time.Since(start))
	fmt.Println("=====================")

	batching.Do(ctx, tasks, func(ctx context.Context, batch []rune) error {
		b := make([]rune, 1)
		for i := range b {
			b[i] = tasks[rand.Intn(len(tasks))]
		}
		runTask3b(string(b))
		time.Sleep(time.Second)
		return nil
	}, batching.BatchSize(1), batching.MaxThreads(1), batching.MinSizeForConcurrency(0))
	fmt.Println(time.Since(start))
}

func runTask3a(name string) {
	department1Id, _ := department.PostDepartment("Company " + name)
	employee.PostEmployee("Yusuf", department1Id)
	employee.PostEmployee("Hasan", department1Id)

	emp1, _ := employee.GetEmployeeByDepartId(department1Id)
	fmt.Println(emp1)
}

func runTask3b(name string) {
	department2Id, _ := department.PostDepartment("Company " + name)
	empAll, _ := employee.GetEmployee()
	lastEmp := empAll[len(empAll)-1]

	employee.PutEmployee(lastEmp.ID, lastEmp.Name, department2Id)

	emp2, _ := employee.GetEmployeeByDepartId(department2Id)
	fmt.Println(emp2)
}
