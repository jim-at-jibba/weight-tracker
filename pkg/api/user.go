package api

import (
	"errors"
	"log"
	"strings"
)

// UserServer contains the methods of the user service
type UserService interface {
	New(user NewUserRequest) error
}

// UserRepository is aht lets our service do db operations
// withour knowng anything about the implimentation
type UserRepository interface {
	CreateUser(NewUserRequest) error
}

type userService struct {
	storage UserRepository
}

// NewUserService ...
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func (u *userService) New(user NewUserRequest) error {
	// do some basic validation
	log.Printf("%+v", user.WeightGoal)
	if user.Email == "" {
		return errors.New("user service - email required")
	}

	if user.Name == "" {
		return errors.New("user service - name required")
	}

	if user.WeightGoal == "" {
		return errors.New("user service - weight goal required")
	}

	// do some basic normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.ToLower(user.Email)

	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
