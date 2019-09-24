package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func MakeQuery(name, tag, comment, url string) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testgo")
	if err != nil {
		panic(err.Error())
	}

	query := fmt.Sprintf("INSERT INTO articles (name, tag, comment, url) VALUES (\"%v\", \"%v\", \"%v\", \"%v\")", name, tag, comment, url)
	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
