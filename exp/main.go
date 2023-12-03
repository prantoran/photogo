package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "photogo_db"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"` // creates unique index `idx_users_email`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&User{}) // drop table if exists

	// Migrate the schema, creates the table `users`
	db.AutoMigrate(&User{})

	user := User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
		},
	}

	fmt.Println("user.ID:", user.ID, "user.CreatedAt:", user.CreatedAt, "user.Model.CreatedAt:", user.Model.CreatedAt)
}
