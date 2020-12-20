package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	username string `json:"username"`
}

//GetUserNames returns all names from the db
func GetUserNames(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT username FROM person;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var courses []*person
	for rows.Next() {
		c := new(person)
		err := rows.Scan(&c.username)
		if err != nil {
			fmt.Println(err)
			return
		}

		courses = append(courses, c)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(courses); err != nil {
		fmt.Println(err)
	}

}
