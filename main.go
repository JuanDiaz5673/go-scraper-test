package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:(password)@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var name string = "ELLIOT"

	delete, err := db.Query("DELETE FROM users WHERE name='%s'", name)
	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()


	fmt.Println("Successfully inserted into user tables")
}