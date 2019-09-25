package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type tagaBase struct {
	db *sql.DB
}

func ConnectDB() {
	tb.db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testgo")

}

var tb tagaBase

func AddArticle(name, tag, comment, url string) {

	query := fmt.Sprintf("INSERT INTO articles (name, tag, comment, url) VALUES (\"%v\", \"%v\", \"%v\", \"%v\")", name, tag, comment, url)
	rows, err := tb.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer rows.Close()

}

func ShowAllNames() string {
	query := fmt.Sprintf("SELECT name FROM articles")
	rows, err := tb.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var names string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = names + name + "\n"
	}
	return names
}

func ShowConcrByName(name string) string { //SELECT * FROM articles WHERE name='test';
	query := fmt.Sprintf("SELECT name, tag, comment, url FROM articles WHERE name=\"%v\"", name)
	rows, err := tb.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var articles []string

	for rows.Next() {
		var name string
		var tag string
		var comment string
		var url string
		if err := rows.Scan(&name, &tag, &comment, &url); err != nil {
			log.Fatal(err)
		}
		article := fmt.Sprintf("Name: %v\nTag: %v\nComment: %v\nURL: %v\n", name, tag, comment, url)

		articles = append(articles, article)
	}
	if len(articles) == 0 {
		return "Article is not found :<"
	}
	return strings.Join(articles, "\n")
}
