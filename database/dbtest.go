package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func MakeQuery(name, tag, comment, url string) {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testgo")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	fmt.Println("\", ")

	// // perform a db.Query insert
	// INSERT INTO articles (name, tag, comment, url) VALUES (hi, John, I, haven't)
	//query := fmt.Sprintf("INSERT INTO articles (name, tag, comment, url) VALUES ('%v', '%v', '%v', '%v')", name, tag, comment, url)
	query := fmt.Sprintf("INSERT INTO articles (name, tag, comment, url) VALUES (\"%v\", \"%v\", \"%v\", \"%v\")", name, tag, comment, url)
	fmt.Println(query)
	insert, err := db.Query(query)

	// if there is an error inserting, handle it
	if err != nil {
		//fmt.Println(err.)
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
