package main

import (
	"fmt"
	"os"

	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host") + ":" + os.Getenv("db_port")

	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: dbName,
		Addr:     dbHost,
	})

	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello")
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*domain.Podcast)(nil), (*domain.Profile)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
