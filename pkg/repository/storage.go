package repository

import (
	"database/sql"
	"errors"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"jamesbest.tech/weight-tracker/pkg/api"

	_ "github.com/lib/pq"
)

type Storage interface {
	RunMigrations(connectionString string) error
	CreateUser(request api.NewUserRequest) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}

	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationPath := filepath.Join("file://", basePath, "pkg/repository/migrations")

	m, err := migrate.New(migrationPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil

	}

	return nil
}

func (s *storage) CreateUser(request api.NewUserRequest) error {
	newUserStatement := `
		INSERT INTO "user" (name, age, height, sex, activity_level, email, weight_goal)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	err := s.db.QueryRow(newUserStatement, request.Name, request.Age, request.Height, request.Sex, request.ActivityLevel, request.Email, request.WeightGoal).Err()

	if err != nil {
		log.Printf("this was the error %v", err.Error())
	}

	return nil
}
