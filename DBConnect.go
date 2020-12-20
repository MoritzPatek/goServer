package main

import (
	"database/sql"
	"fmt"
)

const (
	hostname     = "doYouReallyThink"
	hostport     = 1111
	username     = "iWould"
	password     = "LetStuffLikeThat"
	databasename = "inTheRepoLmao"
)

//DB exists so every controller can access it
var DB *sql.DB

//ConnectDB opens Connection between DB and server
func ConnectDB() {
	pgconstring := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		hostname, hostport, username, password, databasename)
	db, err := sql.Open("postgres", pgconstring)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("You are Successfully connected!")
}
