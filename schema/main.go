package main

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/jcunhasilva/golang-todo-list/config"
	"github.com/jcunhasilva/golang-todo-list/task"
)

func createSchemaFromModel(db *pg.DB) error {
	err := db.CreateTable((*task.Task)(nil), &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Cannot load configurations")
		os.Exit(-1)
	}

	db := pg.Connect(&pg.Options{
		Addr:     config.Host,
		User:     config.User,
		Password: config.Password,
		Database: config.Name,
	})

	if err := createSchemaFromModel(db); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Schema created.")
}
