package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is hte startup err: %s\\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up the db connection
// and routers
func run() error {
	connectionString := ""

	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	// create storage repository
	storage := repository.NewStorage(db)

	// create router dependency
	router := gin.Default()
	router.use(cors.Default())

	// create user service
	userService := api.NewUserService(storage)

	// create weight service
	weightService := api.NewWeightService(storage)

	server := app.NewServer(router, userService, weightService)

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
