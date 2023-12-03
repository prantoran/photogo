package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	Name   string
	Email  string `gorm:"not null;uniqueIndex"` // creates unique index `idx_users_email`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.Logger.LogMode(logger.Info)
	// db.Migrator().DropTable(&User{}) // drop table if exists

	// Migrate the schema, creates the table `users`
	db.AutoMigrate(&User{}, &Order{})

	var u User
	if err := db.Preload("Orders").First(&u).Error; err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(u)

	// createOrder(db, u, 1001, "fake desc #1")
	// createOrder(db, u, 1002, "fake desc #2")
	// createOrder(db, u, 1003, "fake desc #3")

	// errorHandling(db)
	// selectFirstLast(db)
	// checkGormModelEmbedding()
	// insertUser(db)
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error

	if err != nil {
		panic(err)
	}
}

func errorHandling(db *gorm.DB) {
	var u User
	if err := db.Where("email = ?", "blah@blah").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found!")
		default:
			panic(err)
		}
	}
}

func selectFirstLast(db *gorm.DB) {
	var u, v, w User
	db.First(&u)
	fmt.Println(u)
	db.Last(&v)
	fmt.Println(v)
	db.Where("id = ?", v.ID).First(&w)
	fmt.Println(w)

	var users []User
	db.Find(&users)
	fmt.Println(len(users))
	fmt.Println(users)
}

func checkGormModelEmbedding() {
	user := User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
		},
	}

	fmt.Println("user.ID:", user.ID, "user.CreatedAt:", user.CreatedAt, "user.Model.CreatedAt:", user.Model.CreatedAt)
}

func insertUser(db *gorm.DB) {
	name, email := getInfo()
	u := User{
		Name:  name,
		Email: email,
	}
	if err := db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v", u)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email address?")
	email, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	return name, email
}
