package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "users"
)

func getUserByID(db *sql.DB, id int) (*User, error) {
	row := db.QueryRow("SELECT id, name from users WHERE id $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	connString := "user=usernmae dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	user, err := getUserByID(db, 1)

	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}
}
