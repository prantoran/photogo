package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "photogo_db"
)

func insertQueryRow(db *sql.DB) {

	var id int
	err := db.QueryRow(`
		INSERT INTO users(name, email)
		VALUES($1, $2)
		RETURNING id`, "Pinku", "prantoran@gmail.com").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("id: ", id)
}

func selectQueryRow(db *sql.DB) {
	var id int
	var name, email string
	err := db.QueryRow(`
		SELECT id, name, email
		FROM users
		WHERE id=$1`, 1).Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			panic(err)
		}
	}

	fmt.Println("id: ", id, "name:", name, "email:", email)
}

type User struct {
	ID    int
	Name  string
	Email string
}

func selectQuery(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, name, email
		FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		fmt.Println("id: ", user.ID, "name:", user.Name, "email:", user.Email)
	}

	if rows.Err() != nil {
		// handle the err!
	}
	fmt.Println(users)
}

func insertExec(db *sql.DB) {
	for i := 1; i <= 6; i++ {
		userID := 1
		if i > 3 {
			userID = 3
		}
		amount := i * 100
		description := fmt.Sprintf("Day: %d", i)

		_, err := db.Exec(`
			INSERT INTO orders(user_id, amount, description)
			VALUES ($1, $2, $3)`, userID, amount, description)

		if err != nil {
			panic(err)
		}
	}
}

func joinExec(db *sql.DB) {

	rows, err := db.Query(`
		SELECT users.id, users.email, users.name, orders.id, orders.amount, orders.description FROM users
		INNER JOIN orders ON users.id=orders.user_id`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var userID, orderID, amount int
		var email, name, desc string
		if err := rows.Scan(&userID, &email, &name, &orderID, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("userID:", userID, "email:", email, "name:", name, "orderID:", orderID, "amount:", amount, "desc:", desc)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// insertQueryRow(db)
	// selectQueryRow(db)
	// selectQuery(db)
	// insertExec(db)
	joinExec(db)
}
