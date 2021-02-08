package db

import (
	"io/ioutil"
	"sync"
	"database/sql"

	// Need this import for its side effects with sqlx.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

const (
	dsnCreate = "app:password@tcp(gym-rat_db_1:3306)/"
	dsnConnect = "app:password@tcp(gym-rat_db_1:3306)/db?parseTime=true&clientFoundRows=true"
	driver = "mysql"
)

var (
	dbClient     *sqlx.DB
	dbClientSync sync.Once
)

// Client returns an initialised DB client.
// It will initialise the client and create database and the tables in the
// db/tables directory on first call.
func Client() *sqlx.DB {
	dbClientSync.Do(func() {
		createClient, createErr := sql.Open(driver, dsnCreate)
		_,createErr = createClient.Exec("CREATE DATABASE IF NOT EXISTS db")
		if createErr != nil {
				panic(createErr)
		}
		createClient.Close()

		client, err := sqlx.Connect(driver, dsnConnect)
		if err != nil {
			panic(err)
		}

		client.Mapper = reflectx.NewMapper("json")

		err = initTables(client)
		if err != nil {
			panic(err)
		}

		dbClient = client
	})

	return dbClient
}

var TablesDir = "./db/tables/"

func initTables(client *sqlx.DB) error {
	files := [5]string{"workout.sql", "exercise.sql", "workout_exercise.sql", "day.sql", "exerciseset.sql"}

	for _, file := range files {
		contents, err := ioutil.ReadFile(TablesDir + file)
		if err != nil {
			return err
		}

		_, err = client.Exec(string(contents))
		if err != nil {
			return err
		}
	}

	return nil
}
