package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Manual@1995@tcp(mysql:3306)/pixie")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Persons (PersonID, LastName, FirstName, Address, City) VALUES (1, 'Tom B. Erichsen', 'Skagen 21', 'Stavanger', '4006');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yay, values added!")
}
