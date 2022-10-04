package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db_go"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//createEmployee()
	//getEmployee()
	//updateEmployee()
	deleteEmployee()
}

func createEmployee() {
	var employee = Employee{}

	sqlStatement := ` INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning * `

	err = db.QueryRow(sqlStatement, "qwerty s", "qwerty@gmail.com", 21, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func getEmployee() {
	var results = []Employee{}

	sqlStatment := `SELECT * from employees`
	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("employee datas:", results)
}

func updateEmployee() {
	sqlStatement := `UPDATE employees
	SET full_name = $2, email =$3, division = $4, age = $5
	WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1, "qwerty jojo", "qwerty@gmail.com", "qaz", 24)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("update data amount:", count)
}

func deleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("delete data amount:", count)
}
