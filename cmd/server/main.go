package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"jamesbest.tech/weight-tracker/pkg/api"
	"jamesbest.tech/weight-tracker/pkg/app"
	"jamesbest.tech/weight-tracker/pkg/repository"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// main is dump and calls run so its easier to test
// https://pace.dev/blog/2020/02/12/why-you-shouldnt-
// use-func-main-in-golang-by-mat-ryer.html
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup err: %s\\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up the db connection
// and routers
func run() error {
	dbpath := goDotEnvVariable("DB_PATH")

	connectionString := dbpath

	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage repository
	storage := repository.NewStorage(db)

	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	// create user service
	userService := api.NewUserService(storage)

	// create weight service
	// weightService := api.NewWeightService(storage)

	server := app.NewServer(router, userService)

	// start server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping database to ensure it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
