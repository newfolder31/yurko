package usecases

type UserRepository interface {
	Store(user User) error
	FindById(id int) (User, error)
}

type User struct {
	Email    string
	Password string

	IsAdmin bool

	IsActive bool

	FirstName   string
	LastName    string
	FathersName string
}
