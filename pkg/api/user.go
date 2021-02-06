package api

// UserServer contains the methods of the user service
type UserService interface{}

// UserRepository is aht lets our service do db operations
// withour knowng anything about the implimentation
type UserRepository interface{}

type userService struct {
	storage UserRepository
}

// NewUserService ...
func NewUserService(userRepo UserRepository) UserRepository {
	return &userService{
		storage: userRepo,
	}

}
