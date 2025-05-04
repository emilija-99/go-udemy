package main

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID(id int) (*User, error)
}

type Store interface {
	GetByID(id int) (*User, error)
}

type application struct {
	store Store
}

type PostagresUserRespository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostagresUserRespository {
	return &PostagresUserRespository{db: db}
}

func (r *PostagresUserRespository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name from users WHERE id = $1", id)

	var user User

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	connectionString := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepository := NewPostgresUserRepository(db)
	userService := newUserService(UserRepository)

	app := &application{store: userRepository}

	user, err := userService.getUserByID(1)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("user: %+v\n", user)
	}
}

/*

	Repository Patter it is idea of having a centralized repository with
	the methods and then we send implementations or we inject implementation into
	the runtime.
*/
