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

var tb tagaBase

//root:@tcp(127.0.0.1:3306)/testgo

func ConnectDB() {
	var err error
	tb.db, err = sql.Open("mysql", "iF8deKoJkQ:Hgfdk5m3Op@tcp(remotemysql.com:3306)/iF8deKoJkQ")
	if err != nil {
		panic(err.Error())
	}
}

func AddArticle(user, name, tag, comment, url string) {

	query := fmt.Sprintf("INSERT INTO articles (user, name, tag, comment, url) VALUES (\"%v\", \"%v\", \"%v\", \"%v\", \"%v\")", user, name, tag, comment, url)
	rows, err := tb.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
}

func ShowAllNames() string {
	query := fmt.Sprintf("SELECT name FROM articles")
	rows := getRows(tb.db, query)
	defer rows.Close()

	names := genNamesList(rows)
	return names
}

func genNamesList(rows *sql.Rows) string {
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

func ShowConcrByName(user, name string) string { //SELECT * FROM articles WHERE name='test';

	query := fmt.Sprintf("SELECT name, tag, comment, url FROM articles WHERE user=\"%v\" and name=\"%v\"", user, name)
	rows := getRows(tb.db, query)
	defer rows.Close()

	article := genArtclInfo(rows)
	return article
}

func genArtclInfo(rows *sql.Rows) string {
	var articles []string

	for rows.Next() {
		var name, tag, comment, url string
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

func getRows(db *sql.DB, query string) *sql.Rows {
	rows, err := tb.db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
